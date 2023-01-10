package dax

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDAXClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDaxClient(ctrl)
	services := client.Services{
		Dax: m,
	}
	c := types.Cluster{}
	if err := faker.FakeObject(&c); err != nil {
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
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestDAXClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildDAXClustersMock, client.TestOptions{})
}
