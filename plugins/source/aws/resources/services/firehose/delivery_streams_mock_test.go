package firehose

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildKinesisFirehoses(t *testing.T, ctrl *gomock.Controller) client.Services {
	f := mocks.NewMockFirehoseClient(ctrl)

	streams := firehose.ListDeliveryStreamsOutput{}
	require.NoError(t, faker.FakeObject(&streams))
	streams.HasMoreDeliveryStreams = aws.Bool(false)
	streams.DeliveryStreamNames = []string{"test-stream"}
	f.EXPECT().ListDeliveryStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(&streams, nil)

	stream := firehose.DescribeDeliveryStreamOutput{}

	require.NoError(t, faker.FakeObject(&stream))
	stream.DeliveryStreamDescription.Destinations = []types.DestinationDescription{stream.DeliveryStreamDescription.Destinations[0]}

	f.EXPECT().DescribeDeliveryStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&stream, nil)

	tags := firehose.ListTagsForDeliveryStreamOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.HasMoreTags = aws.Bool(false)
	f.EXPECT().ListTagsForDeliveryStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&tags, nil)

	return client.Services{
		Firehose: f,
	}
}

func TestFirehoses(t *testing.T) {
	client.AwsMockTestHelper(t, DeliveryStreams(), buildKinesisFirehoses, client.TestOptions{})
}
