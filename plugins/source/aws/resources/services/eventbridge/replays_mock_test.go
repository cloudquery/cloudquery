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

func buildEventbridgeReplaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)

	var object types.Replay
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListReplays(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListReplaysOutput{
			Replays: []types.Replay{object},
		}, nil)

	var desc eventbridge.DescribeReplayOutput
	require.NoError(t, faker.FakeObject(&desc))

	m.EXPECT().DescribeReplay(gomock.Any(), gomock.Any(), gomock.Any()).Return(&desc, nil)

	var tagsOutput eventbridge.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&tagsOutput))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Eventbridge: m,
	}
}
func TestEventbridgeReplays(t *testing.T) {
	client.AwsMockTestHelper(t, Replays(), buildEventbridgeReplaysMock, client.TestOptions{})
}
