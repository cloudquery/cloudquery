package alerts

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func Query() *schema.Table {
	return &schema.Table{
		Name:      "crowdstrike_alerts_query",
		Resolver:  fetchQuery,
		Transform: transformers.TransformWithStruct(&models.MsaQueryResponse{}),
		Columns:   []schema.Column{},
	}
}
