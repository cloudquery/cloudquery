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

func buildElasticacheUpdateActions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	action := types.UpdateAction{}
	err := faker.FakeObject(&action)
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeUpdateActions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&elasticache.DescribeUpdateActionsOutput{UpdateActions: []types.UpdateAction{action}}, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheUpdateActions(t *testing.T) {
	client.AwsMockTestHelper(t, UpdateActions(), buildElasticacheUpdateActions, client.TestOptions{})
}
