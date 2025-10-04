package app

// These are filled via -ldflags during build/release.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)
