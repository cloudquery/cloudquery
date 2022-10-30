package servicecatalog

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildProvisionedProducts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServiceCatalogClient(ctrl)

	o := servicecatalog.SearchProvisionedProductsOutput{}
	if err := faker.FakeObject(&o); err != nil {
		t.Fatal(err)
	}
	o.NextPageToken = nil

	mk.EXPECT().SearchProvisionedProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	return client.Services{
		ServiceCatalog: mk,
	}
}

func TestProvisionedProducts(t *testing.T) {
	client.AwsMockTestHelper(t, ProvisionedProducts(), buildProvisionedProducts, client.TestOptions{})
}
