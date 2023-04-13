package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildEventbridgeReplaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)

	var object types.Replay
	if err := faker.FakeObject(&object); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListReplays(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListReplaysOutput{
			Replays: []types.Replay{object},
		}, nil)

	var desc eventbridge.DescribeReplayOutput
	if err := faker.FakeObject(&desc); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeReplay(gomock.Any(), gomock.Any(), gomock.Any()).Return(&desc, nil)

	var tagsOutput eventbridge.ListTagsForResourceOutput
	if err := faker.FakeObject(&tagsOutput); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Eventbridge: m,
	}
}
func TestEventbridgeReplays(t *testing.T) {
	client.AwsMockTestHelper(t, Replays(), buildEventbridgeReplaysMock, client.TestOptions{})
}
