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
	bus := types.EventBus{}
	err := faker.FakeObject(&bus)
	if err != nil {
		t.Fatal(err)
	}

	rule := types.Rule{}
	err = faker.FakeObject(&rule)
	if err != nil {
		t.Fatal(err)
	}

	tags := eventbridge.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
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
		Eventbridge: m,
	}
}

func TestEventBridgeEventBuses(t *testing.T) {
	client.AwsMockTestHelper(t, EventBuses(), buildEventBridgeEventBusesMock, client.TestOptions{})
}
