package route53recoveryreadiness

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildReadinessChecks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoveryreadinessClient(ctrl)
	rco := types.ReadinessCheckOutput{}
	require.NoError(t, faker.FakeObject(&rco))

	m.EXPECT().ListReadinessChecks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53recoveryreadiness.ListReadinessChecksOutput{
			ReadinessChecks: []types.ReadinessCheckOutput{rco},
		}, nil)

	return client.Services{
		Route53recoveryreadiness: m,
	}
}

func TestReadinessChecks(t *testing.T) {
	client.AwsMockTestHelper(t, ReadinessChecks(), buildReadinessChecks, client.TestOptions{Region: "us-west-2"})
}
