// Auto generated code - DO NOT EDIT.

package {{.AzurePackageName}}

import (
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

func create{{ .AzureService }}{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMock{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.NetworksClient{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	require.Nil(t, faker.FakeData(&data))

	list := {{ .AzurePackageName }}.{{ .AzureStructName }}ListResult{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}
	mockClient.EXPECT().ListAll(gomock.Any()).Return(list, nil)
	return s
}

func Test{{ .AzureService }}{{ .AzureSubService }}(t *testing.T) {
	client.AzureMockTestHelper(t, {{ .AzureService }}{{ .AzureSubService }}(), create{{ .AzureService }}{{ .AzureSubService }}Mock, client.TestOptions{})
}
