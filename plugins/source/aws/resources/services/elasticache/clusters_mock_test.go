package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildElasticacheClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeCacheClustersOutput{}
	err := faker.FakeObject(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	var ta types.Tag
	if err = faker.FakeObject(&ta); err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeCacheClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)
	mockElasticache.EXPECT().ListTagsForResource(gomock.Any(), &elasticache.ListTagsForResourceInput{ResourceName: output.CacheClusters[0].ARN}, gomock.Any()).Return(&elasticache.ListTagsForResourceOutput{TagList: []types.Tag{ta}}, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildElasticacheClusters, client.TestOptions{})
}
