package changes

import (
	"os"
	"testing"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
	"github.com/stretchr/testify/require"
)

func getDiff(t *testing.T, diffDataFile string) []*gitdiff.File {
	t.Helper()
	patch, err := os.Open(diffDataFile)
	if err != nil {
		t.Fatal(err)
	}

	files, _, err := gitdiff.Parse(patch)
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func Test_parseColumnChange(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name           string
		args           args
		wantName       string
		wantDataType   string
		wantColumnType columnType
	}{
		{name: "Should parse name and data type when change is a column", args: args{line: "|name|String|"}, wantName: "name", wantDataType: "String"},
		{name: "Should parse name, pk and data type when a column is a primary key", args: args{line: "|name (PK)|String|"}, wantName: "name", wantDataType: "String", wantColumnType: columnTypePK},
		{name: "Should return empty strings when change is not a column", args: args{line: "# Table: azure_appservice_site_auth_settings"}, wantName: "", wantDataType: ""},
		{name: "Should parse name, incremental key and data type when a column is an incremental key", args: args{line: "|updated_at (Incremental Key)|Timestamp|"}, wantName: "updated_at", wantDataType: "Timestamp", wantColumnType: columnTypeIncremental},
		{name: "Should parse name, pk and incremental key", args: args{line: "|updated_at (PK) (Incremental Key)|Timestamp|"}, wantName: "updated_at", wantDataType: "Timestamp", wantColumnType: columnTypeIncremental | columnTypePK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotCol := parseColumnChange(tt.args.line)
			require.Equal(t, tt.wantName, gotName)
			require.Equal(t, tt.wantDataType, gotCol.dataType)
			require.Equal(t, tt.wantColumnType, gotCol.columnType)
		})
	}
}

func Test_parsePKChange(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name      string
		args      args
		wantNames []string
	}{
		{name: "PK present", args: args{line: "The composite primary key for this table is (**org**, **id**, **hook_id**)."}, wantNames: []string{"org", "id", "hook_id"}},
		{name: "no PK", args: args{}, wantNames: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNames := parsePKChange(tt.args.line)
			require.Equal(t, tt.wantNames, gotNames)
		})
	}
}

func Test_getChanges(t *testing.T) {
	tests := []struct {
		name         string
		diffDataFile string
		wantChanges  []change
		wantErr      bool
	}{
		{
			name:         "Should return breaking changes",
			diffDataFile: "testdata/pr_4768_diff.txt",
			wantChanges: []change{
				{
					Text:     "Table `azure_web_functions` was renamed to `azure_appservice_functions`",
					Breaking: true,
				},
				{
					Text:     "Table `azure_appservice_functions`: column `web_app_id` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_appservice_functions`: column added with name `site_id` and type `String`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_subscription_locations` was added",
					Breaking: false,
				},
				{
					Text:     "Table `azure_resources_links` was renamed to `azure_subscription_tenants`",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column `name` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column `properties_notes` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column `properties_source_id` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column `properties_target_id` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column added with name `tenant_id` and type `String`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_subscriptions`: column `managed_by_tenants` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscriptions`: column `tags` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscriptions`: column `tenant_id` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscriptions`: column added with name `subscription_id` and type `String`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_web_publishing_profiles` was removed",
					Breaking: true,
				},
			},
			wantErr: false,
		},
		{
			name:         "Should handle PK changes",
			diffDataFile: "testdata/pr_5636_diff.txt",
			wantChanges: []change{
				{
					Text:     "Table `gcp_resourcemanager_projects`: primary key constraint added to column `name`",
					Breaking: true,
				},
				{
					Text:     "Table `gcp_resourcemanager_projects`: primary key constraint added to column `project_id`",
					Breaking: true,
				},
				{
					Text:     "Table `gcp_resourcemanager_projects`: primary key constraint removed from column `_cq_id`",
					Breaking: true,
				},
			},
		},
		{
			name:         "Should mark newly added PK as breaking",
			diffDataFile: "testdata/pr_5802_diff.txt",
			wantChanges: []change{
				{
					Text:     "Table `aws_ses_configuration_sets`: column added with name `arn (PK)` and type `String`",
					Breaking: true,
				},
				{
					Text:     "Table `aws_ses_configuration_sets`: primary key constraint removed from column `account_id`",
					Breaking: true,
				},
				{
					Text:     "Table `aws_ses_configuration_sets`: primary key constraint removed from column `name`",
					Breaking: true,
				},
				{
					Text:     "Table `aws_ses_configuration_sets`: primary key constraint removed from column `region`",
					Breaking: true,
				},
			},
		},
		{
			name:         "Should not report PK order change if they were other PK changes",
			diffDataFile: "testdata/pr_9677_diff.txt",
			wantChanges: []change{
				{
					Text:     "Table `aws_iam_users`: column `id` removed from table",
					Breaking: true,
				},
				{
					Text:     "Table `aws_iam_users`: primary key constraint added to column `arn`",
					Breaking: true,
				},
			},
		},
		{
			name:         "Should mark PK order change as breaking",
			diffDataFile: "testdata/pr_6012_diff.txt",
			wantChanges: []change{
				{
					Text:     "Table `github_hook_deliveries`: primary key order changed from `org, id, hook_id` to `org, hook_id, id`",
					Breaking: true,
				},
				{
					Text:     "Table `github_issues`: column added with name `state_reason` and type `String`",
					Breaking: false,
				},
				{
					Text:     "Table `github_repositories`: column added with name `has_discussions` and type `Bool`",
					Breaking: false,
				},
				{
					Text:     "Table `github_team_members`: primary key order changed from `org, id, team_id` to `org, team_id, id`",
					Breaking: true,
				},
				{
					Text:     "Table `github_team_repositories`: primary key order changed from `org, id, team_id` to `org, team_id, id`",
					Breaking: true,
				},
				{
					Text:     "Table `github_team_repositories`: column added with name `has_discussions` and type `Bool`",
					Breaking: false,
				},
			},
		},
		{
			name:         "Should handle incremental column changes",
			diffDataFile: "testdata/pr_6707_diff.txt",
			wantChanges: []change{
				{
					Text:     "Table `shopify_abandoned_checkouts`: column `updated_at` added to cursor for incremental syncs",
					Breaking: true,
				},
				{
					Text:     "Table `shopify_customers`: column `created_at` removed from cursor for incremental syncs",
					Breaking: true,
				},
				{
					Text:     "Table `shopify_customers`: column `updated_at` added to cursor for incremental syncs",
					Breaking: true,
				},
				{
					Text:     "Table `shopify_orders`: column `created_at` added to cursor for incremental syncs",
					Breaking: true,
				},
				{
					Text:     "Table `shopify_orders`: column `updated_at` added to cursor for incremental syncs",
					Breaking: true,
				},
				{
					Text:     "Table `shopify_price_rules`: column `updated_at` added to cursor for incremental syncs",
					Breaking: true,
				},
				{
					Text:     "Table `shopify_products`: column `updated_at` added to cursor for incremental syncs",
					Breaking: true,
				},
			},
		},
		{
			name:         "Should handle no backticks -> backticks",
			diffDataFile: "testdata/pr_11034_diff.txt",
			wantChanges:  []change{},
		},
		{
			name:         "Should handle newlines",
			diffDataFile: "testdata/pr_11259_diff.txt",
			wantChanges:  []change{},
		},
		{
			name:         "Should ignore removal of _cq_source_name and _cq_sync_time for SDK v4 migration",
			diffDataFile: "testdata/pr_11855_diff.txt",
			wantChanges:  []change{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := getDiff(t, tt.diffDataFile)
			gotChanges, err := GetChanges(diff)
			if tt.wantErr {
				require.Error(t, err)
			}
			require.Equal(t, tt.wantChanges, gotChanges)
		})
	}
}
