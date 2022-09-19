// Auto generated code - DO NOT EDIT.

package container

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
)

func createReplicationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockContainerReplicationsClient(ctrl)
	s := services.Services{
		Container: services.ContainerClient{
			Replications: mockClient,
		},
	}

	data := containerregistry.Replication{}
	require.Nil(t, faker.FakeObject(&data))

	result := containerregistry.NewReplicationListResultPage(containerregistry.ReplicationListResult{Value: &[]containerregistry.Replication{data}}, func(ctx context.Context, result containerregistry.ReplicationListResult) (containerregistry.ReplicationListResult, error) {
		return containerregistry.ReplicationListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
