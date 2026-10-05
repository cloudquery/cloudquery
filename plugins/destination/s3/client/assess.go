package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

const noSpecReason = "no destination spec"

var _ plugin.Assessor = (*Client)(nil)

func (c *Client) AssessTables(_ context.Context, tables []plugin.TablePair, _ plugin.AssessOptions) ([]plugin.TableFinding, error) {
	findings := make([]plugin.TableFinding, len(tables))
	for i, pair := range tables {
		finding, err := c.assessTable(pair)
		if err != nil {
			return nil, fmt.Errorf("failed to assess table %s: %w", pair.TableName(), err)
		}
		findings[i] = finding
	}
	return findings, nil
}

func (c *Client) assessTable(pair plugin.TablePair) (plugin.TableFinding, error) {
	if c.Client == nil {
		return plugin.TableFinding{
			TableName:                pair.TableName(),
			Category:                 plugin.AssessCategoryUnknown,
			CoverageIncomplete:       true,
			CoverageIncompleteReason: noSpecReason,
		}, nil
	}
	return c.AssessTable(pair)
}
