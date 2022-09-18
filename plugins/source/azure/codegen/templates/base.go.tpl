// Auto generated code - DO NOT EDIT.

package {{.AzureService | ToLower}}

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/pkg/errors"
    {{template "imports.go.tpl" .}}
)

{{ if .IsRelation }}
func {{.AzureSubService | ToLowerCamel }}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}
{{ else }}
func {{.AzureSubService}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}
{{ end }}

{{range .Helpers}}
{{.}}
{{end}}