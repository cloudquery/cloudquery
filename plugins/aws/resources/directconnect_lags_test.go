package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDirectconnectLag(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	lag := types.Lag{}
	err := faker.FakeData(&lag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeLagsOutput{
			Lags: []types.Lag{lag},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnectLag(t *testing.T) {
	awsTestHelper(t, DirectconnectLags(), buildDirectconnectLag, TestOptions{})
}
