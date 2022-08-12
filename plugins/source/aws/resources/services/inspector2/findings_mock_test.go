package inspector2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildInspectorV2Findings(t *testing.T, ctrl *gomock.Controller) client.Services {
	inspectorClient := mocks.NewMockInspectorV2Client(ctrl)

	finding := types.Finding{}
	err := faker.FakeData(&finding)
	if err != nil {
		t.Fatal(err)
	}

	inspectorClient.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&inspector2.ListFindingsOutput{Findings: []types.Finding{finding}},
		nil,
	)

	return client.Services{
		InspectorV2: inspectorClient,
	}
}

func TestInspectorV2Findings(t *testing.T) {
	client.AwsMockTestHelper(t, Findings(), buildInspectorV2Findings, client.TestOptions{})
}
