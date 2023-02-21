package secretsmanager

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSecretsmanagerModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSecretsmanagerClient(ctrl)

	secret := types.SecretListEntry{}
	if err := faker.FakeObject(&secret); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListSecrets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&secretsmanager.ListSecretsOutput{SecretList: []types.SecretListEntry{secret}},
		nil,
	)

	dsecret := secretsmanager.DescribeSecretOutput{}
	if err := faker.FakeObject(&dsecret); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeSecret(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dsecret,
		nil,
	)

	var policy secretsmanager.GetResourcePolicyOutput
	if err := faker.FakeObject(&policy); err != nil {
		t.Fatal(err)
	}
	p := `{"key":"value"}`
	policy.ResourcePolicy = &p
	m.EXPECT().GetResourcePolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&policy,
		nil,
	)

	version := secretsmanager.ListSecretVersionIdsOutput{}
	if err := faker.FakeObject(&version); err != nil {
		t.Fatal(err)
	}
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
