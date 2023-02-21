package auditlog

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func AuditLogEntries() *schema.Table {
	return &schema.Table{
		Name:        "launchdarkly_auditlog_entries",
		Description: `https://apidocs.launchdarkly.com/tag/Audit-log#operation/getAuditLogEntries`,
		Resolver:    fetchAuditLogEntries,
		Transform:   transformers.TransformWithStruct(&ldapi.AuditLogEntryListingRep{}, client.SharedTransformers(transformers.WithPrimaryKeys("Id"), transformers.WithSkipFields("Date", "Links"))...),
		Columns: schema.ColumnList{
			{
				Name:     "date",
				Type:     schema.TypeTimestamp,
				Resolver: client.UnixTimeResolver("Date"),
				CreationOptions: schema.ColumnCreationOptions{
					IncrementalKey: true,
				},
			},
		},
		IsIncremental: true,
	}
}

const key = "auditlog_entries"

func fetchAuditLogEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := int64(0)
	if cl.Backend != nil {
		value, err := cl.Backend.Get(ctx, key, cl.ID())
		if err != nil {
			return fmt.Errorf("failed to retrieve state from backend: %w", err)
		}
		if value != "" {
			valInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
			}
			cursor = valInt
		}
	}

	const limit = 20
	for {
		list, b, err := cl.Services.AuditLogApi.GetAuditLogEntries(ctx).Limit(limit).After(cursor).Execute()
		if err != nil {
			return err
		}
		b.Body.Close()
		res <- list.Items

		changed := false
		for i := range list.Items {
			if d := list.Items[i].Date; d > cursor {
				cursor = d
				changed = true
			}
		}

		if cl.Backend != nil && changed {
			if err := cl.Backend.Set(ctx, key, cl.ID(), strconv.FormatInt(cursor, 10)); err != nil {
				return fmt.Errorf("failed to store state in backend: %w", err)
			}
		}

		if len(list.Items) < limit {
			break
		}
	}

	return nil
}
