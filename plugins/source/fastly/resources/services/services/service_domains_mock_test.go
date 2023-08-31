package services

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/golang/mock/gomock"
)

func buildServiceDomainsMock(t *testing.T, m *mocks.MockFastlyClient) {
	d := make([]*fastly.Domain, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListDomains(gomock.Any()).Return(d, nil)
}
