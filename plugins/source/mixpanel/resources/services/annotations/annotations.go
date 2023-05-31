package annotations

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Annotations() *schema.Table {
	return &schema.Table{
		Name:        "mixpanel_annotations",
		Description: `https://developer.mixpanel.com/reference/list-all-annotations-for-project`,
		Resolver:    fetchAnnotations,
		Transform:   client.TransformWithStruct(&mixpanel.Annotation{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.PrimitiveTypes.Int64,
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
