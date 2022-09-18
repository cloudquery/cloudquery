// Auto generated code - DO NOT EDIT.

package network

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func createFlowLogsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkFlowLogsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			FlowLogs: mockClient,
		},
	}

	data := network.FlowLog{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewFlowLogListResultPage(network.FlowLogListResult{Value: &[]network.FlowLog{data}}, func(ctx context.Context, result network.FlowLogListResult) (network.FlowLogListResult, error) {
		return network.FlowLogListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
