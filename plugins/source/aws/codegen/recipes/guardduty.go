package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GuarddutyResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "detectors",
			Struct:     &guardduty.DetectorWrapper{},
			SkipFields: []string{"Id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `resolveGuarddutyARN()`,
				},
				{
					Name:    "id",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
			Relations: []string{
				"DetectorMembers()",
			},
		},
		{
			SubService: "detector_members",
			Struct:     &types.Member{},
			SkipFields: []string{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
				},
				{
					Name:     "detector_arn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentResourceFieldResolver("arn")`,
				},
			},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "guardduty"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("guardduty")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
