package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConfigResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "configuration_recorders",
			Struct:     &models.ConfigurationRecorderWrapper{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `generateConfigRecorderArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "conformance_packs",
			Struct:      &types.ConformancePackDetail{},
			Description: "https://docs.aws.amazon.com/config/latest/APIReference/API_ConformancePackDetail.html",
			SkipFields:  []string{"ConformancePackArn", "ConformancePackInputParameters"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ConformancePackArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"ConformancePackRuleCompliances()",
			},
		},
		{
			SubService: "conformance_pack_rule_compliances",
			Struct:     &models.ConformancePackComplianceWrapper{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "conformance_pack_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "config"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("config")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
