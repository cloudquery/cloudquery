package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildInventories(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.InventoryResultEntity
	require.NoError(t, faker.FakeObject(&i))

	mock.EXPECT().GetInventory(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.GetInventoryOutput{Entities: []types.InventoryResultEntity{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestInventories(t *testing.T) {
	client.AwsMockTestHelper(t, Inventories(), buildInventories, client.TestOptions{})
}
