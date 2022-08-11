package keyvault

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildKeyVaultManagedHSMMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockKeyVaultManagedHSMClient(ctrl)
	var vault keyvault.ManagedHsm
	if err := faker.FakeData(&vault); err != nil {
		t.Fatal(err)
	}
	var limit int32 = 100
	m.EXPECT().ListBySubscription(gomock.Any(), &limit).Return(
		keyvault.NewManagedHsmListResultPage(
			keyvault.ManagedHsmListResult{Value: &[]keyvault.ManagedHsm{vault}},
			func(c context.Context, mhlr keyvault.ManagedHsmListResult) (keyvault.ManagedHsmListResult, error) {
				return keyvault.ManagedHsmListResult{}, nil
			},
		),
		nil,
	)
	return services.Services{
		KeyVault: services.KeyVaultClient{
			ManagedHSM: m,
		},
	}
}

func TestKeyVaultManagedHSM(t *testing.T) {
	client.AzureMockTestHelper(t, KeyvaultManagedHSM(), buildKeyVaultManagedHSMMock, client.TestOptions{})
}
