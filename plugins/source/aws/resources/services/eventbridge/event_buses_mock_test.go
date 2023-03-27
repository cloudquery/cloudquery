package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEventBridgeEventBusesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)

	var bus types.EventBus
	if err := faker.FakeObject(&bus); err != nil {
		t.Fatal(err)
	}

	var rule types.Rule
	if err := faker.FakeObject(&rule); err != nil {
		t.Fatal(err)
	}

	var tags eventbridge.ListTagsForResourceOutput
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}

	var target types.Target
	if err := faker.FakeObject(&target); err != nil {
		t.Fatal(err)
	}

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
