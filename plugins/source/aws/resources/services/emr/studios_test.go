package emr

import (
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func buildStudios(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEmrClient(ctrl)
	var summary types.StudioSummary
	require.NoError(t, faker.FakeObject(&summary))

	mock.EXPECT().ListStudios(gomock.Any(), &emr.ListStudiosInput{}, gomock.Any()).Return(
		&emr.ListStudiosOutput{Studios: []types.StudioSummary{summary}},
		nil,
	)

	var studio types.Studio
	require.NoError(t, faker.FakeObject(&studio))
	mock.EXPECT().DescribeStudio(gomock.Any(), &emr.DescribeStudioInput{StudioId: summary.StudioId}, gomock.Any()).Return(
		&emr.DescribeStudioOutput{Studio: &studio},
		nil,
	)

	return client.Services{Emr: mock}
}

func TestStudios(t *testing.T) {
	client.AwsMockTestHelper(t, Studios(), buildStudios, client.TestOptions{})
}
