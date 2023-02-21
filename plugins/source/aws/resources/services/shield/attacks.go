package shield

import (
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Attacks() *schema.Table {
	return &schema.Table{
		Name:                "aws_shield_attacks",
		Description:         `https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_AttackDetail.html`,
		Resolver:            fetchShieldAttacks,
		PreResourceResolver: getAttack,
		Transform:           transformers.TransformWithStruct(&types.AttackDetail{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("shield"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AttackId"),
				Description: `The unique identifier (ID) of the attack`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
