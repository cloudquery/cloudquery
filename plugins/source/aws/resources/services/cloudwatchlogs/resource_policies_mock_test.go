package cloudwatchlogs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildResourcePolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchlogsClient(ctrl)
	rp := types.ResourcePolicy{}
	require.NoError(t, faker.FakeObject(&rp))
	policyDocument := "{}"
	rp.PolicyDocument = &policyDocument
	m.EXPECT().DescribeResourcePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeResourcePoliciesOutput{
			ResourcePolicies: []types.ResourcePolicy{rp},
		}, nil)
	return client.Services{
		Cloudwatchlogs: m,
	}
}

func TestResourcePolicies(t *testing.T) {
	client.AwsMockTestHelper(t, ResourcePolicies(), buildResourcePolicies, client.TestOptions{})
}
