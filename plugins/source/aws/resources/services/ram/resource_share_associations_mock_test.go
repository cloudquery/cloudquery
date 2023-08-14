package ram

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRamResourceShareAssociationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)

	object := types.ResourceShareAssociation{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().GetResourceShareAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetResourceShareAssociationsOutput{ResourceShareAssociations: []types.ResourceShareAssociation{object}}, nil).MinTimes(1)
	return client.Services{Ram: m}
}

func TestRamResourceShareAssociatedResources(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceShareAssociations(), buildRamResourceShareAssociationsMock, client.TestOptions{})
}
