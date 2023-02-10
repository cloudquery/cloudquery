package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Principals() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_principals",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_Principal.html`,
		Resolver:    fetchRamPrincipals,
		Transform:   transformers.TransformWithStruct(&types.Principal{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			client.AccountPKColumn(true),
			client.RegionPKColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
