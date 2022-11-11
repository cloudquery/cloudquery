package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDirectconnectLag(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	lag := types.Lag{}
	err := faker.FakeObject(&lag)
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
	client.AwsMockTestHelper(t, Lags(), buildDirectconnectLag, client.TestOptions{})
}
