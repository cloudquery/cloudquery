package elbv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildElbv2TargetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticloadbalancingv2Client(ctrl)
	l := elbv2Types.TargetGroup{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeTargetGroupsOutput{
			TargetGroups: []elbv2Types.TargetGroup{l},
		}, nil)

	tags := elasticloadbalancingv2.DescribeTagsOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	th := elasticloadbalancingv2.DescribeTargetHealthOutput{}
	err = faker.FakeObject(&th)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTargetHealth(gomock.Any(), gomock.Any(), gomock.Any()).Return(&th, nil)
	return client.Services{
		Elasticloadbalancingv2: m,
	}
}

func TestElbv2TargetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, TargetGroups(), buildElbv2TargetGroups, client.TestOptions{})
}
