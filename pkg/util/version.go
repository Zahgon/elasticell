package util

// set on build time
var (
	GitCommit = ""
	BuildTime = ""
	GoVersion = ""
	Version   = ""
)

// PrintVersion print version info
func PrintVersion() bool { _ = "STUB: not implemented"; return false }
