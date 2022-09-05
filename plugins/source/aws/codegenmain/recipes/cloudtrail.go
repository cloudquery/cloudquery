package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(parentize(&Resource{
		DefaultColumns:  []codegen.ColumnDefinition{AccountIdColumn},
		AWSStruct:       &types.Trail{},
		AWSService:      "CloudTrail",
		Template:        "resource_get",
		ItemsStruct:     &cloudtrail.DescribeTrailsOutput{},
		ColumnOverrides: map[string]codegen.ColumnDefinition{},
		// TODO query and add tags (one call per region)
	},
		&Resource{
			// TODO missing columns
			AWSStruct:       &types.EventSelector{},
			AWSService:      "Cloudtrail",
			Template:        "resource_get",
			ItemsStruct:     &cloudtrail.GetEventSelectorsOutput{},
			ParentFieldName: "TrailARN",
			ChildFieldName:  "TrailName",
		},
	)...)
}
