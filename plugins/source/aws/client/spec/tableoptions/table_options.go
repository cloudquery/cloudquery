package tableoptions

import (
	"reflect"
)

type (
	customInputValidation interface {
		Validate() error
	}
	defaultsSetter interface {
		SetDefaults()
	}
)

type TableOptions struct {
	AccessAnalyzerFindings *AccessAnalyzerFindings `json:"aws_accessanalyzer_analyzer_findings,omitempty"`
	CloudTrailEvents       *CloudtrailEvents       `json:"aws_cloudtrail_events,omitempty"`
	CloudwatchMetrics      CloudwatchMetrics       `json:"aws_alpha_cloudwatch_metrics,omitempty"`
	CustomCostExplorer     *CostExplorerAPIs       `json:"aws_alpha_costexplorer_cost_custom,omitempty"`
	ECSTasks               *ECSTasks               `json:"aws_ecs_cluster_tasks,omitempty"`
	Inspector2Findings     *Inspector2Findings     `json:"aws_inspector2_findings,omitempty"`
	SecurityHubFindings    *SecurityHubFindings    `json:"aws_securityhub_findings,omitempty"`
}

func (t *TableOptions) SetDefaults() {
	v := reflect.ValueOf(*t)
	for i := 0; i < v.NumField(); i++ {
		table := v.Field(i).Interface()
		if !reflect.ValueOf(table).IsNil() {
			tableInput, ok := table.(defaultsSetter)
			if ok {
				tableInput.SetDefaults()
			}
		}
	}
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
