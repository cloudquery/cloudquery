package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRoute53HealthChecksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	hc := route53Types.HealthCheck{}
	if err := faker.FakeObject(&hc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListHealthChecks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListHealthChecksOutput{
			HealthChecks: []route53Types.HealthCheck{hc},
		}, nil)
	tag := route53Types.Tag{}
	if err := faker.FakeObject(&tag); err != nil {
		t.Fatal(err)
	}
	//m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
	//	&route53.ListTagsForResourceOutput{
	//		ResourceTagSet: &route53Types.ResourceTagSet{
	//			Tags: []route53Types.Tag{tag},
	//		},
	//	}, nil)
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
