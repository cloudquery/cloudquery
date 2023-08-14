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

func buildSysPlugins(t *testing.T, ctrl *gomock.Controller) *client.Services {
	sys := mocks.NewMockSys(ctrl)

	var response api.ListPluginsResponse
	err := faker.FakeObject(&response)
	require.NoError(t, err)

	sys.EXPECT().ListPluginsWithContext(gomock.Any(), gomock.Any()).Return(&response, nil)

	return &client.Services{Sys: sys}
}

func TestSysPlugins(t *testing.T) {
	client.MockTestHelper(t, VaultSysPlugins(), buildSysPlugins)
}
