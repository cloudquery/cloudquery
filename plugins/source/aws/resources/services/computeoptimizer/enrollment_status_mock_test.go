package computeoptimizer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildEnrollmentStatuses(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockComputeoptimizerClient(ctrl)
	services := client.Services{
		Computeoptimizer: m,
	}
	item := computeoptimizer.GetEnrollmentStatusOutput{}
	err := faker.FakeObject(&item)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetEnrollmentStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(&item, nil)

	return services
}

func TestEnrollmentStatus(t *testing.T) {
	client.AwsMockTestHelper(t, EnrollmentStatuses(), buildEnrollmentStatuses, client.TestOptions{})
}
