package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:      &types.Project{},
		AWSService:     "Codebuild",
		AWSSubService:  "Projects",
		Template:       "resource_manual",
		RawResolver:    "resolvers.FetchProjects",
		//Template:       "resource_list_describe",
		//PaginatorStruct:    &codebuild.ListProjectsOutput{},
		//PaginatorGetStruct: &codebuild.BatchGetProjectsInput{},
		//ItemsStruct:        &codebuild.BatchGetProjectsOutput{},
		ColumnOverrides: map[string]codegen.ColumnDefinition{},
	})
}
