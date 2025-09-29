package properties

import (
	"fmt"
	"sort"
)

type Property struct {
	Name        string
	Description string
	Min, Max    uint16
	ID          uint16
	Type        string // "RW", "RO", "WO"
}

var (
	all []Property
	byName map[string]Property
	includeGB bool
)

func BuildRegistry(includeGigabyte bool) {
	includeGB = includeGigabyte
	all = append([]Property(nil), std...) // start with std
	if includeGB {
		all = append(all, gigabyte...)
	}
	sort.Slice(all, func(i, j int) bool { return all[i].Name < all[j].Name })
	byName = make(map[string]Property, len(all))
	for _, p := range all { byName[p.Name] = p }
}

func LookupByName(name string) (Property, bool) { p, ok := byName[name]; return p, ok }
func All() []Property                           { return all }

func PrintPropsTable() {
	fmt.Println("Known properties:")
	namew := 0
	for _, p := range all { if len(p.Name) > namew { namew = len(p.Name) } }
	for _, p := range all {
		fmt.Printf("  %-*s  %-2s  VCP 0x%02X  range:%-11s  %s\n",
			namew, p.Name, p.Type, p.ID, formatRange(p), p.Description)
	}
}

func formatRange(p Property) string {
	if p.Min == 0 && p.Max == 0 { return "-" }
	return fmt.Sprintf("%d–%d", p.Min, p.Max)
}

// ---------- Standard subset ----------
var std = []Property{
	{"brightness", "Monitor brightness (0=darkest, 100=brightest)", 0, 100, 0x10, "RW"},
	{"contrast", "Monitor contrast (0=lowest, 100=highest)", 0, 100, 0x12, "RW"},
	{"sharpness", "Image sharpness (common alt code 0x87)", 0, 100, 0x72, "RW"},
	{"sharpness-alt", "Image sharpness alternate code", 0, 100, 0x87, "RW"},
	{"volume", "Speaker volume (0=mute, 100=max)", 0, 100, 0x62, "RW"},
	{"input-source", "Primary input source selection", 0, 15, 0x60, "RW"},
	{"osd-language", "OSD language index", 0, 255, 0x80, "RW"},
	{"backlight", "Backlight control", 0, 100, 0x8A, "RW"},
	{"color-temp", "Color temperature request", 0, 100, 0x0C, "RW"},
	{"display-scaling", "Display scaling", 0, 6, 0xC0, "RW"},
	{"hue", "Picture hue", 0, 100, 0x73, "RW"},
	{"saturation", "Picture saturation", 0, 100, 0x74, "RW"},
	{"red-black-level", "Red black level", 0, 100, 0x90, "RW"},
	{"green-black-level", "Green black level", 0, 100, 0x91, "RW"},
	{"blue-black-level", "Blue black level", 0, 100, 0x92, "RW"},
	{"clock", "Pixel clock", 0, 65535, 0xB0, "RW"},
	{"phase", "Pixel phase", 0, 65535, 0xB2, "RW"},
	{"active-control", "Active control", 0, 1, 0xB6, "RW"},
	{"image-mode", "Image mode / preset", 0, 255, 0x70, "RW"},
	// RO
	{"hfreq", "Horizontal frequency (kHz)", 0, 0, 0xAC, "RO"},
	{"vfreq", "Vertical frequency (Hz)", 0, 0, 0xAE, "RO"},
	{"display-tech", "Display technology type", 0, 0, 0x7C, "RO"},
	{"fw-level", "Display firmware level", 0, 0, 0x7E, "RO"},
	{"controller-id", "Display controller ID", 0, 0, 0x84, "RO"},
	{"vcp-version", "VCP version", 0, 0, 0xDF, "RO"},
	{"app-enable-key", "Application enable key", 0, 0, 0xC6, "RO"},
	{"display-mode", "Display mode info", 0, 0, 0xC8, "RO"},
	{"display-profile", "Display profile", 0, 0, 0xCC, "RO"},
	// WO
	{"factory-reset", "Restore factory defaults", 0, 0, 0x04, "WO"},
	{"factory-bc-reset", "Restore brightness/contrast defaults", 0, 0, 0x05, "WO"},
	{"factory-color-reset", "Restore factory color defaults", 0, 0, 0x08, "WO"},
	{"color-reset", "Color reset", 0, 0, 0x76, "WO"},
}

// ---------- Gigabyte vendor subset ----------
var gigabyte = []Property{
	{"low-blue-light", "Blue light reduction (0-10)", 0, 10, 0xe00b, "RW"},
	{"kvm-switch", "KVM switch (0=device1, 1=device2)", 0, 1, 0xe069, "RW"},
	{"colour-mode", "Color preset (0=cool,1=sRGB,2=warm,3=user)", 0, 3, 0xe003, "RW"},
	{"count", "On-screen counter value (0-99)", 0, 99, 0xe02a, "RW"},
	{"counter", "On-screen counter visibility (0=hide,1=show)", 0, 1, 0xe028, "RW"},
	{"crosshair", "Gaming crosshair overlay (0=off,1-4=styles)", 0, 4, 0xe037, "RW"},
	{"refresh", "Refresh rate OSD (0=hide,1=show)", 0, 1, 0xe022, "RW"},
	{"rgb-red", "Custom RGB red (0-100)", 0, 100, 0xe004, "RW"},
	{"rgb-green", "Custom RGB green (0-100)", 0, 100, 0xe005, "RW"},
	{"rgb-blue", "Custom RGB blue (0-100)", 0, 100, 0xe006, "RW"},
	{"timer", "Timer mode (0=off,1=count-up,2=count-down)", 0, 2, 0xe023, "RW"},
	{"timer-location", "Timer position (LB=0-2, HB=0-2)", 0x0000, 0x0102, 0xe02b, "RW"},
	{"timer-pause", "Timer pause/resume (0=pause,1=resume)", 0, 1, 0xe027, "RW"},
	{"timer-set", "Timer duration MMSS (e.g. 0x0530=5:30)", 0x0000, 0x633c, 0xe026, "RW"},
	{"pbp-mode", "Multi-window (0=single,1=PIP,2=PBP)", 0, 2, 0xe00e, "RW"},
	{"pbp-source", "Secondary source (0=HDMI1..3=USB-C)", 0, 3, 0xe00f, "RW"},
	{"pbp-switch", "Swap PIP/PBP windows (write 1)", 1, 1, 0xe010, "WO"},
	{"pbp-audio", "Toggle audio source (write 1)", 1, 1, 0xe013, "WO"},
	{"pip-size", "PIP size (0=Large,1=Med,2=Small)", 0, 2, 0xe014, "RW"},
	{"pip-location", "PIP pos (0=TL,1=TR,2=BL,3=BR)", 0, 3, 0xe015, "RW"},
	{"source", "Primary source (0=HDMI1..3=USB-C)", 0, 3, 0xe02d, "RW"},
	{"audio-input", "Audio input routing (0=main,1=aux,2=auto)", 0, 2, 0xe02e, "RW"},
	{"kvm-usb-b", "Map USB-B (0=HDMI1..3=USB-C)", 0, 3, 0xe06b, "RW"},
	{"kvm-usb-c", "Map USB-C (0=HDMI1..3=USB-C)", 0, 3, 0xe06c, "RW"},
}
