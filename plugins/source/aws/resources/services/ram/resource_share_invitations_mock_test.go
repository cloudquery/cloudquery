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

func buildRamResourceShareInvitationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)
	object := types.ResourceShareInvitation{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().GetResourceShareInvitations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetResourceShareInvitationsOutput{
			ResourceShareInvitations: []types.ResourceShareInvitation{object},
		}, nil)

	return client.Services{
		Ram: m,
	}
}
func TestRamResourceShareInvitations(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceShareInvitations(), buildRamResourceShareInvitationsMock, client.TestOptions{})
}
