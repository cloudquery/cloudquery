// Code generated by codegen; DO NOT EDIT.

package plugin

import (
  {{- range .}}
  "github.com/cloudquery/cloudquery/plugins/source/azuredevops/resources/services/{{.Service}}"
  {{- end}}
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{
		  {{- range .}}
      {{.Service}}.{{.SubService | ToCamel}}(),
      {{- end}}
	}
}