package app

import (
	"encoding/hex"
	"fmt"
	"time"

	"leovanalphen/monitorbuddy/internal/cli"
	"leovanalphen/monitorbuddy/internal/ddcci"
	"leovanalphen/monitorbuddy/internal/hidutil"
	"leovanalphen/monitorbuddy/internal/properties"
)

func Run() error {
	// VID/PID & list
	vid, pid := cli.MustVIDPID()
	devs, err := hidutil.List(vid, pid)
	if err != nil {
		return fmt.Errorf("enumerate: %w", err)
	}

	if cli.FlagList() {
		if len(devs) == 0 {
			fmt.Println("No HID devices match the filters.")
			return nil
		}
		fmt.Printf("Devices matching VID=%04x PID=%04x\n", vid, pid)
		hidutil.PrintList(devs)
		return nil
	}
	if len(devs) == 0 {
		return fmt.Errorf("no devices found for VID=%04x PID=%04x. Try -vid 0 -pid 0 and use -list", vid, pid)
	}

	idxs, err := cli.ParseIndexList(cli.FlagMonIdxs())
	if err != nil {
		return fmt.Errorf("bad -mon argument: %w", err)
	}

	handles, err := hidutil.OpenByIndices(devs, idxs)
	if err != nil {
		return err
	}
	defer func() {
		for _, h := range handles {
			h.Close()
		}
	}()

	doGet := cli.FlagGetName() != "" || cli.FlagGetNum() != 0
	doSet := cli.FlagSetName() != "" || cli.FlagSetNum() != 0

	if doGet && doSet {
		return fmt.Errorf("choose exactly one of GET (-get/-getNum) or SET (-set/-setNum)")
	}
	if !doGet && !doSet {
		return fmt.Errorf("no action. Use -list, -props, -get/-getNum, or -set/-setNum with -val")
	}

	if doGet {
		propID, meta, err := resolveGetTarget()
		if err != nil {
			return err
		}
		for i, h := range handles {
			if cli.FlagDry() {
				buf := ddcci.BuildGetFrame(propID, cli.FlagFrameBase())
				fmt.Printf("[mon %d] DRY get 0x%02X\n%s\n", i, propID, hex.Dump(buf))
				continue
			}
			if _, err := h.Write(ddcci.BuildGetFrame(propID, cli.FlagFrameBase())); err != nil {
				return fmt.Errorf("[mon %d] write(get): %w", i, err)
			}
			_ = h.SetNonblock(true)
			defer h.SetNonblock(false)

			time.Sleep(5 * time.Millisecond)

			rb := make([]byte, ddcci.ReportSize)
			n, err := h.Read(rb)
			if err != nil && !ddcci.WouldBlock(err) {
				return fmt.Errorf("[mon %d] read(get): %w", i, err)
			}
			if n == 0 {
				time.Sleep(5 * time.Millisecond)
				n, err = h.Read(rb)
				if err != nil && !ddcci.WouldBlock(err) {
					return fmt.Errorf("[mon %d] read(get) retry: %w", i, err)
				}
			}
			cur, max, vType, perr := ddcci.ParseGetReply(rb[:n])
			if perr != nil {
				return fmt.Errorf("[mon %d] parse: %v\nDump:\n%s", i, perr, hex.Dump(rb[:n]))
			}
			label := vcpLabel(propID, meta)
			fmt.Printf("[mon %d] %s => cur=%d max=%d (type=0x%02X)\n", i, label, cur, max, vType)
		}
		return nil
	}

	// SET
	if cli.FlagVal() == -1 {
		return fmt.Errorf("SET requires -val")
	}
	propID, meta, err := resolveSetTarget(uint16(cli.FlagVal()))
	if err != nil {
		return err
	}

	for i, h := range handles {
		buf := ddcci.BuildSetFrame(propID, uint16(cli.FlagVal()), cli.FlagFrameBase())
		if cli.FlagDry() {
			fmt.Printf("[mon %d] DRY set 0x%02X := 0x%04X\n%s\n", i, propID, cli.FlagVal(), hex.Dump(buf))
			continue
		}
		if _, err := h.Write(buf); err != nil {
			return fmt.Errorf("[mon %d] write(set): %w", i, err)
		}
		label := vcpLabel(propID, meta)
		fmt.Printf("[mon %d] Set %s := %d (0x%04X)\n", i, label, cli.FlagVal(), cli.FlagVal())
	}
	return nil
}

func resolveGetTarget() (uint16, *properties.Property, error) {
	if cli.FlagGetNum() != 0 {
		id := uint16(cli.FlagGetNum())
		return id, nil, nil
	}
	p, ok := properties.LookupByName(cli.FlagGetName())
	if !ok {
		return 0, nil, fmt.Errorf("unknown property name: %s (use -props to list)", cli.FlagGetName())
	}
	return p.ID, &p, nil
}

func resolveSetTarget(val uint16) (uint16, *properties.Property, error) {
	if cli.FlagSetNum() != 0 {
		return uint16(cli.FlagSetNum()), nil, nil
	}
	p, ok := properties.LookupByName(cli.FlagSetName())
	if !ok {
		return 0, nil, fmt.Errorf("unknown property name: %s (use -props to list)", cli.FlagSetName())
	}
	if p.Type == "RO" {
		return 0, nil, fmt.Errorf("property %s is read-only", p.Name)
	}
	if p.Min != 0 || p.Max != 0 {
		if int(val) < int(p.Min) || int(val) > int(p.Max) {
			return 0, nil, fmt.Errorf("value %d out of range %d-%d for %s", val, p.Min, p.Max, p.Name)
		}
	}
	return p.ID, &p, nil
}

func vcpLabel(id uint16, meta *properties.Property) string {
	if meta == nil {
		return fmt.Sprintf("VCP 0x%02X", id)
	}
	return fmt.Sprintf("%s (%s, VCP 0x%02X)", meta.Name, meta.Type, id)
}
