package firehose

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/bxcodec/faker"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/golang/mock/gomock"
)

func buildKinesisFirehoses(t *testing.T, ctrl *gomock.Controller) client.Services {
	f := mocks.NewMockFirehoseClient(ctrl)

	streams := firehose.ListDeliveryStreamsOutput{}
	err := faker.FakeData(&streams)
	if err != nil {
		t.Fatal(err)
	}
	streams.HasMoreDeliveryStreams = aws.Bool(false)
	streams.DeliveryStreamNames = []string{"test-stream"}
	f.EXPECT().ListDeliveryStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(&streams, nil)

	stream := firehose.DescribeDeliveryStreamOutput{}

	err = faker.FakeData(&stream)
	if err != nil {
		t.Fatal(err)
	}
	stream.DeliveryStreamDescription.Destinations = []types.DestinationDescription{stream.DeliveryStreamDescription.Destinations[0]}

	f.EXPECT().DescribeDeliveryStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&stream, nil)

	tags := firehose.ListTagsForDeliveryStreamOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.HasMoreTags = aws.Bool(false)
	f.EXPECT().ListTagsForDeliveryStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&tags, nil)

	return client.Services{
		Firehose: f,
	}
}

func TestFirehoses(t *testing.T) {
	client.AwsMockTestHelper(t, DeliveryStreams(), buildKinesisFirehoses, client.TestOptions{})
}
