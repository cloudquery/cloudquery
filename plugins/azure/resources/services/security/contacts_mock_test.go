package security

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSecurityContactsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSecurityContactsClient(ctrl)

	var contact security.Contact
	if err := faker.FakeData(&contact); err != nil {
		t.Fatal(err)
	}

	result := security.NewContactListPage(
		security.ContactList{Value: &[]security.Contact{contact}},
		func(context.Context, security.ContactList) (security.ContactList, error) {
			return security.ContactList{}, nil
		},
	)
	m.EXPECT().List(gomock.Any()).Return(result, nil)
	return services.Services{
		Security: services.SecurityClient{Contacts: m},
	}
}

func TestSecurityContacts(t *testing.T) {
	client.AzureMockTestHelper(t, SecurityContacts(), buildSecurityContactsMock, client.TestOptions{})
}
