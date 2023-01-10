package elbv1

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv1Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildElbv1LoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticloadbalancingClient(ctrl)
	l := elbv1Types.LoadBalancerDescription{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancing.DescribeLoadBalancersOutput{
			LoadBalancerDescriptions: []elbv1Types.LoadBalancerDescription{l},
		}, nil)

	tag := elbv1Types.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancing.DescribeTagsOutput{
			TagDescriptions: []elbv1Types.TagDescription{
				{
					LoadBalancerName: l.LoadBalancerName,
					Tags:             []elbv1Types.Tag{tag},
				},
			},
		}, nil)

	a := elbv1Types.LoadBalancerAttributes{}
	err = faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLoadBalancerAttributes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancing.DescribeLoadBalancerAttributesOutput{
			LoadBalancerAttributes: &a,
		}, nil)

	p := elbv1Types.PolicyDescription{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancerPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancing.DescribeLoadBalancerPoliciesOutput{
			PolicyDescriptions: []elbv1Types.PolicyDescription{p},
		}, nil)

	return client.Services{
		Elasticloadbalancing: m,
	}
}

func TestElbv1LoadBalancers(t *testing.T) {
	client.AwsMockTestHelper(t, LoadBalancers(), buildElbv1LoadBalancers, client.TestOptions{})
}
