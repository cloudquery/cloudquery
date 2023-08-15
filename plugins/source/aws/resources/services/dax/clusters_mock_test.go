package dax

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDAXClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDaxClient(ctrl)
	services := client.Services{
		Dax: m,
	}
	c := types.Cluster{}
	require.NoError(t, faker.FakeObject(&c))

	daxOutput := &dax.DescribeClustersOutput{
		Clusters: []types.Cluster{c},
	}
	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		daxOutput,
		nil,
	)

	tags := &dax.ListTagsOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	tags.NextToken = nil
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestDAXClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildDAXClustersMock, client.TestOptions{})
}
