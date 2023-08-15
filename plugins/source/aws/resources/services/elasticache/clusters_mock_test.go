package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildElasticacheClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeCacheClustersOutput{}
	require.NoError(t, faker.FakeObject(&output))
	output.Marker = nil

	var ta types.Tag
	require.NoError(t, faker.FakeObject(&ta))

	mockElasticache.EXPECT().DescribeCacheClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)
	mockElasticache.EXPECT().ListTagsForResource(gomock.Any(), &elasticache.ListTagsForResourceInput{ResourceName: output.CacheClusters[0].ARN}, gomock.Any()).Return(&elasticache.ListTagsForResourceOutput{TagList: []types.Tag{ta}}, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildElasticacheClusters, client.TestOptions{})
}
