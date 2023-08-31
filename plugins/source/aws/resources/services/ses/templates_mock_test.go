package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSESTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesv2Client(ctrl)

	tplMeta := types.EmailTemplateMetadata{}
	require.NoError(t, faker.FakeObject(&tplMeta))

	tpl := new(types.EmailTemplateContent)
	require.NoError(t, faker.FakeObject(tpl))

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
		Sesv2: sesClient,
	}
}

func TestSESTemplates(t *testing.T) {
	client.AwsMockTestHelper(t, Templates(), buildSESTemplates, client.TestOptions{})
}
