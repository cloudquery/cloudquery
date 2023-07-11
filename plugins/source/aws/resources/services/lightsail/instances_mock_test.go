package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var instances []types.Instance
	require.NoError(t, faker.FakeObject(&instances))

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
	require.NoError(t, faker.FakeObject(&p))

	mock.EXPECT().GetInstancePortStates(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&p, nil)

	var a lightsail.GetInstanceAccessDetailsOutput
	require.NoError(t, faker.FakeObject(&a))

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
