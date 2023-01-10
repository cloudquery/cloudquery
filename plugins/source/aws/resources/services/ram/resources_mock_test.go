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

func buildRamResourcesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)
	object := types.Resource{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.ListResourcesOutput{
			Resources: []types.Resource{object},
		}, nil).MinTimes(1)

	return client.Services{
		Ram: m,
	}
}
func TestRamResources(t *testing.T) {
	client.AwsMockTestHelper(t, Resources(), buildRamResourcesMock, client.TestOptions{})
}
