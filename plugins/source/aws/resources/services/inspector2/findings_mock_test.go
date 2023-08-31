package inspector2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildInspectorV2Findings(t *testing.T, ctrl *gomock.Controller) client.Services {
	inspectorClient := mocks.NewMockInspector2Client(ctrl)

	finding := types.Finding{}
	require.NoError(t, faker.FakeObject(&finding))

	inspectorClient.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&inspector2.ListFindingsOutput{Findings: []types.Finding{finding}},
		nil,
	)

	return client.Services{
		Inspector2: inspectorClient,
	}
}

func TestInspectorV2Findings(t *testing.T) {
	client.AwsMockTestHelper(t, Findings(), buildInspectorV2Findings, client.TestOptions{})
}
