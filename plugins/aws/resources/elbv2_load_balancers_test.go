package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElbv2LoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElbV2Client(ctrl)
	l := elbv2Types.LoadBalancer{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeLoadBalancersOutput{
			LoadBalancers: []elbv2Types.LoadBalancer{l},
		}, nil)

	m.EXPECT().DescribeLoadBalancerAttributes(
		gomock.Any(),
		&elasticloadbalancingv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: l.LoadBalancerArn},
		gomock.Any(),
	).Return(fakeLoadBalancerAttributes(), nil)

	return client.Services{
		ELBv2: m,
	}
}

func fakeLoadBalancerAttributes() *elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput {
	attr := func(key, value string) elbv2Types.LoadBalancerAttribute {
		return elbv2Types.LoadBalancerAttribute{Key: &key, Value: &value}
	}
	return &elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput{Attributes: []elbv2Types.LoadBalancerAttribute{
		attr("access_logs.s3.enabled", "true"),
		attr("access_logs.s3.bucket", "bucket"),
		attr("access_logs.s3.prefix", "prefix"),
		attr("deletion_protection.enabled", "true"),
		attr("idle_timeout.timeout_seconds", "10"),
		attr("routing.http.desync_mitigation_mode", "mode"),
		attr("routing.http.drop_invalid_header_fields.enabled", "true"),
		attr("routing.http.x_amzn_tls_version_and_cipher_suite.enabled", "true"),
		attr("routing.http.xff_client_port.enabled", "true"),
		attr("routing.http2.enabled", "true"),
		attr("waf.fail_open.enabled", "true"),
		attr("load_balancing.cross_zone.enabled", "true"),
	}}
}

func TestElbv2LoadBalancers(t *testing.T) {
	awsTestHelper(t, Elbv2LoadBalancers(), buildElbv2LoadBalancers, TestOptions{})
}
