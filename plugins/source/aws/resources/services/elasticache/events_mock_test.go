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

func buildElasticacheEvents(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	event := types.Event{}
	require.NoError(t, faker.FakeObject(&event))

	mockElasticache.EXPECT().DescribeEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(&elasticache.DescribeEventsOutput{Events: []types.Event{event}}, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheEvents(t *testing.T) {
	client.AwsMockTestHelper(t, Events(), buildElasticacheEvents, client.TestOptions{})
}
