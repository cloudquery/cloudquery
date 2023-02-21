package plugin

import (
"github.com/cloudquery/plugin-sdk/schema"
{{- range $resource := . }}
  "github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/{{ $resource.Service }}"
{{- end }}
)

func tables() []*schema.Table {
return []*schema.Table{
{{- range $resource := . }}
    {{ $resource.Service }}.{{ $resource.TableName | ToPascal }}(),
{{- end }}
}
}
