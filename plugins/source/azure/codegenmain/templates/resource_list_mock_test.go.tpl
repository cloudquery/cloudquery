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
  {{range .MockImports}}
  "{{.}}"
  {{end}}
)

func createMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMock{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.NetworksClient{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	require.Nil(t, faker.FakeData(&data))

	page := {{ .AzurePackageName }}.New{{ .AzureStructName }}ListResultPage(network.{{ .AzureStructName }}ListResult{Value: &[]network.{{ .AzureStructName }}{data}}, func(ctx context.Context, result network.{{ .AzureStructName }}ListResult) (network.{{ .AzureStructName }}ListResult, error) {
		return network.{{ .AzureStructName }}ListResult{}, nil
	})

	mockClient.EXPECT().{{ .ListFunction }}(gomock.Any()).Return(page, nil)
	return s
}

func Test{{ .AzureService }}{{ .AzureSubService }}(t *testing.T) {
	client.AzureMockTestHelper(t, {{ .AzureService }}{{ .AzureSubService }}(), createMock, client.TestOptions{})
}
