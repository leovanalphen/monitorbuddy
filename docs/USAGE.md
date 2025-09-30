# Monitor Buddy — Usage Guide

This document provides detailed usage instructions for the **Monitor Buddy CLI**.

---

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows (MSYS2)](#windows-msys2)
- [Basic Commands](#basic-commands)
  - [Listing devices](#listing-devices)
  - [Reading properties](#reading-properties)
  - [Setting properties](#setting-properties)
- [Advanced Options](#advanced-options)
  - [VID/PID filters](#vidpid-filters)
  - [Vendor properties](#vendor-properties)
  - [Frame base tuning](#frame-base-tuning)
  - [Dry-run mode](#dry-run-mode)
- [Examples](#examples)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)

---

## Overview

`monitor-buddy` is a command-line tool to control monitor settings using **DDC/CI over USB HID**. It supports both standard VCP codes and vendor‑specific extensions (e.g., Gigabyte/AORUS).

## Installation

### Linux

```bash
sudo apt install libhidapi-dev libudev-dev
go install github.com/YOURNAME/monitor-buddy@latest
```

### macOS

```zsh
brew install hidapi
CGO_CFLAGS="-I$(brew --prefix hidapi)/include" \
CGO_LDFLAGS="-L$(brew --prefix hidapi)/lib" \
go install github.com/YOURNAME/monitor-buddy@latest
```

### Windows (MSYS2)

```cmd
# Ensure MSYS2 UCRT64 toolchain is installed
make win-ucrt64
```

## Basic Commands

### Listing devices

```bash
monitor-buddy -list
```

### Reading properties

```bash
monitor-buddy -get brightness -mon 0
```

### Setting properties

```bash
monitor-buddy -set brightness -val 50 -mon 0
```

## Advanced Options

### VID/PID filters

```bash
monitor-buddy -list -vid 0x0bda -pid 0x1100
```

### Vendor properties

Enable/disable vendor properties:

```bash
monitor-buddy -props -include-gigabyte=false
```

### Frame base tuning

Some monitors require a different frame length base:

```bash
monitor-buddy -set brightness -val 20 -frame-base 0x80
```

### Dry-run mode

```bash
monitor-buddy -set brightness -val 30 -n
```

## Examples

- Set brightness on multiple monitors:

```bash
monitor-buddy -set brightness -val 70 -mon 0,1
```

- Read firmware level:

```bash
monitor-buddy -get fw-level -mon 0
```

## Troubleshooting

- **No devices found:** try `-vid 0 -pid 0` and re-run `-list`.
- **Windows build errors:** ensure MSYS2 UCRT64 gcc is installed.
- **Monitor not responding:** verify USB upstream cable is connected.

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines.
