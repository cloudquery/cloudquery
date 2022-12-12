package domain

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDomains(t *testing.T, ctrl *gomock.Controller) client.VercelServices {
	mock := mocks.NewMockVercelServices(ctrl)

	var d vercel.Domain
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListDomains(gomock.Any(), gomock.Any()).Return([]vercel.Domain{d}, &vercel.Paginator{}, nil)

	var r vercel.DomainRecord
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListDomainRecords(gomock.Any(), d.Name, gomock.Any()).Return([]vercel.DomainRecord{r}, &vercel.Paginator{}, nil)

	return mock
}

func TestDomains(t *testing.T) {
	client.MockTestHelper(t, Domains(), buildDomains)
}
