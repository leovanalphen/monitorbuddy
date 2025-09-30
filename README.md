# Monitor Buddy — Cross-Platform Monitor Control

Monitorbuddy (`mbuddy`) is a command-line tool to **read and change** monitor settings via **DDC/CI over the Realtek HID** (as used by Gigabyte/AORUS OSD Sidekick and similar monitor software). It is open-source and extensible for other vendors.

---

## Project Overview

- **MonitorBuddy CLI (Open Source)**  
  Cross-platform monitor control for Linux, macOS, and Windows.  
  Licensed under [Apache 2.0](./LICENSE), with portions under MIT (see [NOTICE](./NOTICE)).

- **MonitorBuddy Stream Deck Plugin (Paid)**  
  A polished Elgato Stream Deck plugin powered by the CLI backend.  
  Allows one-click brightness, input switching, KVM toggles, and more.  
  *(coming soon)*

---

## Supported / Verified Monitors

- **Gigabyte M32U** (original target)
- **Gigabyte M32Q** ([text](https://github.com/kelvie/gbmonctl/issues/11))
- **Gigabyte M32QC** ([text](https://github.com/kelvie/gbmonctl/issues/7))
- **Gigabyte M27Q** ([text](https://github.com/kelvie/gbmonctl/issues/9))
- **AORUS CO49DQ** (verified; same HID protocol)
- **Any Realtek HID-based Gigabyte/AORUS monitor** (VID `0x0BDA`, PID `0x1100`)  
- **Other brands coming soon** (experimental)

Any Gigabyte that exposes a Realtek HID device (usually **VID `0x0BDA`**, PID `0x1100`) is *likely* compatible.

---

## Installation / Building

### Linux (Debian/Ubuntu)

```bash
sudo apt install libhidapi-dev libudev-dev
go install github.com/leovanalphen/monitorbuddy@latest
```

### macOS (Homebrew)

```bash
brew install hidapi
CGO_CFLAGS="-I$(brew --prefix hidapi)/include" \
CGO_LDFLAGS="-L$(brew --prefix hidapi)/lib" \
go install github.com/leovanalphen/monitorbuddy@latest
```

### Windows (MSYS2 **ucrt64** toolchain)

1. Install **MSYS2** from [https://www.msys2.org/](https://www.msys2.org/) and set up the **ucrt64** environment (install the `ucrt64` gcc packages if needed).

2. From **PowerShell**:

```powershell
$env:Path = "C:\\msys64\\ucrt64\\bin;$env:Path"
$env:CGO_ENABLED = "1"
$env:CC  = "C:\\msys64\\ucrt64\\bin\\gcc.exe"
$env:CXX = "C:\\msys64\\ucrt64\\bin\\g++.exe"

go build -v -o mbuddy.exe ./cmd/monitorbuddy
```

> If the linker later complains about SetupAPI, add:\
> `setx CGO_LDFLAGS "-lsetupapi"`

#### VS Code one-click build (Windows)

Create `.vscode/tasks.json`:

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "go build (ucrt64)",
      "type": "shell",
      "command": "go",
      "args": ["build", "-v", "-o", "mbuddy.exe", "."],
      "options": {
        "env": {
          "PATH": "C:/msys64/ucrt64/bin:${env:PATH}",
          "CGO_ENABLED": "1",
          "CC": "C:/msys64/ucrt64/bin/gcc.exe",
          "CXX": "C:/msys64/ucrt64/bin/g++.exe",
          "CGO_LDFLAGS": "-lsetupapi"
        }
      },
      "group": { "kind": "build", "isDefault": true }
    }
  ]
}
```

#### Makefile (Linux/macOS/Windows with MSYS2)

A `Makefile` is included for convenience with cross‑platform targets:

```makefile
# Common
make build     # default (on Windows calls win-ucrt64)
make clean     # clean artifacts and cache

# Windows (MSYS2)
make win-ucrt64       # build with UCRT64 toolchain (recommended)
make win-mingw64      # build with MinGW64 toolchain (alternative)
make win-ucrt64-setupapi  # build with SetupAPI explicitly linked

# Linux
make linux

# macOS
make mac
```

On Windows you must have MSYS2 installed and use the correct gcc (`ucrt64` recommended). On Linux/macOS, ensure `hidapi` development headers are installed.

Example:

```bash
# build and list devices
make build
./mbuddy.exe -list
```

---

## Usage

Detailed usage is documented in [docs/USAGE.md](./docs/USAGE.md).

Quick examples:

```bash
# Default filters (Realtek 0x0BDA:0x1100)
mbuddy -list

# Read current brightness
mbuddy -get brightness -mon 0

# Set brightness to 75
mbuddy -set brightness -val 75 -mon 0
```

---

## Property Reference

See [docs/PROPERTIES.md](./docs/PROPERTIES.md) for a full property reference, including:

- Standard VCP properties (DDC/CI)
- Gigabyte / AORUS vendor extensions

---

## Stream Deck Plugin (Commercial)

A paid Stream Deck plugin built on MonitorBuddy is in the works.\
👉 **[Plugin page — coming soon]**

This plugin uses MonitorBuddy as a backend to control monitors directly from your Stream Deck.

---

## Troubleshooting

- **No devices found**: try `-vid 0 -pid 0` and run `-list`.
- **Windows build errors**: use MSYS2 **ucrt64** or **mingw64** gcc; MSVC `cl/link` won’t work.
- **Nothing changes on monitor**: ensure the monitor’s **USB upstream** is connected.
- **Parsing errors on read**: some models are sensitive to timing; retries are included, try rerunning.

---

## Contributing

PRs welcome! Especially:

- Confirming additional compatible models
- Filling in enumerations (per-model input maps, languages)
- New vendor-specific codes and descriptions

## Acknowledgements

This project was originally inspired by and based on [gbmonctl](https://github.com/kelvie/gbmonctl) by Kelvie Wong, licensed under MIT.\
See [NOTICE](./NOTICE) for detailed attribution and licensing information.
