package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildKmsAliases(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	aliases := kms.ListAliasesOutput{}
	require.NoError(t, faker.FakeObject(&aliases))
	aliases.NextMarker = nil
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&aliases, nil)

	return client.Services{
		Kms: m,
	}
}

func TestKmsAliases(t *testing.T) {
	client.AwsMockTestHelper(t, Aliases(), buildKmsAliases, client.TestOptions{})
}
