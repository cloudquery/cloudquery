package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildJobsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	node := types.CodeGenConfigurationNode{}

	require.NoError(t, faker.FakeObject(&node))
	job := types.Job{
		CodeGenConfigurationNodes: map[string]types.CodeGenConfigurationNode{"test": node},
		ExecutionClass:            types.ExecutionClassFlex,
	}
	require.NoError(t, faker.FakeObject(&job))
	m.EXPECT().GetJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(&glue.GetJobsOutput{Jobs: []types.Job{job}}, nil)

	var jobRuns glue.GetJobRunsOutput
	require.NoError(t, faker.FakeObject(&jobRuns))
	jobRuns.NextToken = nil
	m.EXPECT().GetJobRuns(gomock.Any(), gomock.Any(), gomock.Any()).Return(&jobRuns, nil)

	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return client.Services{
		Glue: m,
	}
}

func TestJobs(t *testing.T) {
	client.AwsMockTestHelper(t, Jobs(), buildJobsMock, client.TestOptions{})
}
