package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRoute53HealthChecksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	hc := route53Types.HealthCheck{}
	require.NoError(t, faker.FakeObject(&hc))

	m.EXPECT().ListHealthChecks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListHealthChecksOutput{
			HealthChecks: []route53Types.HealthCheck{hc},
		}, nil)
	tag := route53Types.Tag{}
	require.NoError(t, faker.FakeObject(&tag))

	m.EXPECT().ListTagsForResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTagsForResourcesOutput{
			ResourceTagSets: []route53Types.ResourceTagSet{
				{
					ResourceId: hc.Id,
					Tags:       []route53Types.Tag{tag},
				},
			},
		}, nil)
	return client.Services{
		Route53: m,
	}
}

func TestRoute53HealthCheck(t *testing.T) {
	client.AwsMockTestHelper(t, HealthChecks(), buildRoute53HealthChecksMock, client.TestOptions{})
}
