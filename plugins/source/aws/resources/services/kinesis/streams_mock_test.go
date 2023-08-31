package kinesis

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type customKinesisClient struct {
	ConsumerCount *int32
	// EncryptionType types.EncryptionType
	// EnhancedMonitoring []EnhancedMetrics
	KeyId                   *string
	OpenShardCount          *int32
	RetentionPeriodHours    *int32
	StreamARN               *string
	StreamCreationTimestamp *time.Time
	StreamModeDetails       *types.StreamModeDetails
	StreamName              *string
	StreamStatus            types.StreamStatus
}

func buildKinesisStreams(t *testing.T, ctrl *gomock.Controller) client.Services {
	k := mocks.NewMockKinesisClient(ctrl)

	streams := kinesis.ListStreamsOutput{}
	require.NoError(t, faker.FakeObject(&streams))
	streams.HasMoreStreams = aws.Bool(false)
	streams.NextToken = nil
	k.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(&streams, nil)

	stream := kinesis.DescribeStreamSummaryOutput{
		StreamDescriptionSummary: &types.StreamDescriptionSummary{
			EnhancedMonitoring: []types.EnhancedMetrics{{
				ShardLevelMetrics: []types.MetricsName{types.MetricsNameAll},
			}}},
	}
	customKinesisClient := customKinesisClient{}
	require.NoError(t, faker.FakeObject(&customKinesisClient))

	stream.StreamDescriptionSummary.ConsumerCount = customKinesisClient.ConsumerCount
	stream.StreamDescriptionSummary.KeyId = customKinesisClient.KeyId
	stream.StreamDescriptionSummary.OpenShardCount = customKinesisClient.OpenShardCount
	stream.StreamDescriptionSummary.RetentionPeriodHours = customKinesisClient.RetentionPeriodHours
	stream.StreamDescriptionSummary.StreamARN = customKinesisClient.StreamARN
	stream.StreamDescriptionSummary.StreamCreationTimestamp = customKinesisClient.StreamCreationTimestamp
	stream.StreamDescriptionSummary.StreamModeDetails = customKinesisClient.StreamModeDetails
	stream.StreamDescriptionSummary.StreamName = customKinesisClient.StreamName
	stream.StreamDescriptionSummary.StreamStatus = customKinesisClient.StreamStatus
	k.EXPECT().DescribeStreamSummary(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&stream, nil)

	tags := kinesis.ListTagsForStreamOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.HasMoreTags = aws.Bool(false)
	k.EXPECT().ListTagsForStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&tags, nil)

	return client.Services{
		Kinesis: k,
	}
}

func TestStreams(t *testing.T) {
	client.AwsMockTestHelper(t, Streams(), buildKinesisStreams, client.TestOptions{})
}
