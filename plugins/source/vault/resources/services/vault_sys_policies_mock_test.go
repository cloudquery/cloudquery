package services

import (
	"github.com/cloudquery/plugins/vault/client"
	"github.com/cloudquery/plugins/vault/client/services/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildSysPolicies(_ *testing.T, ctrl *gomock.Controller) *client.Services {
	sys := mocks.NewMockSys(ctrl)

	policies := []string{"policy1", "policy2", "policy3"}

	sys.EXPECT().ListPoliciesWithContext(gomock.Any()).Return(policies, nil)

	return &client.Services{Sys: sys}
}

func TestSysPolicies(t *testing.T) {
	client.MockTestHelper(t, VaultSysPolicies(), buildSysPolicies)
}
