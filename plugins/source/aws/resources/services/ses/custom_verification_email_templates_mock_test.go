package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCustomVerificationEmailTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesv2Client(ctrl)

	metadata := types.CustomVerificationEmailTemplateMetadata{}
	if err := faker.FakeObject(&metadata); err != nil {
		t.Fatal(err)
	}

	get := new(sesv2.GetCustomVerificationEmailTemplateOutput)
	if err := faker.FakeObject(get); err != nil {
		t.Fatal(err)
	}
	metadata.TemplateName = get.TemplateName

	sesClient.EXPECT().ListCustomVerificationEmailTemplates(gomock.Any(), gomock.Any()).Return(
		&sesv2.ListCustomVerificationEmailTemplatesOutput{
			CustomVerificationEmailTemplates: []types.CustomVerificationEmailTemplateMetadata{metadata},
		},
		nil,
	)
	sesClient.EXPECT().GetCustomVerificationEmailTemplate(gomock.Any(), gomock.Any()).Return(
		get,
		nil,
	)

	return client.Services{
		Sesv2: sesClient,
	}
}

func TestCustomVerificationEmailTemplates(t *testing.T) {
	client.AwsMockTestHelper(t, CustomVerificationEmailTemplates(), buildCustomVerificationEmailTemplates, client.TestOptions{})
}
