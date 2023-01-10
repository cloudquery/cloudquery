package ram

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRamResourceTypesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)
	object := types.ServiceNameAndResourceType{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListResourceTypes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.ListResourceTypesOutput{ResourceTypes: []types.ServiceNameAndResourceType{object}}, nil)

	return client.Services{Ram: m}
}

func TestRamResourceTypes(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceTypes(), buildRamResourceTypesMock, client.TestOptions{})
}
