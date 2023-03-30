package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Folders() *schema.Table {
	tableName := "aws_quicksight_folders"
	return &schema.Table{
		Name:                tableName,
		Description:         "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Folder.html",
		Resolver:            fetchQuicksightFolders,
		PreResourceResolver: getFolder,
		Transform:           client.TransformWithStruct(&types.Folder{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns:             []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
	}
}
