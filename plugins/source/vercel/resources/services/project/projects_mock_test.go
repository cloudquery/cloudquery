package project

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildProjects(t *testing.T, ctrl *gomock.Controller) client.VercelServices {
	mock := mocks.NewMockVercelServices(ctrl)

	var p vercel.Project
	if err := faker.FakeObject(&p); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListProjects(gomock.Any(), gomock.Any()).Return([]vercel.Project{p}, &vercel.Paginator{}, nil)

	var e vercel.ProjectEnv
	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListProjectEnvs(gomock.Any(), p.ID, gomock.Any()).Return([]vercel.ProjectEnv{e}, &vercel.Paginator{}, nil)

	return mock
}

func TestProjects(t *testing.T) {
	client.MockTestHelper(t, Projects(), buildProjects)
}
