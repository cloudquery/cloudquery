package glacier

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildVaultsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlacierClient(ctrl)

	v := glacier.ListVaultsOutput{}
	require.NoError(t, faker.FakeObject(&v))
	v.Marker = nil
	m.EXPECT().ListVaults(gomock.Any(), gomock.Any(), gomock.Any()).Return(&v, nil)

	ap := glacier.GetVaultAccessPolicyOutput{}
	require.NoError(t, faker.FakeObject(&ap))
	ap.Policy.Policy = aws.String(`{"some":"policy"}`)
	m.EXPECT().GetVaultAccessPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ap, nil)

	lp := glacier.GetVaultLockOutput{}
	require.NoError(t, faker.FakeObject(&lp))
	lp.Policy = aws.String(`{"some":"policy"}`)
	m.EXPECT().GetVaultLock(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lp, nil)

	vn := glacier.GetVaultNotificationsOutput{}
	require.NoError(t, faker.FakeObject(&vn))
	m.EXPECT().GetVaultNotifications(gomock.Any(), gomock.Any(), gomock.Any()).Return(&vn, nil)

	tags := glacier.ListTagsForVaultOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().ListTagsForVault(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{
		Glacier: m,
	}
}

func TestVaults(t *testing.T) {
	client.AwsMockTestHelper(t, Vaults(), buildVaultsMock, client.TestOptions{})
}
