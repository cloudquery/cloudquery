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

func buildSysAuths(t *testing.T, ctrl *gomock.Controller) *client.Services {
	sys := mocks.NewMockSys(ctrl)

	var auths map[string]*api.AuthMount
	err := faker.FakeObject(&auths)
	require.NoError(t, err)

	sys.EXPECT().ListAuthWithContext(gomock.Any()).Return(auths, nil)

	return &client.Services{Sys: sys}
}

func TestSysAuths(t *testing.T) {
	client.MockTestHelper(t, VaultSysAuths(), buildSysAuths)
}
