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

func buildElbv2TargetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElbV2Client(ctrl)
	l := elbv2Types.TargetGroup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeTargetGroupsOutput{
			TargetGroups: []elbv2Types.TargetGroup{l},
		}, nil)

	tags := elasticloadbalancingv2.DescribeTagsOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)
	return client.Services{
		ELBv2: m,
	}
}

func TestElbv2TargetGroups(t *testing.T) {
	awsTestHelper(t, Elbv2TargetGroups(), buildElbv2TargetGroups, TestOptions{})
}
