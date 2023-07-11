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

func buildCustomVerificationEmailTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesv2Client(ctrl)

	metadata := types.CustomVerificationEmailTemplateMetadata{}
	require.NoError(t, faker.FakeObject(&metadata))

	get := new(sesv2.GetCustomVerificationEmailTemplateOutput)
	require.NoError(t, faker.FakeObject(get))

	metadata.TemplateName = get.TemplateName

	sesClient.EXPECT().ListCustomVerificationEmailTemplates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sesv2.ListCustomVerificationEmailTemplatesOutput{
			CustomVerificationEmailTemplates: []types.CustomVerificationEmailTemplateMetadata{metadata},
		},
		nil,
	)
	sesClient.EXPECT().GetCustomVerificationEmailTemplate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
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
