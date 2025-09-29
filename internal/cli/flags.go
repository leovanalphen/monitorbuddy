package cli

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	flagList      = flag.Bool("list", false, "List candidate HID monitor devices and exit")
	flagMonIdxs   = flag.String("mon", "", "Comma-separated monitor indices to target (from -list). If empty, use first match.")
	flagVID       = flag.String("vid", "0x0bda", "Vendor ID filter in hex (e.g., 0x0bda for Realtek). Use 0 to allow any.")
	flagPID       = flag.String("pid", "0x1100", "Product ID filter in hex. Use 0 to allow any.")
	flagIncludeGB = flag.Bool("include-gigabyte", true, "Include Gigabyte/Aorus vendor VCP properties")
	flagFrameBase = flag.Int("frame-base", 0x81, "Frame length base byte (try 0x81 or 0x80 depending on firmware)")
	flagDry       = flag.Bool("n", false, "Dry run: print frames instead of writing")

	flagGetName = flag.String("get", "", "Get property by name (e.g., brightness)")
	flagGetNum  = flag.Uint("getNum", 0, "Get property by VCP numeric code")
	flagSetName = flag.String("set", "", "Set property by name (requires -val)")
	flagSetNum  = flag.Uint("setNum", 0, "Set property by VCP numeric code (requires -val)")
	flagVal     = flag.Int("val", -1, "Value for -set or -setNum")

	flagHelpProps = flag.Bool("props", false, "List known property names and exit")
)

func Parse() { flag.Parse() }

func FlagList() bool      { return *flagList }
func FlagMonIdxs() string { return *flagMonIdxs }
func FlagVID() string     { return *flagVID }
func FlagPID() string     { return *flagPID }
func FlagIncludeGB() bool { return *flagIncludeGB }
func FlagFrameBase() int  { return *flagFrameBase }
func FlagDry() bool       { return *flagDry }
func FlagGetName() string { return *flagGetName }
func FlagGetNum() uint    { return *flagGetNum }
func FlagSetName() string { return *flagSetName }
func FlagSetNum() uint    { return *flagSetNum }
func FlagVal() int        { return *flagVal }
func FlagHelpProps() bool { return *flagHelpProps }

func ParseHexU16(s string) (uint16, error) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		v, err := strconv.ParseUint(s[2:], 16, 16)
		return uint16(v), err
	}
	v, err := strconv.ParseUint(s, 16, 16)
	return uint16(v), err
}

func MustVIDPID() (uint16, uint16) {
	vid, err := ParseHexU16(*flagVID)
	if err != nil {
		panic(fmt.Errorf("invalid -vid: %v", err))
	}
	pid, err := ParseHexU16(*flagPID)
	if err != nil {
		panic(fmt.Errorf("invalid -pid: %v", err))
	}
	return vid, pid
}

func ParseIndexList(s string) ([]int, error) {
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	parts := strings.Split(s, ",")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		v, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}
