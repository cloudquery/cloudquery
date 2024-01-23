package access_applications

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildAccessApplications(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var accessApplication cloudflare.AccessApplication
	if err := faker.FakeObject(&accessApplication); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListAccessApplications(
		gomock.Any(),
		cloudflare.ZoneIdentifier(client.TestZoneID),
		cloudflare.ListAccessApplicationsParams{
			ResultInfo: cloudflare.ResultInfo{
				Page:    1,
				PerPage: 200,
			},
		},
	).Return(
		[]cloudflare.AccessApplication{accessApplication},
		&cloudflare.ResultInfo{
			Page:    1,
			PerPage: 1,
			Count:   1,
			Total:   1,
		},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestAccessApplications(t *testing.T) {
	client.MockTestHelper(t, AccessApplications(), buildAccessApplications)
}
