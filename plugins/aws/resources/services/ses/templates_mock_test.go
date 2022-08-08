package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSESTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSESClient(ctrl)

	tplMeta := types.EmailTemplateMetadata{}
	err := faker.FakeData(&tplMeta)
	if err != nil {
		t.Fatal(err)
	}

	tpl := new(types.EmailTemplateContent)
	err = faker.FakeData(tpl)
	if err != nil {
		t.Fatal(err)
	}

	sesClient.EXPECT().ListEmailTemplates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sesv2.ListEmailTemplatesOutput{TemplatesMetadata: []types.EmailTemplateMetadata{tplMeta}},
		nil,
	)
	sesClient.EXPECT().GetEmailTemplate(gomock.Any(), &sesv2.GetEmailTemplateInput{TemplateName: tplMeta.TemplateName}, gomock.Any()).Return(
		&sesv2.GetEmailTemplateOutput{
			TemplateName:    tplMeta.TemplateName,
			TemplateContent: tpl,
		}, nil,
	)

	return client.Services{
		SES: sesClient,
	}
}

func TestSESTemplates(t *testing.T) {
	client.AwsMockTestHelper(t, Templates(), buildSESTemplates, client.TestOptions{})
}
