package annotations

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Annotations() *schema.Table {
	return &schema.Table{
		Name:      "mixpanel_annotations",
		Resolver:  fetchAnnotations,
		Transform: transformers.TransformWithStruct(&mixpanel.Cohort{}, transformers.WithPrimaryKeys("id")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProjectID,
			},
		},
	}
}

func fetchAnnotations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	ret, err := cl.Services.ListAnnotations(ctx)
	if err != nil {
		return err
	}
	if ret.Status != mixpanel.StatusOK {
		return fmt.Errorf("api call failed: %s", ret.Error)
	}

	res <- ret.Results
	return nil
}
