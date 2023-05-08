package client

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_options"
)

type TableOptions struct {
	CloudTrailEvents       *table_options.CloudtrailAPIs         `json:"aws_cloudtrail_events,omitempty"`
	AccessAnalyzerFindings *table_options.AccessanalyzerFindings `json:"aws_accessanalyzer_analyzer_findings,omitempty"`
	Inspector2Findings     *table_options.Inspector2APIs         `json:"aws_inspector2_findings,omitempty"`
}
