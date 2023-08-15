package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDirectconnectLag(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	lag := types.Lag{}
	require.NoError(t, faker.FakeObject(&lag))
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
