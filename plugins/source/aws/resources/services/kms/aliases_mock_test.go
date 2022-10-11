package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildKmsAliases(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	aliases := kms.ListAliasesOutput{}
	err := faker.FakeData(&aliases)
	if err != nil {
		t.Fatal(err)
	}
	aliases.NextMarker = nil
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&aliases, nil)

	return client.Services{
		KMS: m,
	}
}

func TestKmsAliases(t *testing.T) {
	client.AwsMockTestHelper(t, Aliases(), buildKmsAliases, client.TestOptions{})
}
