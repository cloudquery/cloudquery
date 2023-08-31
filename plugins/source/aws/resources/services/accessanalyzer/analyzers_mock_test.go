package accessanalyzer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAccessAnalyzer(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAccessanalyzerClient(ctrl)
	u := types.AnalyzerSummary{}
	require.NoError(t, faker.FakeObject(&u))

	f := types.FindingSummary{}
	require.NoError(t, faker.FakeObject(&f))

	m.EXPECT().ListAnalyzers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListAnalyzersOutput{
			Analyzers: []types.AnalyzerSummary{u},
		}, nil)

	m.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListFindingsOutput{
			Findings: []types.FindingSummary{f},
		}, nil)

	arch := types.ArchiveRuleSummary{}
	require.NoError(t, faker.FakeObject(&arch))

	m.EXPECT().ListArchiveRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListArchiveRulesOutput{
			ArchiveRules: []types.ArchiveRuleSummary{arch},
		}, nil)

	return client.Services{
		Accessanalyzer: m,
	}
}

func TestAccessAnalyzerAnalyzer(t *testing.T) {
	client.AwsMockTestHelper(t, Analyzers(), buildAccessAnalyzer, client.TestOptions{})
}
