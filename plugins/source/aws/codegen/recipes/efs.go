package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var efsResources = []*Resource{
	{
		SubService:  "filesystems",
		Struct:      &types.FileSystemDescription{},
		Description: "https://docs.aws.amazon.com/efs/latest/ug/API_FileSystemDescription.html",
		SkipFields:  []string{"FileSystemArn"},
		ExtraColumns: append(
			defaultRegionalColumns,
			[]codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("FileSystemArn")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "backup_policy_status",
					Type:     schema.TypeString,
					Resolver: `ResolveEfsFilesystemBackupPolicyStatus`,
				},
			}...),
	},
}

func EFSResources() []*Resource {
	for _, r := range efsResources {
		r.Service = "efs"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("elasticfilesystem")`
	}
	return efsResources
}
