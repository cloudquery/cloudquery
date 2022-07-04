package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	sesTypes "github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSESTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSESClient(ctrl)

	tplMeta := sesTypes.TemplateMetadata{}
	err := faker.FakeData(&tplMeta)
	if err != nil {
		t.Fatal(err)
	}

	tpl := new(sesTypes.Template)
	err = faker.FakeData(tpl)
	if err != nil {
		t.Fatal(err)
	}

	tpl.TemplateName = aws.String(aws.ToString(tplMeta.Name))

	sesClient.EXPECT().ListTemplates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ses.ListTemplatesOutput{TemplatesMetadata: []sesTypes.TemplateMetadata{tplMeta}},
		nil,
	)
	sesClient.EXPECT().GetTemplate(gomock.Any(), &ses.GetTemplateInput{TemplateName: tplMeta.Name}, gomock.Any()).Return(
		&ses.GetTemplateOutput{Template: tpl}, nil,
	)

	return client.Services{
		SES: sesClient,
	}
}

func TestSESTemplates(t *testing.T) {
	client.AwsMockTestHelper(t, Templates(), buildSESTemplates, client.TestOptions{})
}
