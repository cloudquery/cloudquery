package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildMlTransformsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var transforms glue.GetMLTransformsOutput
	require.NoError(t, faker.FakeData(&transforms))
	transforms.NextToken = nil
	m.EXPECT().GetMLTransforms(
		gomock.Any(),
		gomock.Any(),
	).Return(&transforms, nil)

	var runs glue.GetMLTaskRunsOutput
	require.NoError(t, faker.FakeData(&runs))
	runs.NextToken = nil
	m.EXPECT().GetMLTaskRuns(
		gomock.Any(),
		gomock.Any(),
	).Return(&runs, nil)

	m.EXPECT().GetTags(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return client.Services{
		Glue: m,
	}
}

func TestMlTransforms(t *testing.T) {
	client.AwsMockTestHelper(t, MlTransforms(), buildMlTransformsMock, client.TestOptions{})
}
