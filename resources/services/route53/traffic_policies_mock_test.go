// +build mock

package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRoute53TrafficPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	tps := route53Types.TrafficPolicySummary{}
	if err := faker.FakeData(&tps); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTrafficPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPoliciesOutput{
			TrafficPolicySummaries: []route53Types.TrafficPolicySummary{tps},
		}, nil)
	tp := route53Types.TrafficPolicy{}
	if err := faker.FakeData(&tp); err != nil {
		t.Fatal(err)
	}
	tp.Id = tps.Id
	jsonStr := "{\"test\": \"test\"}"
	tp.Document = &jsonStr
	m.EXPECT().ListTrafficPolicyVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPolicyVersionsOutput{
			TrafficPolicies: []route53Types.TrafficPolicy{tp},
		}, nil)
	return client.Services{
		Route53: m,
	}
}

func TestRoute53TrafficPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, Route53TrafficPolicies(), buildRoute53TrafficPoliciesMock, client.TestOptions{})
}
