package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildContainerServicesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	dep := types.ContainerServiceDeployment{State: "test", Containers: map[string]types.Container{"test": {Image: aws.String("test")}}}
	err := faker.FakeDataSkipFields(&dep, []string{"Containers", "State"})
	if err != nil {
		t.Fatal(err)
	}
	service := types.ContainerService{CurrentDeployment: &dep, NextDeployment: &dep, Power: "test", ResourceType: "test", State: "test"}
	err = faker.FakeDataSkipFields(&service, []string{"CurrentDeployment", "NextDeployment", "Power", "ResourceType", "State"})
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetContainerServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lightsail.GetContainerServicesOutput{ContainerServices: []types.ContainerService{service}}, nil)

	m.EXPECT().GetContainerServiceDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lightsail.GetContainerServiceDeploymentsOutput{Deployments: []types.ContainerServiceDeployment{dep}}, nil)

	i := lightsail.GetContainerImagesOutput{}
	err = faker.FakeData(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetContainerImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(&i, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestContainerServices(t *testing.T) {
	client.AwsMockTestHelper(t, ContainerServices(), buildContainerServicesMock, client.TestOptions{})
}
