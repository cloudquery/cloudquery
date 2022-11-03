package codepipeline

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildWebhooks(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCodepipelineClient(ctrl)

	var webhook types.ListWebhookItem
	if err := faker.FakeObject(&webhook); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWebhooks(
		gomock.Any(),
		&codepipeline.ListWebhooksInput{},
		gomock.Any(),
	).Return(
		&codepipeline.ListWebhooksOutput{Webhooks: []types.ListWebhookItem{webhook}},
		nil,
	)

	return client.Services{Codepipeline: mock}
}

func TestCodePipelineWebhooks(t *testing.T) {
	client.AwsMockTestHelper(t, Webhooks(), buildWebhooks, client.TestOptions{})
}
