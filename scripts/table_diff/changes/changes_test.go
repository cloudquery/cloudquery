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
		name         string
		args         args
		wantName     string
		wantDataType string
	}{
		{name: "Should parse name and data type when change is a column", args: args{line: "|name|String|"}, wantName: "name", wantDataType: "String"},
		{name: "Should return empty strings when change is not a column", args: args{line: "# Table: azure_appservice_site_auth_settings"}, wantName: "", wantDataType: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotDataType := parseColumnChange(tt.args.line)
			require.Equal(t, tt.wantName, gotName)
			require.Equal(t, tt.wantDataType, gotDataType)
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
					Text:     "Table `azure_appservice_functions`: column removed `web_app_id` from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_appservice_functions`: column added with name `site_id` and type `String`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `function_app_id`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `href`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `kind`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `language`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `script_href`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `script_root_path_href`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `secrets_file_href`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `test_data_href`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `test_data`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_appservice_functions`: column order changed for `type`",
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
					Text:     "Table `azure_subscription_tenants`: column removed `name` from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column removed `properties_notes` from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column removed `properties_source_id` from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column removed `properties_target_id` from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscription_tenants`: column added with name `tenant_id` and type `String`",
					Breaking: false,
				},
				{
					Text:     "Table `azure_subscriptions`: column removed `managed_by_tenants` from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscriptions`: column removed `tags` from table",
					Breaking: true,
				},
				{
					Text:     "Table `azure_subscriptions`: column removed `tenant_id` from table",
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
