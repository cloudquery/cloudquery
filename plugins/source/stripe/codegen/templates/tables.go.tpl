package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
{{- range $resource := . }}
  "github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/{{ $resource.Service }}"
{{- end }}
)

func rawTables() []*schema.Table {
return []*schema.Table{
{{- range $resource := . }}
    {{ $resource.Service }}.{{ $resource.TableName | ToPascal }}(),
{{- end }}
}
}
