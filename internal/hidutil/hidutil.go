package hidutil

import (
	"fmt"
	"strings"

	"github.com/sstallion/go-hid"
)

type DeviceRef struct {
	Index int
	Info  *hid.DeviceInfo // VendorID, ProductID, MfrStr, ProductStr, SerialNbr, Path
}

func List(vid, pid uint16) ([]DeviceRef, error) {
	var devs []*hid.DeviceInfo
	err := hid.Enumerate(vid, pid, func(info *hid.DeviceInfo) error {
		devs = append(devs, info)
		return nil
	})
	if err != nil {
		return nil, err
	}
	out := make([]DeviceRef, 0, len(devs))
	for i, d := range devs {
		out = append(out, DeviceRef{Index: i, Info: d})
	}
	return out, nil
}

func PrintList(devs []DeviceRef) {
	for _, d := range devs {
		man := strings.TrimSpace(d.Info.MfrStr)
		prod := strings.TrimSpace(d.Info.ProductStr)
		serial := strings.TrimSpace(d.Info.SerialNbr)
		fmt.Printf("[%d] %04x:%04x  %-20s  %-30s  %s\n",
			d.Index, d.Info.VendorID, d.Info.ProductID, man, prod, serial)
	}
}

func OpenByIndices(devs []DeviceRef, idxs []int) ([]*hid.Device, error) {
	if len(idxs) == 0 && len(devs) > 0 {
		idxs = []int{devs[0].Index} // default to first
	}
	var handles []*hid.Device
	for _, want := range idxs {
		var di *hid.DeviceInfo
		for _, dr := range devs {
			if dr.Index == want {
				di = dr.Info
				break
			}
		}
		if di == nil {
			return nil, fmt.Errorf("no device with index %d", want)
		}
		h, err := hid.OpenPath(di.Path)
		if err != nil {
			return nil, fmt.Errorf("open index %d failed: %w", want, err)
		}
		handles = append(handles, h)
	}
	return handles, nil
}
