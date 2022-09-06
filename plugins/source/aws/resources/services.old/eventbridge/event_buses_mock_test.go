package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEventBridgeEventBusesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventBridgeClient(ctrl)
	bus := types.EventBus{}
	err := faker.FakeData(&bus)
	if err != nil {
		t.Fatal(err)
	}

	rule := types.Rule{}
	err = faker.FakeData(&rule)
	if err != nil {
		t.Fatal(err)
	}

	tags := eventbridge.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
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

	return client.Services{
		EventBridge: m,
	}
}

func TestEventBridgeEventBuses(t *testing.T) {
	client.AwsMockTestHelper(t, EventBuses(), buildEventBridgeEventBusesMock, client.TestOptions{})
}
