package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDAXClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDAXClient(ctrl)
	services := client.Services{
		DAX: m,
	}
	c := types.Cluster{}
	if err := faker.FakeData(&c); err != nil {
		t.Fatal(err)
	}
	daxOutput := &dax.DescribeClustersOutput{
		Clusters: []types.Cluster{c},
	}
	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		daxOutput,
		nil,
	)

	tags := &dax.ListTagsOutput{}
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestDAXClusters(t *testing.T) {
	awsTestHelper(t, DaxClusters(), buildDAXClustersMock, TestOptions{})
}
