package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

const (
	behaviorNoChange     = "makes no changes"
	behaviorKeepVertices = "keeps the existing vertices"
)

var _ plugin.Assessor = (*Client)(nil)

func (*Client) AssessTables(_ context.Context, tables []plugin.TablePair, _ plugin.AssessOptions) ([]plugin.TableFinding, error) {
	findings := make([]plugin.TableFinding, len(tables))
	for i, pair := range tables {
		findings[i] = assessTable(pair)
	}
	return findings, nil
}

func assessTable(pair plugin.TablePair) plugin.TableFinding {
	if pair.New == nil {
		return tableFinding(pair.Old.Name, plugin.AssessCategoryTableRemoved, behaviorKeepVertices)
	}
	finding := tableFinding(pair.New.Name, plugin.AssessCategoryNoChange, behaviorNoChange)
	if pair.Old == nil {
		return finding
	}
	for _, change := range pair.New.GetChanges(pair.Old) {
		if change.ColumnName == "" {
			continue
		}
		finding.Columns = append(finding.Columns, plugin.ColumnFinding{
			ColumnName:         change.ColumnName,
			Category:           plugin.AssessCategoryNoChange,
			SafeModeBehavior:   behaviorNoChange,
			ForcedModeBehavior: behaviorNoChange,
		})
	}
	return finding
}

func tableFinding(name string, category plugin.AssessCategory, behavior string) plugin.TableFinding {
	return plugin.TableFinding{
		TableName:          name,
		Category:           category,
		SafeModeBehavior:   behavior,
		ForcedModeBehavior: behavior,
	}
}
