package projects

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createProjects(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockProjectsService(ctrl)

	var data []godo.Project
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	createResources(t, m)

	return client.Services{
		Projects: m,
	}
}

func TestProjects(t *testing.T) {
	client.MockTestHelper(t, Projects(), createProjects, client.TestOptions{})
}
