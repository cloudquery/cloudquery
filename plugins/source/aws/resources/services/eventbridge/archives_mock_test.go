package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEventbridgeArchivesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)
	object := types.Archive{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListArchives(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListArchivesOutput{
			Archives: []types.Archive{object},
		}, nil)

	tagsOutput := eventbridge.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Eventbridge: m,
	}
}
func TestEventbridgeArchives(t *testing.T) {
	client.AwsMockTestHelper(t, Archives(), buildEventbridgeArchivesMock, client.TestOptions{})
}
