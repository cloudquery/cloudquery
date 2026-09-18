package analytics

import (
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"strings"
	"sync"

	"github.com/apache/arrow-go/v18/arrow"
)

const (
	AccountDimension    = "account"
	RepositoryDimension = "repository"

	maxHashesPerDimension  = 500
	maxTrackedPerDimension = 5000

	hashPepper = "cq-cli-account-analytics-v1"
	hashLength = 16
)

type columnSpec struct {
	dimension string
	table     string
	column    string
}

var idColumnsBySource = map[string][]columnSpec{
	"aws":        {{dimension: AccountDimension, column: "account_id"}},
	"gcp":        {{dimension: AccountDimension, column: "project_id"}},
	"azure":      {{dimension: AccountDimension, column: "subscription_id"}},
	"k8s":        {{dimension: AccountDimension, column: "cloud_cluster_id"}},
	"cloudflare": {{dimension: AccountDimension, column: "account_id"}},
	"github": {
		{dimension: AccountDimension, column: "org"},
		{dimension: RepositoryDimension, table: "github_repositories", column: "id"},
	},
}

type IDSummary struct {
	Hashes    []string
	Count     int
	Truncated bool
}

type IDCollector struct {
	specs []columnSpec

	mu      sync.Mutex
	values  map[string]map[string]struct{}
	dropped map[string]bool
}

func NewIDCollector(sourcePath string) *IDCollector {
	specs, ok := idColumnsBySource[sourceName(sourcePath)]
	if !ok {
		return nil
	}
	return &IDCollector{
		specs:   specs,
		values:  make(map[string]map[string]struct{}),
		dropped: make(map[string]bool),
	}
}

func sourceName(sourcePath string) string {
	_, name, found := strings.Cut(sourcePath, "/")
	if !found {
		return sourcePath
	}
	return name
}

func (c *IDCollector) Observe(tableName string, record arrow.RecordBatch) {
	if c == nil || record == nil || record.NumRows() == 0 {
		return
	}

	for _, spec := range c.specs {
		if spec.table != "" && spec.table != tableName {
			continue
		}
		indices := record.Schema().FieldIndices(spec.column)
		if len(indices) == 0 {
			continue
		}
		c.observeColumn(spec.dimension, record.Column(indices[0]))
	}
}

func (c *IDCollector) observeColumn(dimension string, column arrow.Array) {
	c.mu.Lock()
	defer c.mu.Unlock()

	values, ok := c.values[dimension]
	if !ok {
		values = make(map[string]struct{})
		c.values[dimension] = values
	}

	for i := range column.Len() {
		if column.IsNull(i) {
			continue
		}
		value := column.ValueStr(i)
		if value == "" {
			continue
		}
		if _, seen := values[value]; seen {
			continue
		}
		if len(values) >= maxTrackedPerDimension {
			c.dropped[dimension] = true
			return
		}
		values[value] = struct{}{}
	}
}

func (c *IDCollector) Summaries() map[string]IDSummary {
	if c == nil {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	summaries := make(map[string]IDSummary, len(c.values))
	for dimension, values := range c.values {
		if len(values) == 0 {
			continue
		}
		hashes := make([]string, 0, len(values))
		for value := range values {
			hashes = append(hashes, hashID(value))
		}
		slices.Sort(hashes)

		truncated := c.dropped[dimension]
		if len(hashes) > maxHashesPerDimension {
			hashes = hashes[:maxHashesPerDimension]
			truncated = true
		}
		summaries[dimension] = IDSummary{
			Hashes:    hashes,
			Count:     len(values),
			Truncated: truncated,
		}
	}
	if len(summaries) == 0 {
		return nil
	}
	return summaries
}

func hashID(value string) string {
	sum := sha256.Sum256([]byte(hashPepper + value))
	return hex.EncodeToString(sum[:])[:hashLength]
}
