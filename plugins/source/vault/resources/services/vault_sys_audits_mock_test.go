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

func buildSysAudits(t *testing.T, ctrl *gomock.Controller) *client.Services {
	sys := mocks.NewMockSys(ctrl)

	var audits map[string]*api.Audit
	err := faker.FakeObject(&audits)
	require.NoError(t, err)

	sys.EXPECT().ListAuditWithContext(gomock.Any()).Return(audits, nil)

	return &client.Services{Sys: sys}
}

func TestSysAudits(t *testing.T) {
	client.MockTestHelper(t, VaultSysAudits(), buildSysAudits)
}
