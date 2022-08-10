package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWorkflowsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var name string
	require.NoError(t, faker.FakeData(&name))
	m.EXPECT().ListWorkflows(
		gomock.Any(),
		&glue.ListWorkflowsInput{MaxResults: aws.Int32(25)},
	).Return(
		&glue.ListWorkflowsOutput{Workflows: []string{name}},
		nil,
	)

	var w types.Workflow
	require.NoError(t, faker.FakeData(&w))
	w.Name = &name
	m.EXPECT().GetWorkflow(
		gomock.Any(),
		&glue.GetWorkflowInput{Name: aws.String(name)},
	).Return(
		&glue.GetWorkflowOutput{Workflow: &w},
		nil,
	)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: aws.String("arn:aws:glue:us-east-1:testAccount:workflow/" + name)},
	).Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return client.Services{
		Glue: m,
	}
}

func TestWorkflows(t *testing.T) {
	client.AwsMockTestHelper(t, Workflows(), buildWorkflowsMock, client.TestOptions{})
}
