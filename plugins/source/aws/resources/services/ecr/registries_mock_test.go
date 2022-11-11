package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEcrRegistriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	var registryId string
	err := faker.FakeObject(&registryId)
	if err != nil {
		t.Fatal(err)
	}
	rcs := types.ReplicationConfiguration{}
	err = faker.FakeObject(&rcs)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRegistry(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeRegistryOutput{
			ReplicationConfiguration: &rcs,
			RegistryId:               aws.String(registryId),
		}, nil)

	return client.Services{
		Ecr: m,
	}
}

func TestEcrRegistries(t *testing.T) {
	client.AwsMockTestHelper(t, Registries(), buildEcrRegistriesMock, client.TestOptions{})
}
