package tableoptions

import (
	"reflect"
)

type (
	customInputValidation interface {
		Validate() error
	}
)

// TableOptions allows users to override the default options for specific tables.
type TableOptions struct {
	// Override options for `aws_accessanalyzer_analyzer_findings` table.
	AccessAnalyzerFindings *AccessAnalyzerFindings `json:"aws_accessanalyzer_analyzer_findings,omitempty"`

	// Override options for `aws_cloudtrail_events` table.
	CloudTrailEvents *CloudtrailEvents `json:"aws_cloudtrail_events,omitempty"`

	// Override options for `aws_alpha_cloudwatch_metrics` table.
	CloudwatchMetrics CloudwatchMetrics `json:"aws_alpha_cloudwatch_metrics,omitempty"`

	// Override options for `aws_alpha_costexplorer_cost_custom` table.
	CustomCostExplorer *CostExplorerAPIs `json:"aws_alpha_costexplorer_cost_custom,omitempty"`

	// Override options for `aws_ecs_cluster_tasks` table.
	ECSTasks *ECSTasks `json:"aws_ecs_cluster_tasks,omitempty"`

	// Override options for `aws_inspector2_findings` table.
	Inspector2Findings *Inspector2Findings `json:"aws_inspector2_findings,omitempty"`

	// Override options for `aws_securityhub_findings` table.
	SecurityHubFindings *SecurityHubFindings `json:"aws_securityhub_findings,omitempty"`
}

func (t *TableOptions) Validate() error {
	v := reflect.ValueOf(*t)
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
