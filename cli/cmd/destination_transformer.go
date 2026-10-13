package cmd

import (
	"context"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformer"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
)

func newDestinationRecordTransformer(destinationSpec specs.Destination, sourceName string, syncTime time.Time, syncGroupId string, cqColumnsNotNull bool) *transformer.RecordTransformer {
	opts := []transformer.RecordTransformerOption{
		transformer.WithSourceNameColumn(sourceName),
		transformer.WithSyncTimeColumn(syncTime),
	}
	if cqColumnsNotNull {
		opts = append(opts, transformer.WithCQColumnsNotNull())
	}
	if syncGroupId != "" {
		opts = append(opts, transformer.WithSyncGroupIdColumn(syncGroupId))
	}
	if destinationSpec.WriteMode == specs.WriteModeAppend {
		opts = append(opts, transformer.WithRemovePKs(), transformer.WithRemoveUniqueConstraints())
	} else if destinationSpec.PKMode == specs.PKModeCQID {
		opts = append(opts, transformer.WithRemovePKs(), transformer.WithCQIDPrimaryKey())
	}
	return transformer.NewRecordTransformer(opts...)
}

func transformSchemaForDestination(ctx context.Context, recordTransformer *transformer.RecordTransformer, transformerClients []plugin.PluginClient, sc *arrow.Schema) ([]byte, error) {
	transformedSchemaBytes, err := plugin.SchemaToBytes(recordTransformer.TransformSchema(sc))
	if err != nil {
		return nil, err
	}
	for _, transformerClient := range transformerClients {
		resp, err := transformerClient.TransformSchema(ctx, &plugin.TransformSchema_Request{Schema: transformedSchemaBytes})
		if err != nil {
			return nil, err
		}
		transformedSchemaBytes = resp.Schema
	}
	return transformedSchemaBytes, nil
}
