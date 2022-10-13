package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEcrRegistriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	var registryId string
	err := faker.FakeData(&registryId)
	if err != nil {
		t.Fatal(err)
	}
	rcs := types.ReplicationConfiguration{}
	err = faker.FakeData(&rcs)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRegistry(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeRegistryOutput{
			ReplicationConfiguration: &rcs,
			RegistryId:               aws.String(registryId),
		}, nil)

	return client.Services{
		ECR: m,
	}
}

func TestEcrRegistries(t *testing.T) {
	client.AwsMockTestHelper(t, Registries(), buildEcrRegistriesMock, client.TestOptions{})
}
