package databases

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createBackups(t *testing.T, m *mocks.MockDatabasesService) {
	var data []godo.DatabaseBackup
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}
