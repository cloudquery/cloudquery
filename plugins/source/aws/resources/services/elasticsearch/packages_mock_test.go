package elasticsearch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildElasticSearchPackages(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticsearchserviceClient(ctrl)

	var pkg types.PackageDetails
	require.NoError(t, faker.FakeObject(&pkg))

	m.EXPECT().DescribePackages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.DescribePackagesOutput{
			PackageDetailsList: []types.PackageDetails{pkg},
		},
		nil,
	)

	return client.Services{Elasticsearchservice: m}
}

func TestElasticSearchPackages(t *testing.T) {
	client.AwsMockTestHelper(t, Packages(), buildElasticSearchPackages, client.TestOptions{})
}
