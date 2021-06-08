package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAccessAnalyzer(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAnalyzerClient(ctrl)
	u := types.AnalyzerSummary{}
	if err := faker.FakeData(&u); err != nil {
		t.Fatal(err)
	}
	f := types.FindingSummary{}
	if err := faker.FakeData(&f); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAnalyzers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListAnalyzersOutput{
			Analyzers: []types.AnalyzerSummary{u},
		}, nil)

	m.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListFindingsOutput{
			Findings: []types.FindingSummary{f},
		}, nil)

	return client.Services{
		Analyzer: m,
	}
}

func TestAccessAnalyzerAnalyzer(t *testing.T) {
	awsTestHelper(t, AccessAnalyzerAnalyzer(), buildAccessAnalyzer, TestOptions{})
}
