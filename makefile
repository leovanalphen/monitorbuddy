# Makefile for monitorbuddy

APP      ?= mbuddy
OUT      ?= $(APP).exe
PKG      ?= ./cmd/monitorbuddy
LDFLAGS  ?=
GOFLAGS  ?= -v

# ---- Common targets ------------------------------------------------

.PHONY: all clean
all: build

clean:
	go clean -cache
	go clean -x
	-@rm -f $(APP) $(APP).exe

# ---- Windows (MSYS2 UCRT64 / MinGW64) ------------------------------

# Recommended: UCRT64 toolchain
.PHONY: win-ucrt64
win-ucrt64:
	@echo ">> Building with MSYS2 UCRT64 (gcc)"
	set CGO_ENABLED=1 && \
	set CC=C:\msys64\ucrt64\bin\gcc.exe && \
	set CXX=C:\msys64\ucrt64\bin\g++.exe && \
	set PATH=C:\msys64\ucrt64\bin;%PATH% && \
	go build $(GOFLAGS) -o $(OUT) $(PKG)

# Alternative: MinGW64 toolchain
.PHONY: win-mingw64
win-mingw64:
	@echo ">> Building with MSYS2 MinGW64 (gcc)"
	set CGO_ENABLED=1 && \
	set CC=C:\msys64\mingw64\bin\gcc.exe && \
	set CXX=C:\msys64\mingw64\bin\g++.exe && \
	set PATH=C:\msys64\mingw64\bin;%PATH% && \
	go build $(GOFLAGS) -o $(OUT) $(PKG)

# If your setup needs SetupAPI explicitly (usually not with MSYS2):
.PHONY: win-ucrt64-setupapi
win-ucrt64-setupapi:
	@echo ">> Building UCRT64 with -lsetupapi"
	set CGO_ENABLED=1 && \
	set CC=C:\msys64\ucrt64\bin\gcc.exe && \
	set CXX=C:\msys64\ucrt64\bin\g++.exe && \
	set PATH=C:\msys64\ucrt64\bin;%PATH% && \
	set CGO_LDFLAGS=-lsetupapi && \
	go build $(GOFLAGS) -o $(OUT) $(PKG)

# ---- Linux ---------------------------------------------------------

.PHONY: linux
linux:
	@echo ">> Building on Linux"
	CGO_ENABLED=1 go build $(GOFLAGS) -o $(APP) $(PKG)

# ---- macOS (Homebrew hidapi) --------------------------------------

# Assumes: brew install hidapi
.PHONY: mac
mac:
	@echo ">> Building on macOS with Homebrew hidapi"
	CGO_ENABLED=1 \
	CGO_CFLAGS="-I$$(brew --prefix hidapi)/include" \
	CGO_LDFLAGS="-L$$(brew --prefix hidapi)/lib" \
	go build $(GOFLAGS) -o $(APP) $(PKG)

# ---- Convenience meta target --------------------------------------

.PHONY: build
build: win-ucrt64
