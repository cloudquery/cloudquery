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

func buildElasticacheReplicationGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeReplicationGroupsOutput{}
	err := faker.FakeObject(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeReplicationGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)

	var ta types.Tag
	if err = faker.FakeObject(&ta); err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().ListTagsForResource(gomock.Any(), &elasticache.ListTagsForResourceInput{ResourceName: output.ReplicationGroups[0].ARN}, gomock.Any()).Return(&elasticache.ListTagsForResourceOutput{TagList: []types.Tag{ta}}, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheReplicationGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ReplicationGroups(), buildElasticacheReplicationGroups, client.TestOptions{})
}
