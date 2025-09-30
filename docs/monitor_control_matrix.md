📊 Monitor Control Matrix
=========================

| Vendor | Brightness / Contrast / Color (VCP 0x10/0x12/0x14…) | Input Select (VCP 0x60) | Audio Volume (0x62) | Preset Modes (0xDC) | KVM Switch | PiP / PBP | RGB / LEDs | Crosshair / Game OSD | Other Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **Dell / Alienware** | ✅ DDC/CI | ✅ DDC/CI | ✅ DDC/CI (if audio) | ✅ DDC/CI | ❌ | ❌ (OSD only) | ❌ | ❌ | Dell Command | Monitor is 100% DDC/CI; very complete. |
| **Gigabyte / AORUS** | ✅ DDC/CI | ✅ DDC/CI + HID | ✅ DDC/CI (if audio) | ✅ DDC/CI + HID | ✅ USB HID | ✅ USB HID | ✅ USB HID | ✅ USB HID | “Dashboard” telemetry via HID. |
| **ASUS (ROG / ProArt w/ KVM)** | ✅ DDC/CI | ✅ DDC/CI | ✅ DDC/CI (some) | ✅ DDC/CI | ✅ USB HID | ✅ USB HID | ✅ USB HID (Aura Sync) | ✅ USB HID | Their “DisplayWidget Center” talks HID. |
| **MSI (Gaming / Prestige w/ KVM)** | ✅ DDC/CI | ✅ DDC/CI | ✅ DDC/CI (some) | ✅ DDC/CI | ✅ USB HID | ✅ USB HID | ✅ USB HID (Mystic Light) | ✅ USB HID | “Gaming OSD” app = HID. |
| **BenQ (PD, EX, Zowie)** | ✅ DDC/CI | ✅ DDC/CI | ✅ DDC/CI (some) | ✅ DDC/CI | ✅ via USB HID (Hotkey Puck) | ✅ via USB HID | ❌ | ❌ (some Zowie HID extras) | Hotkey Puck = vendor HID reports. |
| **LG (UltraFine, UltraGear w/ KVM)** | ✅ DDC/CI | ✅ DDC/CI | ✅ DDC/CI (if speakers) | ✅ DDC/CI | ✅ USB HID (on KVM models) | ✅ USB HID (on high-end models) | ✅ USB HID (some gaming models) | ✅ USB HID (Game Mode OSD) | “OnScreen Control” uses both DDC/CI + HID depending on model. |
| **HP / Z-series** | ✅ DDC/CI | ✅ DDC/CI | ✅ (if audio) | ✅ DDC/CI | ❌ | ❌ | ❌ | ❌ | Tend to stick strictly to MCCS. |
| **Eizo / NEC** | ✅ DDC/CI | ✅ DDC/CI | ✅ DDC/CI (if audio) | ✅ DDC/CI | ❌ | ❌ | ❌ | ❌ | Very standards-compliant. |

---

🔑 Key Takeaways
================

* **Always present over DDC/CI**: brightness, contrast, color gains, input select, audio (if built-in), preset modes.
* **Vendor extras** (KVM, PiP/PBP, LEDs, crosshair, telemetry) → **almost always USB HID or vendor-specific USB protocol**.
* **Dell/Alienware, HP, Eizo, NEC** = safe to treat as DDC/CI-only.
* **Gaming vendors (Gigabyte, ASUS, MSI, BenQ, LG)** = dual path (DDC/CI basics + HID extras).

