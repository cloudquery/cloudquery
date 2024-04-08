package hooks

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildHooks(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Hook
	require.NoError(t, faker.FakeObject(&cs))
	cs.Config = map[string]any{"key": "value"}
	cs.LastResponse = map[string]any{"key": "value"}
	mock.EXPECT().ListHooks(gomock.Any(), "testorg", gomock.Any()).Return([]*github.Hook{&cs}, &github.Response{}, nil)

	var hd *github.HookDelivery
	require.NoError(t, faker.FakeObject(&hd))
	hd.Request = nil
	hd.Response = nil
	mock.EXPECT().ListHookDeliveries(gomock.Any(), "testorg", *cs.ID, gomock.Any()).Return([]*github.HookDelivery{hd}, &github.Response{}, nil)

	var hdGet *github.HookDelivery
	require.NoError(t, faker.FakeObject(&hdGet))
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
