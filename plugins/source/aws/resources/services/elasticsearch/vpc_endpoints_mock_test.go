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

func buildElasticSearchVpcEndpoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticsearchserviceClient(ctrl)

	var summary types.VpcEndpointSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListVpcEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.ListVpcEndpointsOutput{
			VpcEndpointSummaryList: []types.VpcEndpointSummary{summary},
		},
		nil,
	)

	var endpoint types.VpcEndpoint
	require.NoError(t, faker.FakeObject(&endpoint))

	endpoint.VpcEndpointId = summary.VpcEndpointId

	m.EXPECT().DescribeVpcEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.DescribeVpcEndpointsOutput{
			VpcEndpoints: []types.VpcEndpoint{endpoint},
		},
		nil,
	)

	return client.Services{Elasticsearchservice: m}
}

func TestElasticSearchVpcEndpoints(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpoints(), buildElasticSearchVpcEndpoints, client.TestOptions{})
}
