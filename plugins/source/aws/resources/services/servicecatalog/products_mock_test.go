package servicecatalog

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildProducts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServiceCatalogClient(ctrl)
	ma := mocks.NewMockServiceCatalogAppRegistryClient(ctrl)

	o := servicecatalog.SearchProductsAsAdminOutput{}
	if err := faker.FakeObject(&o); err != nil {
		t.Fatal(err)
	}
	o.NextPageToken = nil

	mk.EXPECT().SearchProductsAsAdmin(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	to := servicecatalogappregistry.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}

	ma.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&to,
		nil,
	)

	return client.Services{
		ServiceCatalog:   mk,
		ServiceCatalogAR: ma,
	}
}

func TestProducts(t *testing.T) {
	client.AwsMockTestHelper(t, Products(), buildProducts, client.TestOptions{})
}
