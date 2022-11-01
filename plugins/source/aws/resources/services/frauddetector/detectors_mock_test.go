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

func buildDetectors(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFraudDetectorClient(ctrl)

	data := types.Detector{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	fdClient.EXPECT().GetDetectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetDetectorsOutput{Detectors: []types.Detector{data}}, nil,
	)

	buildRules(t, fdClient)
	addTagsCall(t, fdClient)

	return client.Services{
		FraudDetector: fdClient,
	}
}

func TestDetectors(t *testing.T) {
	client.AwsMockTestHelper(t, Detectors(), buildDetectors, client.TestOptions{})
}
