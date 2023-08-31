package hooks

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildHooks(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Hook
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	cs.Config = make(map[string]any)
	cs.LastResponse = make(map[string]any)
	mock.EXPECT().ListHooks(gomock.Any(), "testorg", gomock.Any()).Return([]*github.Hook{&cs}, &github.Response{}, nil)

	var hd *github.HookDelivery
	if err := faker.FakeObject(&hd); err != nil {
		t.Fatal(err)
	}
	hd.Request = nil
	hd.Response = nil
	mock.EXPECT().ListHookDeliveries(gomock.Any(), "testorg", *cs.ID, gomock.Any()).Return([]*github.HookDelivery{hd}, &github.Response{}, nil)

	var hdGet *github.HookDelivery
	if err := faker.FakeObject(&hdGet); err != nil {
		t.Fatal(err)
	}
	rawRequest := json.RawMessage("{}")
	rawResponse := json.RawMessage("{}")
	hdGet.Request = &github.HookRequest{Headers: make(map[string]string), RawPayload: &rawRequest}
	hdGet.Response = &github.HookResponse{Headers: make(map[string]string), RawPayload: &rawResponse}
	mock.EXPECT().GetHookDelivery(gomock.Any(), "testorg", *cs.ID, *hd.ID).Return(hdGet, &github.Response{}, nil)

	return client.GithubServices{Organizations: mock}
}

func TestHooks(t *testing.T) {
	client.GithubMockTestHelper(t, Hooks(), buildHooks, client.TestOptions{})
}
