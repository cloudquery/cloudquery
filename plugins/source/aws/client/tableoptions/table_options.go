package tableoptions

import "reflect"

type customInputValidation interface {
	Validate() error
}

type TableOptions struct {
	CloudwatchMetrics      CloudwatchMetrics       `json:"aws_alpha_cloudwatch_metrics,omitempty"`
	CloudTrailEvents       *CloudtrailAPIs         `json:"aws_cloudtrail_events,omitempty"`
	AccessAnalyzerFindings *AccessanalyzerFindings `json:"aws_accessanalyzer_analyzer_findings,omitempty"`
	Inspector2Findings     *Inspector2APIs         `json:"aws_inspector2_findings,omitempty"`
	CustomCostExplorer     *CostExplorerAPIs       `json:"aws_alpha_costexplorer_cost_custom,omitempty"`
	SecurityHubFindings    *SecurityHubAPIs        `json:"aws_securityhub_findings,omitempty"`
	ECSTasks               *ECSTaskAPIs            `json:"aws_ecs_cluster_tasks,omitempty"`
}

func (t TableOptions) Validate() error {
	v := reflect.ValueOf(t)
	for i := 0; i < v.NumField(); i++ {
		table := v.Field(i).Interface()
		if !reflect.ValueOf(table).IsNil() {
			tableInput := table.(customInputValidation)
			err := tableInput.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
