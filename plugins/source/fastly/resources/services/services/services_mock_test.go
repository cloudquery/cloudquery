package services

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/golang/mock/gomock"
)

type mockListServicesPaginator struct {
	items       []*fastly.Service
	pages       int
	currentPage int
}

func (m *mockListServicesPaginator) Remaining() int {
	return m.pages - m.currentPage
}

func (m *mockListServicesPaginator) GetNext() ([]*fastly.Service, error) {
	m.currentPage++
	return m.items[m.currentPage-2 : m.currentPage-1], nil
}

func (m *mockListServicesPaginator) HasNext() bool {
	return m.currentPage <= m.pages
}

func buildServicesMock(t *testing.T, ctrl *gomock.Controller) services.FastlyClient {
	m := mocks.NewMockFastlyClient(ctrl)

	buildServiceVersionsMock(t, m)
	buildServiceBackendsMock(t, m)
	buildServiceDomainsMock(t, m)
	buildServiceHealthChecksMock(t, m)

	d := make([]*fastly.Service, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	p := &mockListServicesPaginator{
		items:       d,
		pages:       1,
		currentPage: 1,
	}
	m.EXPECT().NewListServicesPaginator(gomock.Any()).Return(p)
	return m
}

func TestServices(t *testing.T) {
	client.MockTestHelper(t, Services(), buildServicesMock, client.TestOptions{})
}
