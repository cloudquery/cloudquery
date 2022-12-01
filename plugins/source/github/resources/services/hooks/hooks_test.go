package hooks

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
)

func buildHooks(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Hook
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	cs.Config = make(map[string]interface{})
	cs.LastResponse = make(map[string]interface{})
	mock.EXPECT().ListHooks(gomock.Any(), "testorg", gomock.Any()).Return([]*github.Hook{&cs}, &github.Response{}, nil)

	var hd *github.HookDelivery
	if err := faker.FakeObject(&hd); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListHookDeliveries(gomock.Any(), "testorg", *cs.ID, gomock.Any()).Return([]*github.HookDelivery{hd}, &github.Response{}, nil)
	return client.GithubServices{Organizations: mock}
}

func TestHooks(t *testing.T) {
	client.GithubMockTestHelper(t, Hooks(), buildHooks, client.TestOptions{})
}
