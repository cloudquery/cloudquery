package secretsmanager

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSecretsmanagerModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSecretsmanagerClient(ctrl)

	secret := types.SecretListEntry{}
	require.NoError(t, faker.FakeObject(&secret))

	m.EXPECT().ListSecrets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&secretsmanager.ListSecretsOutput{SecretList: []types.SecretListEntry{secret}},
		nil,
	)

	dsecret := secretsmanager.DescribeSecretOutput{}
	require.NoError(t, faker.FakeObject(&dsecret))

	m.EXPECT().DescribeSecret(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dsecret,
		nil,
	)

	var policy secretsmanager.GetResourcePolicyOutput
	require.NoError(t, faker.FakeObject(&policy))

	p := `{"key":"value"}`
	policy.ResourcePolicy = &p
	m.EXPECT().GetResourcePolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&policy,
		nil,
	)

	version := secretsmanager.ListSecretVersionIdsOutput{}
	require.NoError(t, faker.FakeObject(&version))

	version.NextToken = nil
	m.EXPECT().ListSecretVersionIds(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&version,
		nil,
	)

	return client.Services{
		Secretsmanager: m,
	}
}

func TestSecretsManagerModels(t *testing.T) {
	client.AwsMockTestHelper(t, Secrets(), buildSecretsmanagerModels, client.TestOptions{})
}
