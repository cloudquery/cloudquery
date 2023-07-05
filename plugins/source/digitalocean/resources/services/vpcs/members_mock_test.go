package vpcs

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createMembers(t *testing.T, m *mocks.MockVpcsService) {
	var data []*godo.VPCMember
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}
