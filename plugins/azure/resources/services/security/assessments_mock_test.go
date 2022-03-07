package security

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSecurityAssessments(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockAssessmentsClient(ctrl)
	var p security.AssessmentProperties
	if err := faker.FakeDataSkipFields(&p, []string{"ResourceDetails"}); err != nil {
		t.Fatal(err)
	}
	p.ResourceDetails = security.ResourceDetails{Source: security.SourceResourceDetails}
	var a security.Assessment
	if err := faker.FakeDataSkipFields(&a, []string{"AssessmentProperties"}); err != nil {
		t.Fatal(err)
	}
	a.AssessmentProperties = &p
	m.EXPECT().List(gomock.Any(), "/subscriptions/test_sub").Return(
		security.NewAssessmentListPage(
			security.AssessmentList{Value: &[]security.Assessment{a}},
			func(c context.Context, al security.AssessmentList) (security.AssessmentList, error) {
				return security.AssessmentList{}, nil
			},
		),
		nil,
	)
	return services.Services{
		Security: services.SecurityClient{
			Assessments: m,
		},
	}
}

func TestSecurityAssessments(t *testing.T) {
	client.AzureMockTestHelper(t, SecurityAssessments(), buildSecurityAssessments, client.TestOptions{})
}
