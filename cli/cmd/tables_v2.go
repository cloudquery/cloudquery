package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/source/v2"
)

func tablesV2(ctx context.Context, sourceClient *managedplugin.Client, path string, format string) error {
	pbSourceClient := source.NewSourceClient(sourceClient.Conn)
	if _, err := pbSourceClient.GenDocs(ctx, &source.GenDocs_Request{
		Format: source.GenDocs_FORMAT(source.GenDocs_FORMAT_value[format]),
		Path:   path,
	}); err != nil {
		return fmt.Errorf("failed to generate docs for %s. Error: %w", sourceClient.Name(), err)
	}
	return nil
}
