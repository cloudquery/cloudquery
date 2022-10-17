package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildNeptuneGlobalClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var gc types.GlobalCluster
	if err := faker.FakeObject(&gc); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeGlobalClusters(gomock.Any(), &neptune.DescribeGlobalClustersInput{}, gomock.Any()).Return(
		&neptune.DescribeGlobalClustersOutput{GlobalClusters: []types.GlobalCluster{gc}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: gc.GlobalClusterArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return client.Services{Neptune: mock}
}

func TestNeptuneGlobalCluster(t *testing.T) {
	client.AwsMockTestHelper(t, GlobalClusters(), buildNeptuneGlobalClusters, client.TestOptions{})
}
