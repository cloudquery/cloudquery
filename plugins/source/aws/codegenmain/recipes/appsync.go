package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn, NamespaceColumn},
		AWSStruct:      &types.GraphqlApi{},
		AWSService:     "AppSync",
		Template:       "resource_get",
		ItemsStruct:    &appsync.ListGraphqlApisOutput{},
		PrimaryKeys:    []string{"arn"},
		TrimPrefix:     "api_",
	},
	)
}
