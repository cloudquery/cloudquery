package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildLoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var lb lightsail.GetLoadBalancersOutput
	if err := faker.FakeData(&lb); err != nil {
		t.Fatal(err)
	}
	lb.NextPageToken = nil

	mock.EXPECT().GetLoadBalancers(
		gomock.Any(),
		&lightsail.GetLoadBalancersInput{},
		gomock.Any(),
	).Return(&lb, nil)

	var lbc lightsail.GetLoadBalancerTlsCertificatesOutput
	if err := faker.FakeData(&lbc); err != nil {
		t.Fatal(err)
	}

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
