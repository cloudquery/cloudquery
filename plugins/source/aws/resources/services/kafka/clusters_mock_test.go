package kafka

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildKafkaClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKafkaClient(ctrl)
	object := types.Cluster{}
	require.NoError(t, faker.FakeObject(&object))
	buildKafkaNodesMock(t, m)
	buildKafkaClusterOperationsMock(t, m)

	m.EXPECT().ListClustersV2(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kafka.ListClustersV2Output{
			ClusterInfoList: []types.Cluster{object},
		}, nil)

	m.EXPECT().DescribeClusterV2(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kafka.DescribeClusterV2Output{
			ClusterInfo: &object,
		}, nil)

	tagsOutput := kafka.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()

	return client.Services{
		Kafka: m,
	}
}
func TestKafkaClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildKafkaClustersMock, client.TestOptions{})
}
