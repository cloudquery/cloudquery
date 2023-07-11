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

	mk.EXPECT().SearchProvisionedProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return(&o, nil)

	pao := servicecatalog.DescribeProvisioningArtifactOutput{}
	require.NoError(t, faker.FakeObject(&pao))
	mk.EXPECT().DescribeProvisioningArtifact(gomock.Any(), gomock.Any(), gomock.Any()).Return(&pao, nil)

	ppo := servicecatalog.DescribeProvisioningParametersOutput{}
	require.NoError(t, faker.FakeObject(&ppo))
	mk.EXPECT().DescribeProvisioningParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ppo, nil)

	llpo := servicecatalog.ListLaunchPathsOutput{}
	require.NoError(t, faker.FakeObject(&llpo))
	llpo.NextPageToken = nil
	mk.EXPECT().ListLaunchPaths(gomock.Any(), gomock.Any(), gomock.Any()).Return(&llpo, nil).MinTimes(1)

	return client.Services{
		Servicecatalog: mk,
	}
}

func TestProvisionedProducts(t *testing.T) {
	client.AwsMockTestHelper(t, ProvisionedProducts(), buildProvisionedProducts, client.TestOptions{})
}
