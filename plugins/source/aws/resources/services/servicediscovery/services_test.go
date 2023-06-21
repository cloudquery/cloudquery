package servicediscovery

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockServicediscoveryClient(ctrl)

	var ss types.ServiceSummary
	require.NoError(t, faker.FakeObject(&ss))
	m.EXPECT().ListServices(
		gomock.Any(),
		&servicediscovery.ListServicesInput{MaxResults: aws.Int32(100)},
		gomock.Any(),
	).Return(
		&servicediscovery.ListServicesOutput{Services: []types.ServiceSummary{ss}},
		nil,
	)

	var service types.Service
	require.NoError(t, faker.FakeObject(&service))
	service.Arn = ss.Arn
	service.Id = ss.Id

	m.EXPECT().GetService(
		gomock.Any(),
		&servicediscovery.GetServiceInput{Id: ss.Id},
		gomock.Any(),
	).Return(
		&servicediscovery.GetServiceOutput{Service: &service},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&servicediscovery.ListTagsForResourceInput{ResourceARN: service.Arn},
		gomock.Any(),
	).Return(
		&servicediscovery.ListTagsForResourceOutput{Tags: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}}},
		nil,
	)

	return client.Services{
		Servicediscovery: m,
	}
}

func TestServices(t *testing.T) {
	client.AwsMockTestHelper(t, Services(), buildServices, client.TestOptions{})
}
