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

func buildProducts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServicecatalogClient(ctrl)

	o := servicecatalog.SearchProductsAsAdminOutput{}
	require.NoError(t, faker.FakeObject(&o))

	o.NextPageToken = nil

	mk.EXPECT().SearchProductsAsAdmin(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)
	po := servicecatalog.DescribeProductAsAdminOutput{}
	require.NoError(t, faker.FakeObject(&po))

	mk.EXPECT().DescribeProductAsAdmin(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&po,
		nil,
	)

	return client.Services{
		Servicecatalog: mk,
	}
}

func TestProducts(t *testing.T) {
	client.AwsMockTestHelper(t, Products(), buildProducts, client.TestOptions{})
}
