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

func buildRamResourceSharesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)
	object := types.ResourceShare{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetResourceShares(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetResourceSharesOutput{ResourceShares: []types.ResourceShare{object}}, nil)

	buildRamResourceShareAssociatedResourcesMock(t, m)
	buildRamResourceShareAssociatedPrincipalsMock(t, m)

	return client.Services{Ram: m}
}

func TestRamResourceShares(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceShares(), buildRamResourceSharesMock, client.TestOptions{})
}
