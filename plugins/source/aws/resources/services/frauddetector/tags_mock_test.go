package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func addTagsCall(t *testing.T, client *mocks.MockFrauddetectorClient) {
	var data []types.Tag
	require.NoError(t, faker.FakeObject(&data))

	client.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.ListTagsForResourceOutput{Tags: data}, nil,
	)
}
