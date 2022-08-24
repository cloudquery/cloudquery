// Auto generated code - DO NOT EDIT.

package {{.AzurePackageName}}

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func Test{{ .AzureService }}{{ .AzureSubService }}(t *testing.T) {
	client.AzureMockTestHelper(t, {{ .AzureService }}{{ .AzureSubService }}(), create{{ .AzureService }}{{ .AzureSubService }}Mock, client.TestOptions{})
}