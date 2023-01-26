package securityhub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildFindings(t *testing.T, ctrl *gomock.Controller) client.Services {
	shMock := mocks.NewMockSecurityhubClient(ctrl)
	findings := types.AwsSecurityFinding{}
	err := faker.FakeObject(&findings)
	if err != nil {
		t.Fatal(err)
	}

	shMock.EXPECT().GetFindings(
		gomock.Any(),
		&securityhub.GetFindingsInput{
			MaxResults: 100,
		},
	).Return(
		&securityhub.GetFindingsOutput{
			Findings: []types.AwsSecurityFinding{findings},
		},
		nil,
	)

	return client.Services{Securityhub: shMock}
}

func TestFindings(t *testing.T) {
	client.AwsMockTestHelper(t, Findings(), buildFindings, client.TestOptions{})
}
