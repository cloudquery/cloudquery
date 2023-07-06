package wellarchitected

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildShareInvitationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWellarchitectedClient(ctrl)
	for _, shareResourceType := range types.ShareResourceType("").Values() {
		var summary types.ShareInvitationSummary
		require.NoError(t, faker.FakeObject(&summary))
		summary.ShareResourceType = shareResourceType

		m.EXPECT().ListShareInvitations(gomock.Any(),
			&wellarchitected.ListShareInvitationsInput{
				MaxResults:        50,
				ShareResourceType: shareResourceType,
			}, gomock.Any()).
			Return(
				&wellarchitected.ListShareInvitationsOutput{
					ShareInvitationSummaries: []types.ShareInvitationSummary{summary},
				},
				nil,
			)
	}
	return client.Services{Wellarchitected: m}
}

func TestShareInvitations(t *testing.T) {
	client.AwsMockTestHelper(t, ShareInvitations(), buildShareInvitationsMock, client.TestOptions{})
}
