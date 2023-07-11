package traffic

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildReferrers(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)

	var referrers []*github.TrafficReferrer
	if err := faker.FakeObject(&referrers); err != nil {
		t.Fatal(err)
	}

	repositoriesMock.EXPECT().ListTrafficReferrers(gomock.Any(), "test string", "test string").Return(referrers, nil, nil)
	return client.GithubServices{Repositories: repositoriesMock}
}

func TestReferrers(t *testing.T) {
	client.GithubMockTestHelper(t, Referrers(), buildReferrers, client.TestOptions{})
}
