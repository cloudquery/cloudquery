package plugin

// Don't move this file to a different package, it's used by Go releaser to embed the version in the binary.
var (
	Name    = "reverser"
	Kind    = "destination" // breaks otherwise
	Team    = "cloudquery"
	Version = "development"
)
