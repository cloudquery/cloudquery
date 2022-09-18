// Auto generated code - DO NOT EDIT.

package security

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
)

func TestSecurityContacts(t *testing.T) {
	client.MockTestHelper(t, Contacts(), createContactsMock)
}

func createContactsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityContactsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Contacts: mockClient,
		},
	}

	data := security.Contact{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.NewContactListPage(security.ContactList{Value: &[]security.Contact{data}}, func(ctx context.Context, result security.ContactList) (security.ContactList, error) {
		return security.ContactList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
