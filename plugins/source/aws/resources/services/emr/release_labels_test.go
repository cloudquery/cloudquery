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

func buildReleaseLabels(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEmrClient(ctrl)

	releaseLabels := []string{"emr-6.0.0", "emr-5.0.0"}

	mock.EXPECT().ListReleaseLabels(gomock.Any(), &emr.ListReleaseLabelsInput{}, gomock.Any()).Return(
		&emr.ListReleaseLabelsOutput{ReleaseLabels: releaseLabels},
		nil,
	)

	for i := range releaseLabels {
		mock.EXPECT().DescribeReleaseLabel(gomock.Any(), &emr.DescribeReleaseLabelInput{ReleaseLabel: &releaseLabels[i]}, gomock.Any()).Return(
			&emr.DescribeReleaseLabelOutput{ReleaseLabel: &releaseLabels[i]},
			nil,
		)

		var instanceTypes []types.SupportedInstanceType
		require.NoError(t, faker.FakeObject(&instanceTypes))

		mock.EXPECT().ListSupportedInstanceTypes(gomock.Any(), &emr.ListSupportedInstanceTypesInput{ReleaseLabel: &releaseLabels[i]}, gomock.Any()).Return(
			&emr.ListSupportedInstanceTypesOutput{
				SupportedInstanceTypes: instanceTypes,
			},
			nil,
		)
	}

	return client.Services{Emr: mock}
}

func TestReleaseLabels(t *testing.T) {
	client.AwsMockTestHelper(t, ReleaseLabels(), buildReleaseLabels, client.TestOptions{})
}
