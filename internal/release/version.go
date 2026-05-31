package release

import "runtime/debug"

var versionString string

// Version returns the current version of the application.
func Version() string {
	if versionString != "" {
		return versionString
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		v := info.Main.Version
		if v != "" && v != "(devel)" {
			return v
		}
	}

	return "dev"
}
