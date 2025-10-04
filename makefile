# Makefile for monitorbuddy (cross-platform)

APP      ?= mbuddy
PKG      ?= ./cmd/monitorbuddy
GOFLAGS  ?= -v

# --- Derive OS/Arch -------------------------------------------------
GOOS     ?= $(shell go env GOOS)
GOARCH   ?= $(shell go env GOARCH)

# --- Version metadata via ldflags -----------------------------------
VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)

# DATE + env wrapper differ per-OS
ifeq ($(GOOS),windows)
  # ISO8601 UTC via PowerShell
  DATE    := $(shell powershell -NoProfile -Command "(Get-Date).ToUniversalTime().ToString('yyyy-MM-ddTHH:mm:ssZ')")
  BUILDENV := set CGO_ENABLED=1 &&
  EXE      := .exe
  LDQUOTE  := "
else
  DATE    := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
  BUILDENV := CGO_ENABLED=1
  EXE      :=
  LDQUOTE  := '
endif

LDFLAGS  ?= -s -w \
	-X leovanalphen/monitorbuddy/internal/app.Version=$(VERSION) \
	-X leovanalphen/monitorbuddy/internal/app.Commit=$(COMMIT) \
	-X leovanalphen/monitorbuddy/internal/app.Date=$(DATE)

# --- Targets --------------------------------------------------------

.PHONY: all build test clean help
all: build

build: ## Build for host OS/Arch with version metadata
	@echo ">> Building $(APP) ($(GOOS)/$(GOARCH)) version=$(VERSION) commit=$(COMMIT)"
	$(BUILDENV) go build $(GOFLAGS) -ldflags $(LDQUOTE)$(LDFLAGS)$(LDQUOTE) -o $(APP)$(EXE) $(PKG)

test:  ## Run tests
	@echo ">> go test"
	go test ./...

clean: ## Clean build artifacts and caches
	go clean -cache
	go clean -x
	-@rm -f $(APP) $(APP).exe

# Optional helpers if you want explicit OS targets

.PHONY: linux mac win
linux:
	@echo ">> Building on Linux"
	GOOS=linux CGO_ENABLED=1 go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(APP) $(PKG)

mac:   # requires: brew install hidapi
	@echo ">> Building on macOS with Homebrew hidapi"
	GOOS=darwin CGO_ENABLED=1 \
	CGO_CFLAGS="-I$$(brew --prefix hidapi)/include" \
	CGO_LDFLAGS="-L$$(brew --prefix hidapi)/lib" \
	go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(APP) $(PKG)

win:
	@echo ">> Building on Windows"
	GOOS=windows CGO_ENABLED=1 go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(APP).exe $(PKG)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'
