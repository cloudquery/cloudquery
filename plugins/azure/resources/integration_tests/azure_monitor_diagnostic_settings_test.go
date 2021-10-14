package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationMonitorDiagnosticSettings(t *testing.T) {
	awsTestIntegrationHelper(t, resources.MonitorDiagnosticSettings(), []string{
		"azure_monitor_diagnostic_settings.tf",
		"azure_storage_accounts.tf",
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.MonitorDiagnosticSettings().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-ds", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"type": "Microsoft.Insights/diagnosticSettings",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_monitor_diagnostic_setting_logs",
					ForeignKeyName: "diagnostic_setting_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"retention_policy_enabled": false,
							"enabled":                  true,
							"retention_policy_days":    float64(1),
							"category":                 "VMProtectionAlerts",
						},
					}},
				},
				{
					Name:           "azure_monitor_diagnostic_setting_metrics",
					ForeignKeyName: "diagnostic_setting_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"retention_policy_enabled": false,
							"enabled":                  true,
							"retention_policy_days":    float64(0),
							"category":                 "AllMetrics",
						},
					}},
				},
			},
		}
	})
}
