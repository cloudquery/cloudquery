package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildInventorySchemas(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.InventoryItemSchema
	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetInventorySchema(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.GetInventorySchemaOutput{Schemas: []types.InventoryItemSchema{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestInventorySchemas(t *testing.T) {
	client.AwsMockTestHelper(t, InventorySchemas(), buildInventorySchemas, client.TestOptions{})
}
