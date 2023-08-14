package services

import (
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/vault/client"
	"github.com/cloudquery/plugins/vault/client/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/require"
	"testing"
)

func buildSysMounts(t *testing.T, ctrl *gomock.Controller) *client.Services {
	sys := mocks.NewMockSys(ctrl)

	var mounts map[string]*api.MountOutput
	err := faker.FakeObject(&mounts)
	require.NoError(t, err)

	sys.EXPECT().ListMountsWithContext(gomock.Any()).Return(mounts, nil)

	return &client.Services{Sys: sys}
}

func TestSysMounts(t *testing.T) {
	client.MockTestHelper(t, VaultSysMounts(), buildSysMounts)
}
