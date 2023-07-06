package projects

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createResources(t *testing.T, m *mocks.MockProjectsService) {
	var data []godo.ProjectResource
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}
