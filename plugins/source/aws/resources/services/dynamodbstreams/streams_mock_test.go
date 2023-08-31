package dynamodbstreams

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDynamodbstreamsStreamsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbstreamsClient(ctrl)
	services := client.Services{
		Dynamodbstreams: m,
	}
	stream := types.Stream{}
	require.NoError(t, faker.FakeObject(&stream))

	m.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dynamodbstreams.ListStreamsOutput{
			Streams: []types.Stream{stream},
		},
		nil,
	)

	streamDescription := types.StreamDescription{}
	require.NoError(t, faker.FakeObject(&streamDescription))

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
