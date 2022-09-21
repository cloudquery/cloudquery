package domains

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createRecords(t *testing.T, m *mocks.MockDomainsService) {
	var data []godo.DomainRecord
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().Records(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}
