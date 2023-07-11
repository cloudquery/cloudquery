package codebuild

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSourceCredentials(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodebuildClient(ctrl)

	sourceCredentials := codebuild.ListSourceCredentialsOutput{}
	require.NoError(t, faker.FakeObject(&sourceCredentials))

	m.EXPECT().ListSourceCredentials(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&sourceCredentials,
		nil,
	)
	return client.Services{Codebuild: m}
}

func TestSourceCredentials(t *testing.T) {
	client.AwsMockTestHelper(t, SourceCredentials(), buildSourceCredentials, client.TestOptions{})
}
