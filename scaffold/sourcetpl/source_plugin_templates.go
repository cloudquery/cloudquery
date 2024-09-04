package sourcetpl

import "embed"

//go:embed templates/source/*
var SourcePluginTemplatesFS embed.FS
