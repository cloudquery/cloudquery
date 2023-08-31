package account

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildContacts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAccountClient(ctrl)

	var ci types.ContactInformation
	require.NoError(t, faker.FakeObject(&ci))

	mock.EXPECT().GetContactInformation(
		gomock.Any(),
		&account.GetContactInformationInput{},
		gomock.Any(),
	).Return(
		&account.GetContactInformationOutput{ContactInformation: &ci},
		nil,
	)

	return client.Services{Account: mock}
}

func TestContacts(t *testing.T) {
	client.AwsMockTestHelper(t, Contacts(), buildContacts, client.TestOptions{})
}
