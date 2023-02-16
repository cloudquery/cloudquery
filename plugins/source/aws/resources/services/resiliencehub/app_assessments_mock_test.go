package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAppAssessments(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppAssessmentsOutput
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	l.NextToken = nil
	mock.EXPECT().ListAppAssessments(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)

	var d resiliencehub.DescribeAppAssessmentOutput
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}

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
