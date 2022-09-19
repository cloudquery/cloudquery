// Auto generated code - DO NOT EDIT.

package {{.AzureService | ToLower }}

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
    {{template "imports.go.tpl" .}}
)

{{ if not .IsRelation }}
func Test{{ .AzureService }}{{ .AzureSubService }}(t *testing.T) {
	client.MockTestHelper(t, {{ .AzureSubService }}(), create{{ .AzureSubService }}Mock)
}
{{ end }}

{{range .MockHelpers}}
{{.}}
{{end}}