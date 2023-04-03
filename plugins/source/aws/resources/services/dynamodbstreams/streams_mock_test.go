package dynamodbstreams

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDynamodbstreamsStreamsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbstreamsClient(ctrl)
	services := client.Services{
		Dynamodbstreams: m,
	}
	stream := types.Stream{}
	if err := faker.FakeObject(&stream); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dynamodbstreams.ListStreamsOutput{
			Streams: []types.Stream{stream},
		},
		nil,
	)

	streamDescription := types.StreamDescription{}
	if err := faker.FakeObject(&streamDescription); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeStream(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dynamodbstreams.DescribeStreamOutput{
			StreamDescription: &streamDescription,
		},
		nil,
	)

	return services
}

func TestDynamodbStreams(t *testing.T) {
	client.AwsMockTestHelper(t, Streams(), buildDynamodbstreamsStreamsMock, client.TestOptions{})
}
