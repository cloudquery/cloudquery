//go:build mock
// +build mock

package elbv1

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv1Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElbv1LoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElbV1Client(ctrl)
	l := elbv1Types.LoadBalancerDescription{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancing.DescribeLoadBalancersOutput{
			LoadBalancerDescriptions: []elbv1Types.LoadBalancerDescription{l},
		}, nil)

	tag := elbv1Types.Tag{}
	err = faker.FakeData(&tag)
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
	err = faker.FakeData(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLoadBalancerAttributes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancing.DescribeLoadBalancerAttributesOutput{
			LoadBalancerAttributes: &a,
		}, nil)

	p := elbv1Types.PolicyDescription{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancerPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancing.DescribeLoadBalancerPoliciesOutput{
			PolicyDescriptions: []elbv1Types.PolicyDescription{p},
		}, nil)

	return client.Services{
		ELBv1: m,
	}
}

func TestElbv1LoadBalancers(t *testing.T) {
	client.AwsMockTestHelper(t, Elbv1LoadBalancers(), buildElbv1LoadBalancers, client.TestOptions{})
}
