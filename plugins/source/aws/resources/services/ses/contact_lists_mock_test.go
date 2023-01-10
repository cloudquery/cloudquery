package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildContactLists(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesv2Client(ctrl)

	cs := sesv2.GetContactListOutput{}
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}

	sesClient.EXPECT().ListContactLists(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sesv2.ListContactListsOutput{ContactLists: []types.ContactList{{ContactListName: cs.ContactListName}}},
		nil,
	)
	sesClient.EXPECT().GetContactList(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cs,
		nil,
	)

	return client.Services{
		Sesv2: sesClient,
	}
}

func TestContactLists(t *testing.T) {
	client.AwsMockTestHelper(t, ContactLists(), buildContactLists, client.TestOptions{})
}
