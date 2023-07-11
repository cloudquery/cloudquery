package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAppAssessments(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppAssessmentsOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListAppAssessments(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)

	var d resiliencehub.DescribeAppAssessmentOutput
	require.NoError(t, faker.FakeObject(&d))

	mock.EXPECT().DescribeAppAssessment(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&d, nil)

	buildAppComponentCompliances(t, mock)
	buildComponentRecommendations(t, mock)
	buildAppAlarmRecommendations(t, mock)
	buildAppTestRecommendations(t, mock)
	buildRecommendationsTemplates(t, mock)
	buildSopAlarmRecommendations(t, mock)
}
