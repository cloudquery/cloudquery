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

func buildElasticacheUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeUsersOutput{}
	require.NoError(t, faker.FakeObject(&output))
	output.Marker = nil

	mockElasticache.EXPECT().DescribeUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheUsers(t *testing.T) {
	client.AwsMockTestHelper(t, Users(), buildElasticacheUsers, client.TestOptions{})
}
