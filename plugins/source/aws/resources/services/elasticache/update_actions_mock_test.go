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

func buildElasticacheUpdateActions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	action := types.UpdateAction{}
	require.NoError(t, faker.FakeObject(&action))

	mockElasticache.EXPECT().DescribeUpdateActions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&elasticache.DescribeUpdateActionsOutput{UpdateActions: []types.UpdateAction{action}}, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheUpdateActions(t *testing.T) {
	client.AwsMockTestHelper(t, UpdateActions(), buildElasticacheUpdateActions, client.TestOptions{})
}
