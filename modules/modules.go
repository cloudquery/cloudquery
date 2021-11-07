package modules

import "embed"

//go:embed configs/*.hcl
var FS embed.FS
