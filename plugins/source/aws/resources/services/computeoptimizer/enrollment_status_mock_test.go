package computeoptimizer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEnrollmentStatuses(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := computeoptimizer.GetEnrollmentStatusOutput{}
	require.NoError(t, faker.FakeObject(&item))

	m.EXPECT().GetEnrollmentStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(&item, nil)

	return services
}

func TestEnrollmentStatus(t *testing.T) {
	client.AwsMockTestHelper(t, EnrollmentStatuses(), buildEnrollmentStatuses, client.TestOptions{})
}
