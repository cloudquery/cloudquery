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

func buildRules(t *testing.T, client *mocks.MockFrauddetectorClient) {
	data := types.RuleDetail{}
	require.NoError(t, faker.FakeObject(&data))

	client.EXPECT().GetRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetRulesOutput{RuleDetails: []types.RuleDetail{data}}, nil,
	)
}
