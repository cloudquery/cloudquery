package plugin

import (
	"strings"
	"testing"

	"github.com/gertd/go-pluralize"
)

func TestAWS(t *testing.T) {
	// Note: this test is simple, but serves as a smoke test.
	// The AWS() call below also catches duplicate columns and other issues
	// that may have been missed if mock tests are incomplete.
	p := AWS()
	name := p.Name()
	if name != "aws" {
		t.Errorf("Name() = %q, want %q", name, "aws")
	}
}

// This test ensures that all tables have proper name and description.
func TestAWSTables(t *testing.T) {
	pluralize := pluralize.NewClient()
	descriptions := make(map[string]string)
	tables := getTables().FlattenTables()
	for _, table := range tables {
		if !ignorePluralName(table.Name) && !pluralize.IsPlural(table.Name[strings.LastIndex(table.Name, ",")+1:]) {
			t.Errorf("invalid table name: %s. must be plural.", table.Name)
		}
		if ignoreTableDescription(table.Name) {
			continue
		}
		if val, ok := descriptions[table.Description]; ok || table.Description == "" {
			t.Errorf("duplicate description for %s and %s", val, table.Name)
		} else {
			descriptions[table.Description] = table.Name
		}
	}
}

func ignoreTableDescription(tableName string) bool {
	tablesToIgnore := map[string]bool{
		"aws_resiliencehub_suggested_resiliency_policies":       true,
		"aws_rds_cluster_parameter_group_parameters":            true,
		"aws_rds_db_parameter_group_db_parameters":              true,
		"aws_rds_cluster_parameters":                            true,
		"aws_iam_group_attached_policies":                       true,
		"aws_iam_role_attached_policies":                        true,
		"aws_iam_group_last_accessed_details":                   true,
		"aws_iam_role_last_accessed_details":                    true,
		"aws_iam_policy_last_accessed_details":                  true,
		"aws_ssoadmin_permission_set_customer_managed_policies": true,
		"aws_ssoadmin_permission_set_managed_policies":          true,
		"aws_organizations_account_parents":                     true,
		"aws_organizations_organizational_unit_parents":         true,
		"aws_stepfunctions_executions":                          true,
		"aws_stepfunctions_map_run_executions":                  true,
		"aws_iam_user_groups":                                   true,
		"aws_iam_groups":                                        true,
	}
	_, ok := tablesToIgnore[tableName]
	return ok
}

func ignorePluralName(tableName string) bool {
	tableNamesToIgnore := map[string]bool{
		"aws_costexplorer_cost_30d":          true,
		"aws_costexplorer_cost_forecast_30d": true,
	}
	_, ok := tableNamesToIgnore[tableName]
	return ok
}
