package elastictranscoder

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildElastictranscoderPresetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElastictranscoderClient(ctrl)
	object := types.Preset{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListPresets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elastictranscoder.ListPresetsOutput{Presets: []types.Preset{object}},
		nil,
	)

	return client.Services{
		Elastictranscoder: m,
	}
}
func TestElastictranscoderPresets(t *testing.T) {
	client.AwsMockTestHelper(t, Presets(), buildElastictranscoderPresetsMock, client.TestOptions{})
}
