package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFraudDetectorClient(ctrl)

	data := types.RuleDetail{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	fdClient.EXPECT().GetRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetRulesOutput{RuleDetails: []types.RuleDetail{data}}, nil,
	)

	return client.Services{
		FraudDetector: fdClient,
	}
}

func TestRules(t *testing.T) {
	client.AwsMockTestHelper(t, Rules(), buildRules, client.TestOptions{})
}
