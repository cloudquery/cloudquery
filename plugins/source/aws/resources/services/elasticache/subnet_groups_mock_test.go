package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildElasticacheSubnetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeCacheSubnetGroupsOutput{}
	require.NoError(t, faker.FakeObject(&output))
	output.Marker = nil

	mockElasticache.EXPECT().DescribeCacheSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheSubnetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SubnetGroups(), buildElasticacheSubnetGroups, client.TestOptions{})
}
