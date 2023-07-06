package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildStacksWithTemplate(tmpl string) func(t *testing.T, ctrl *gomock.Controller) client.Services {
	return func(t *testing.T, ctrl *gomock.Controller) client.Services {
		mock := mocks.NewMockCloudformationClient(ctrl)

		var stack types.Stack
		require.NoError(t, faker.FakeObject(&stack))
		mock.EXPECT().DescribeStacks(
			gomock.Any(),
			&cloudformation.DescribeStacksInput{},
			gomock.Any(),
		).Return(
			&cloudformation.DescribeStacksOutput{Stacks: []types.Stack{stack}},
			nil,
		)

		var resource types.StackResourceSummary
		require.NoError(t, faker.FakeObject(&resource))
		mock.EXPECT().ListStackResources(
			gomock.Any(),
			&cloudformation.ListStackResourcesInput{StackName: stack.StackName},
			gomock.Any(),
		).Return(
			&cloudformation.ListStackResourcesOutput{StackResourceSummaries: []types.StackResourceSummary{resource}},
			nil,
		)

		var summary cloudformation.GetTemplateSummaryOutput
		require.NoError(t, faker.FakeObject(&summary))
		summary.Metadata = aws.String(`{ "some": "metadata" }`) // Required as faker doesn't handle this.

		mock.EXPECT().GetTemplateSummary(
			gomock.Any(),
			&cloudformation.GetTemplateSummaryInput{StackName: stack.StackName},
			gomock.Any(),
		).Return(
			&summary,
			nil,
		)

		var template cloudformation.GetTemplateOutput
		require.NoError(t, faker.FakeObject(&template))
		template.TemplateBody = aws.String(tmpl) // Required as faker doesn't handle this.

		mock.EXPECT().GetTemplate(
			gomock.Any(),
			&cloudformation.GetTemplateInput{StackName: stack.StackName},
			gomock.Any(),
		).Return(
			&template,
			nil,
		)

		return client.Services{Cloudformation: mock}
	}
}

func TestCloudformationStacksWithYAML(t *testing.T) {
	// check that cloudformation templates are parsed correctly for both YAML and JSON
	client.AwsMockTestHelper(t, Stacks(), buildStacksWithTemplate(`---
AWSTemplateFormatVersion: "version date"`), client.TestOptions{})
}

func TestCloudformationStacksWithJSON(t *testing.T) {
	client.AwsMockTestHelper(t, Stacks(), buildStacksWithTemplate(`{"test": "key"}`), client.TestOptions{})
}
