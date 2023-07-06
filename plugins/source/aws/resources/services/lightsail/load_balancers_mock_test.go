package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildLoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var lb lightsail.GetLoadBalancersOutput
	require.NoError(t, faker.FakeObject(&lb))

	lb.NextPageToken = nil

	mock.EXPECT().GetLoadBalancers(
		gomock.Any(),
		&lightsail.GetLoadBalancersInput{},
		gomock.Any(),
	).Return(&lb, nil)

	var lbc lightsail.GetLoadBalancerTlsCertificatesOutput
	require.NoError(t, faker.FakeObject(&lbc))

	mock.EXPECT().GetLoadBalancerTlsCertificates(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&lbc, nil)

	return client.Services{Lightsail: mock}
}

func TestLoadBalancers(t *testing.T) {
	client.AwsMockTestHelper(t, LoadBalancers(), buildLoadBalancers, client.TestOptions{})
}
