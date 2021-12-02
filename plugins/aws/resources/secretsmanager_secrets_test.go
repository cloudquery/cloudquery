package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	types "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSecretsmanagerModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSecretsManagerClient(ctrl)

	secret := types.SecretListEntry{}
	if err := faker.FakeData(&secret); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListSecrets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&secretsmanager.ListSecretsOutput{SecretList: []types.SecretListEntry{secret}},
		nil,
	)

	dsecret := secretsmanager.DescribeSecretOutput{}
	if err := faker.FakeData(&dsecret); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeSecret(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dsecret,
		nil,
	)

	var policy secretsmanager.GetResourcePolicyOutput
	if err := faker.FakeData(&policy); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetResourcePolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&policy,
		nil,
	)

	return client.Services{
		SecretsManager: m,
	}
}

func TestSecretsManagerModels(t *testing.T) {
	awsTestHelper(t, SecretsmanagerSecrets(), buildSecretsmanagerModels, TestOptions{})
}
