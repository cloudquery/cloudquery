package plugin

// Don't move this file to a different package, it's used by Go releaser to embed the version in the binary.
var (
	Name    = "s3"
	Kind    = "destination"
	Team    = "cloudquery"
	Version = "development"
)
