package databases

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createReplicas(t *testing.T, m *mocks.MockDatabasesService) {
	var data []godo.DatabaseReplica
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListReplicas(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}
