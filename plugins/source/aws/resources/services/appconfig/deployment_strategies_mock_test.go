package appconfig

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDeploymentStrategies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppconfigClient(ctrl)

	var ds types.DeploymentStrategy
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().ListDeploymentStrategies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appconfig.ListDeploymentStrategiesOutput{
			Items: []types.DeploymentStrategy{ds},
		},
		nil,
	)

	return client.Services{Appconfig: m}
}

func TestDeploymentStrategies(t *testing.T) {
	client.AwsMockTestHelper(t, DeploymentStrategies(), buildDeploymentStrategies, client.TestOptions{})
}
