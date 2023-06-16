package servicecatalog

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildProvisionedProducts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServicecatalogClient(ctrl)

	o := servicecatalog.SearchProvisionedProductsOutput{}
	require.NoError(t, faker.FakeObject(&o))

	o.NextPageToken = nil

	mk.EXPECT().SearchProvisionedProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	return client.Services{
		Servicecatalog: mk,
	}
}

func TestProvisionedProducts(t *testing.T) {
	client.AwsMockTestHelper(t, ProvisionedProducts(), buildProvisionedProducts, client.TestOptions{})
}
