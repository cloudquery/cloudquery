package incidents

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Crowdscore() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_incidents_crowdscore",
		Resolver:  fetchCrowdscore,
		Transform: transformers.TransformWithStruct(&models.DomainEnvironmentScore{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
