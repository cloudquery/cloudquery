package codepipeline

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildPipelines(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCodepipelineClient(ctrl)

	var pipeSummary types.PipelineSummary
	if err := faker.FakeObject(&pipeSummary); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListPipelines(
		gomock.Any(),
		&codepipeline.ListPipelinesInput{},
		gomock.Any(),
	).Return(
		&codepipeline.ListPipelinesOutput{Pipelines: []types.PipelineSummary{pipeSummary}},
		nil,
	)

	var resource codepipeline.GetPipelineOutput
	if err := faker.FakeObject(&resource); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetPipeline(
		gomock.Any(),
		&codepipeline.GetPipelineInput{Name: pipeSummary.Name},
		gomock.Any(),
	).Return(
		&resource,
		nil,
	)

	tags := &codepipeline.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)

	return client.Services{Codepipeline: mock}
}

func TestCodePipelinePipelines(t *testing.T) {
	client.AwsMockTestHelper(t, Pipelines(), buildPipelines, client.TestOptions{})
}
