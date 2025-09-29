# Monitor Buddy — Property Reference

This document lists the supported **VESA VCP (Virtual Control Panel)** codes and known **vendor-specific extensions** (initially focused on Gigabyte/AORUS), grouped by type.

---

## Legend

- **RW** – Read/Write (continuous values)
- **RO** – Read-Only (queried values)
- **WO** – Write-Only (momentary triggers)

---

## Standard VCP Properties (subset)

| Name              | Type | Range   | VCP  | Description                     |
| ----------------- | ---- | ------- | ---- | ------------------------------- |
| brightness        | RW   | 0–100   | 0x10 | Monitor brightness              |
| contrast          | RW   | 0–100   | 0x12 | Monitor contrast                |
| sharpness         | RW   | 0–100   | 0x72 | Image sharpness                 |
| sharpness-alt     | RW   | 0–100   | 0x87 | Alternate sharpness control     |
| volume            | RW   | 0–100   | 0x62 | Speaker volume                  |
| input-source      | RW   | 0–15    | 0x60 | Input source selection          |
| osd-language      | RW   | 0–255   | 0x80 | OSD language index              |
| backlight         | RW   | 0–100   | 0x8A | Backlight level                 |
| color-temp        | RW   | 0–100   | 0x0C | Color temperature request       |
| display-scaling   | RW   | 0–6     | 0xC0 | Display scaling mode            |
| hue               | RW   | 0–100   | 0x73 | Hue adjustment                  |
| saturation        | RW   | 0–100   | 0x74 | Saturation adjustment           |
| red-black-level   | RW   | 0–100   | 0x90 | Red black level                 |
| green-black-level | RW   | 0–100   | 0x91 | Green black level               |
| blue-black-level  | RW   | 0–100   | 0x92 | Blue black level                |
| clock             | RW   | 0–65535 | 0xB0 | Pixel clock                     |
| phase             | RW   | 0–65535 | 0xB2 | Pixel phase                     |
| active-control    | RW   | 0–1     | 0xB6 | Active control (auto adjust)    |
| image-mode        | RW   | 0–255   | 0x70 | Image/preset mode               |

### Read-Only VCP Properties

| Name           | Type | VCP  | Description                   |
| -------------- | ---- | ---- | ----------------------------- |
| hfreq          | RO   | 0xAC | Horizontal frequency (kHz)    |
| vfreq          | RO   | 0xAE | Vertical frequency (Hz)       |
| display-tech   | RO   | 0x7C | Display technology type       |
| fw-level       | RO   | 0x7E | Firmware level                |
| controller-id  | RO   | 0x84 | Display controller ID         |
| vcp-version    | RO   | 0xDF | VCP version                   |
| app-enable-key | RO   | 0xC6 | Application enable key        |
| display-mode   | RO   | 0xC8 | Display mode info             |
| display-profile| RO   | 0xCC | Display profile info          |

### Write-Only (Momentary) VCP Properties

| Name               | Type | VCP  | Description                         |
| ------------------ | ---- | ---- | ----------------------------------- |
| factory-reset       | WO   | 0x04 | Restore factory defaults            |
| factory-bc-reset    | WO   | 0x05 | Reset brightness/contrast           |
| factory-color-reset | WO   | 0x08 | Reset color calibration             |
| color-reset         | WO   | 0x76 | Reset color settings                |

---

## Gigabyte / AORUS Vendor Extensions (0xE000–0xEFFF)

| Name         | Type | Range   | Hex  | Description                          |
| ------------ | ---- | ------- | ---- | ------------------------------------ |
| low-blue-light | RW | 0–10    | E00B | Blue-light reduction                 |
| kvm-switch     | RW | 0–1     | E069 | Switch KVM to device 0/1             |
| colour-mode    | RW | 0–3     | E003 | 0=cool, 1=sRGB, 2=warm, 3=user       |
| count          | RW | 0–99    | E02A | On-screen counter value              |
| counter        | RW | 0–1     | E028 | Counter visibility (0=off,1=on)      |
| crosshair      | RW | 0–4     | E037 | Gaming crosshair overlay             |
| refresh        | RW | 0–1     | E022 | Show refresh rate on screen          |
| rgb-red        | RW | 0–100   | E004 | Custom RGB (red)                     |
| rgb-green      | RW | 0–100   | E005 | Custom RGB (green)                   |
| rgb-blue       | RW | 0–100   | E006 | Custom RGB (blue)                    |
| timer          | RW | 0–2     | E023 | Timer mode (0=off,1=up,2=down)       |
| timer-location | RW | 0–0x102 | E02B | Timer screen position                |
| timer-pause    | RW | 0–1     | E027 | Pause/resume timer                   |
| timer-set      | RW | 0–0x633C| E026 | Timer duration (MM:SS)               |
| pbp-mode       | RW | 0–2     | E00E | Picture-by-picture / picture-in-pic  |
| pbp-source     | RW | 0–3     | E00F | Secondary video source               |
| pbp-switch     | WO | 1       | E010 | Swap PBP/PIP windows                 |
| pbp-audio      | WO | 1       | E013 | Toggle PBP/PIP audio                 |
| pip-size       | RW | 0–2     | E014 | PIP window size                      |
| pip-location   | RW | 0–3     | E015 | PIP window location                  |
| source         | RW | 0–3     | E02D | Primary source                       |
| audio-input    | RW | 0–2     | E02E | Audio input routing                  |
| kvm-usb-b      | RW | 0–3     | E06B | Map USB-B upstream                   |
| kvm-usb-c      | RW | 0–3     | E06C | Map USB-C upstream                   |

---

## Notes

- Vendor-specific codes are **not standardized**; values may differ by model or firmware version.
- Use `-props` to list all properties recognized by your build of Monitor Buddy.
- Safe practice: use dry-run (`-n`) before writing to unfamiliar properties.
- If your monitor is not listed here, contributions are welcome!

---
