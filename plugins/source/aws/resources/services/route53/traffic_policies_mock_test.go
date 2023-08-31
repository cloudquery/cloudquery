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

func buildRoute53TrafficPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	tps := route53Types.TrafficPolicySummary{}
	require.NoError(t, faker.FakeObject(&tps))

	m.EXPECT().ListTrafficPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPoliciesOutput{
			TrafficPolicySummaries: []route53Types.TrafficPolicySummary{tps},
		}, nil)
	tp := route53Types.TrafficPolicy{}
	require.NoError(t, faker.FakeObject(&tp))

	tp.Id = tps.Id
	document := `{"AWSPolicyFormatVersion":"2015-10-01","RecordType":"A","Endpoints":{"endpoint-geoproximity-vfcf":{"Type":"value","Value":"1.0.0.1"},"endpoint-geoproximity-gPSy":{"Type":"value","Value":"1.0.0.2"}},"Rules":{"geoproximity-start-Hfni":{"RuleType":"geoproximity","GeoproximityLocations":[{"EndpointReference":"endpoint-geoproximity-vfcf","Bias":0,"Region":"aws:route53:us-east-1","EvaluateTargetHealth":true},{"Bias":0,"Region":"aws:route53:us-east-2","EvaluateTargetHealth":true,"EndpointReference":"endpoint-geoproximity-gPSy"}]}},"StartRule":"geoproximity-start-Hfni"}`
	tp.Document = &document
	m.EXPECT().ListTrafficPolicyVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPolicyVersionsOutput{
			TrafficPolicies: []route53Types.TrafficPolicy{tp},
		}, nil)
	return client.Services{
		Route53: m,
	}
}

func TestRoute53TrafficPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, TrafficPolicies(), buildRoute53TrafficPoliciesMock, client.TestOptions{})
}
