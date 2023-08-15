package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildContainerServicesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	dep := types.ContainerServiceDeployment{State: "test", Containers: map[string]types.Container{"test": {Image: aws.String("test")}}}
	require.NoError(t, faker.FakeObject(&dep))
	service := types.ContainerService{CurrentDeployment: &dep, NextDeployment: &dep, Power: "test", ResourceType: "test", State: "test"}
	require.NoError(t, faker.FakeObject(&service))
	m.EXPECT().GetContainerServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lightsail.GetContainerServicesOutput{ContainerServices: []types.ContainerService{service}}, nil)

	m.EXPECT().GetContainerServiceDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lightsail.GetContainerServiceDeploymentsOutput{Deployments: []types.ContainerServiceDeployment{dep}}, nil)

	i := lightsail.GetContainerImagesOutput{}
	require.NoError(t, faker.FakeObject(&i))
	m.EXPECT().GetContainerImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(&i, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestContainerServices(t *testing.T) {
	client.AwsMockTestHelper(t, ContainerServices(), buildContainerServicesMock, client.TestOptions{})
}
