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

func buildLabels(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFraudDetectorClient(ctrl)

	data := types.Label{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	fdClient.EXPECT().GetLabels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetLabelsOutput{Labels: []types.Label{data}}, nil,
	)

	addTagsCall(t, fdClient)

	return client.Services{
		FraudDetector: fdClient,
	}
}

func TestLabels(t *testing.T) {
	client.AwsMockTestHelper(t, Labels(), buildLabels, client.TestOptions{})
}
