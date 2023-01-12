package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"
{{- range $resource := . }}
	"github.com/cloudquery/cloudquery/plugins/source/oracle/resources/services/{{ $resource.Service }}"
{{- end }}
)

func AutogenTables() []*schema.Table {
	return []*schema.Table{
    {{- range $resource := . }}
        {{ $resource.Service }}.{{ $resource.SubService | ToCamel }}(),
    {{- end }}
	}
}
