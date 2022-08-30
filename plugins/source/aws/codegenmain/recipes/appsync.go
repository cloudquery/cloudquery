package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var AppsyncResources = []*Resource{
	{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn, NamespaceColumn},
		AWSStruct:      &types.GraphqlApi{},
		AWSService:     "AppSync",
		Template:       "resource_get",
		ItemsStruct:    &appsync.ListGraphqlApisOutput{},
		TrimPrefix:     "api_",
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	},
}
