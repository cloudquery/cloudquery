package account

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/golang/mock/gomock"
)

func buildAccountUsersMock(t *testing.T, ctrl *gomock.Controller) services.FastlyClient {
	m := mocks.NewMockFastlyClient(ctrl)
	f := make([]*fastly.User, 0, 1)
	err := faker.FakeObject(&f)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetCurrentUser().Return(f[0], nil)
	m.EXPECT().ListCustomerUsers(gomock.Any()).Return(f, nil)
	return m
}

func TestAccountUsers(t *testing.T) {
	client.MockTestHelper(t, AccountUsers(), buildAccountUsersMock, client.TestOptions{})
}
