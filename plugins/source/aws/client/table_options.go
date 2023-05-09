package client

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_options"
)

type TableOptions struct {
	CloudTrailEvents       *table_options.CloudtrailAPIs         `json:"aws_cloudtrail_events,omitempty"`
	AccessAnalyzerFindings *table_options.AccessanalyzerFindings `json:"aws_accessanalyzer_analyzer_findings,omitempty"`
	Inspector2Findings     *table_options.Inspector2APIs         `json:"aws_inspector2_findings,omitempty"`
}

func (t TableOptions) Validate() error {
	if t.CloudTrailEvents != nil {
		err := t.CloudTrailEvents.Validate()
		if err != nil {
			return err
		}
	}

	if t.AccessAnalyzerFindings != nil {
		err := t.AccessAnalyzerFindings.Validate()
		if err != nil {
			return err
		}
	}

	if t.Inspector2Findings != nil {
		err := t.Inspector2Findings.Validate()
		if err != nil {
			return err
		}
	}
	
	return nil
}
