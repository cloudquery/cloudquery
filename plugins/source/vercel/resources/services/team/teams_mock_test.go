package team

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildTeams(t *testing.T, ctrl *gomock.Controller) client.VercelServices {
	mock := mocks.NewMockVercelServices(ctrl)

	var vt vercel.Team
	if err := faker.FakeObject(&vt); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListTeams(gomock.Any(), gomock.Any()).Return([]vercel.Team{vt}, &vercel.Paginator{}, nil)

	var m vercel.TeamMember
	if err := faker.FakeObject(&m); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListTeamMembers(gomock.Any(), vt.ID, gomock.Any()).Return([]vercel.TeamMember{m}, &vercel.Paginator{}, nil)

	return mock
}

func TestTeams(t *testing.T) {
	client.MockTestHelper(t, Teams(), buildTeams)
}
