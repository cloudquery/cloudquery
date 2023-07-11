package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEventBridgeEventBusesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)

	var bus types.EventBus
	require.NoError(t, faker.FakeObject(&bus))

	var rule types.Rule
	require.NoError(t, faker.FakeObject(&rule))

	var tags eventbridge.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&tags))

	var target types.Target
	require.NoError(t, faker.FakeObject(&target))

	m.EXPECT().ListEventBuses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListEventBusesOutput{
			EventBuses: []types.EventBus{bus},
		}, nil)
	m.EXPECT().ListRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListRulesOutput{
			Rules: []types.Rule{rule},
		}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(2).Return(
		&tags, nil)
	m.EXPECT().ListTargetsByRule(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListTargetsByRuleOutput{
			Targets: []types.Target{target},
		}, nil)

	return client.Services{
		Eventbridge: m,
	}
}

func TestEventBridgeEventBuses(t *testing.T) {
	client.AwsMockTestHelper(t, EventBuses(), buildEventBridgeEventBusesMock, client.TestOptions{})
}
