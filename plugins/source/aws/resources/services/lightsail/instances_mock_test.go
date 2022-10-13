package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var instances []types.Instance
	if err := faker.FakeObject(&instances); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetInstances(
		gomock.Any(),
		&lightsail.GetInstancesInput{},
		gomock.Any(),
	).Return(
		&lightsail.GetInstancesOutput{
			Instances: instances,
		},
		nil,
	)

	var p lightsail.GetInstancePortStatesOutput
	if err := faker.FakeObject(&p); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetInstancePortStates(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&p, nil)

	var a lightsail.GetInstanceAccessDetailsOutput
	if err := faker.FakeObject(&a); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetInstanceAccessDetails(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&a, nil)

	return client.Services{Lightsail: mock}
}

func TestLightsailInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildInstances, client.TestOptions{})
}
