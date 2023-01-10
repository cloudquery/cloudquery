package elbv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildElbv2LoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticloadbalancingv2Client(ctrl)
	w := mocks.NewMockWafv2Client(ctrl)
	l := elbv2Types.LoadBalancer{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.Type = elbv2Types.LoadBalancerTypeEnumApplication

	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeLoadBalancersOutput{
			LoadBalancers: []elbv2Types.LoadBalancer{l},
		}, nil)

	m.EXPECT().DescribeLoadBalancerAttributes(
		gomock.Any(),
		&elasticloadbalancingv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: l.LoadBalancerArn},
		gomock.Any(),
	).Return(fakeLoadBalancerAttributes(), nil)

	webAcl := types.WebACL{}
	err = faker.FakeObject(&webAcl)
	if err != nil {
		t.Fatal(err)
	}

	w.EXPECT().GetWebACLForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&wafv2.GetWebACLForResourceOutput{WebACL: &webAcl}, nil).AnyTimes()

	tags := elasticloadbalancingv2.DescribeTagsOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).Times(2).Return(&tags, nil)

	lis := elbv2Types.Listener{}
	if err := faker.FakeObject(&lis); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeListeners(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeListenersOutput{
			Listeners: []elbv2Types.Listener{lis},
		}, nil)

	c := elbv2Types.Certificate{}
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeListenerCertificates(
		gomock.Any(),
		&elasticloadbalancingv2.DescribeListenerCertificatesInput{ListenerArn: lis.ListenerArn},
		gomock.Any(),
	).Return(&elasticloadbalancingv2.DescribeListenerCertificatesOutput{
		Certificates: []elbv2Types.Certificate{c},
	}, nil)

	return client.Services{
		Elasticloadbalancingv2: m,
		Wafv2:                  w,
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
	client.AwsMockTestHelper(t, LoadBalancers(), buildElbv2LoadBalancers, client.TestOptions{})
}
