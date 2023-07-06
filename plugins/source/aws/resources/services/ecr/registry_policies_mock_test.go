package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEcrRegistryPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	var registryId string
	require.NoError(t, faker.FakeObject(&registryId))

	policyText := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": [
					"ecr:GetAuthorizationToken",
					"ecr:BatchCheckLayerAvailability",
					"ecr:GetDownloadUrlForLayer",
					"ecr:GetRepositoryPolicy",
					"ecr:DescribeRepositories",
					"ecr:ListImages",
					"ecr:DescribeImages",
					"ecr:BatchGetImage",
					"ecr:GetLifecyclePolicy",
					"ecr:GetLifecyclePolicyPreview",
					"ecr:ListTagsForResource",
					"ecr:DescribeImageScanFindings"
				],
				"Resource": "*"
			}
		]
	}`
	m.EXPECT().GetRegistryPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.GetRegistryPolicyOutput{
			PolicyText: aws.String(policyText),
			RegistryId: aws.String(registryId),
		}, nil)

	return client.Services{
		Ecr: m,
	}
}

func TestEcrRegistryPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, RegistryPolicies(), buildEcrRegistryPoliciesMock, client.TestOptions{})
}
