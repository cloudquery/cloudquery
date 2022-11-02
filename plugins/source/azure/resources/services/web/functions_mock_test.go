// Auto generated code - DO NOT EDIT.

package web

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func createFunctionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebFunctionsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			Functions: mockClient,
		},
	}

	data := web.FunctionEnvelope{}
	require.Nil(t, faker.FakeObject(&data))

	result := web.NewFunctionEnvelopeCollectionPage(web.FunctionEnvelopeCollection{Value: &[]web.FunctionEnvelope{data}}, func(ctx context.Context, result web.FunctionEnvelopeCollection) (web.FunctionEnvelopeCollection, error) {
		return web.FunctionEnvelopeCollection{}, nil
	})

	mockClient.EXPECT().ListFunctions(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
