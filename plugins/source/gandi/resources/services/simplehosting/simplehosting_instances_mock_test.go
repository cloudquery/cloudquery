package simplehosting

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/go-gandi/go-gandi/simplehosting"
	"github.com/golang/mock/gomock"
)

func buildInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSimpleHostingClient(ctrl)

	var i simplehosting.Instance
	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListInstances().Return([]simplehosting.Instance{i}, nil)

	var v simplehosting.Vhost
	if err := faker.FakeObject(&v); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListVhosts(i.ID).Return([]simplehosting.Vhost{v}, nil)

	return client.Services{
		SimpleHostingClient: mock,
	}
}

func TestInstances(t *testing.T) {
	client.MockTestHelper(t, SimplehostingInstances(), buildInstances)
}
