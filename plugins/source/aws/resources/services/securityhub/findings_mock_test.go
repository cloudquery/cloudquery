package securityhub

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildFindings(t *testing.T, ctrl *gomock.Controller) client.Services {
	shMock := mocks.NewMockSecurityhubClient(ctrl)
	findings := types.AwsSecurityFinding{}
	require.NoError(t, faker.FakeObject(&findings))
	findings.CreatedAt = aws.String(time.Now().Format(time.RFC3339))
	findings.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))
	findings.FirstObservedAt = aws.String(time.Now().Format(time.RFC3339))
	findings.LastObservedAt = aws.String(time.Now().Format(time.RFC3339))

	shMock.EXPECT().GetFindings(
		gomock.Any(),
		&securityhub.GetFindingsInput{
			MaxResults: 100,
		},
		gomock.Any(),
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
