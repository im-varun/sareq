package release

var versionString string = "dev"

// Version returns the current version of the application.
func Version() string {
	return versionString
}
