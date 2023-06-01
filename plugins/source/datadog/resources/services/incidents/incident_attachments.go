package incidents

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func IncidentAttachments() *schema.Table {
	return &schema.Table{
		Name:      "datadog_incident_attachments",
		Transform: client.TransformWithStruct(&datadogV2.IncidentAttachmentData{}, transformers.WithPrimaryKeys("Id")),
		Resolver:  fetchIncidentAttachments,
		Columns: []schema.Column{
			client.AccountNameColumn,
			{
				Name:       "incident_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIncidentAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(datadogV2.IncidentResponseData)
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.IncidentsAPI.ListIncidentAttachments(ctx, p.Id)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
