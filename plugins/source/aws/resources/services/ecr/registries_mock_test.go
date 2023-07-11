package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEcrRegistriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	var registryId string
	require.NoError(t, faker.FakeObject(&registryId))

	rcs := types.ReplicationConfiguration{}
	require.NoError(t, faker.FakeObject(&rcs))

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
