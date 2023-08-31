package elbv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildTargetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticloadbalancingv2Client(ctrl)
	l := elbv2Types.TargetGroup{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticloadbalancingv2.DescribeTargetGroupsOutput{
			TargetGroups: []elbv2Types.TargetGroup{l},
		}, nil)

	tags := elasticloadbalancingv2.DescribeTagsOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	th := elasticloadbalancingv2.DescribeTargetHealthOutput{}
	require.NoError(t, faker.FakeObject(&th))
	m.EXPECT().DescribeTargetHealth(gomock.Any(), gomock.Any(), gomock.Any()).Return(&th, nil)
	return client.Services{
		Elasticloadbalancingv2: m,
	}
}

func TestElbv2TargetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, TargetGroups(), buildTargetGroups, client.TestOptions{})
}
