package analytics

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"maps"
	"slices"
	"sync"

	"github.com/apache/arrow-go/v18/arrow"
)

const (
	AccountDimension    = "account"
	RepositoryDimension = "repository"

	maxHashesPerDimension  = 500
	maxTrackedPerDimension = 50000

	hashLength = 16
)

type columnSpec struct {
	dimension string
	table     string
	column    string
}

var idColumnsBySource = map[string][]columnSpec{
	"cloudquery/aws":        {{dimension: AccountDimension, column: "account_id"}},
	"cloudquery/gcp":        {{dimension: AccountDimension, column: "project_id"}},
	"cloudquery/azure":      {{dimension: AccountDimension, column: "subscription_id"}},
	"cloudquery/k8s":        {{dimension: AccountDimension, column: "cloud_cluster_id"}},
	"cloudquery/cloudflare": {{dimension: AccountDimension, column: "account_id"}},
	"cloudquery/github": {
		{dimension: AccountDimension, column: "org"},
		{dimension: RepositoryDimension, table: "github_repositories", column: "id"},
	},
}

type IDSummary struct {
	Hashes       []string
	Count        int
	Truncated    bool
	CountIsFloor bool
}

type IDCollector interface {
	Observe(tableName string, record arrow.RecordBatch)
	Summaries() map[string]IDSummary
}

type NoopIDCollector struct{}

func (NoopIDCollector) Observe(string, arrow.RecordBatch) {}

func (NoopIDCollector) Summaries() map[string]IDSummary { return nil }

type idCollector struct {
	specs []columnSpec
	salt  string

	mu      sync.Mutex
	hashes  map[string]map[string]struct{}
	dropped map[string]bool
}

func SupportedIDSource(sourcePath string) bool {
	_, ok := idColumnsBySource[sourcePath]
	return ok
}

func NewIDCollector(sourcePath, salt string) IDCollector {
	specs, ok := idColumnsBySource[sourcePath]
	if !ok || salt == "" {
		return NoopIDCollector{}
	}
	return &idCollector{
		specs:   specs,
		salt:    salt,
		hashes:  make(map[string]map[string]struct{}),
		dropped: make(map[string]bool),
	}
}

func (c *idCollector) Observe(tableName string, record arrow.RecordBatch) {
	if record == nil || record.NumRows() == 0 {
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

func (c *idCollector) observeColumn(dimension string, column arrow.Array) {
	c.mu.Lock()
	defer c.mu.Unlock()

	hashes, ok := c.hashes[dimension]
	if !ok {
		hashes = make(map[string]struct{})
		c.hashes[dimension] = hashes
	}

	for i := range column.Len() {
		if column.IsNull(i) {
			continue
		}
		value := column.ValueStr(i)
		if value == "" {
			continue
		}
		hash := c.hashID(value)
		if _, seen := hashes[hash]; seen {
			continue
		}
		if len(hashes) >= maxTrackedPerDimension {
			c.dropped[dimension] = true
			return
		}
		hashes[hash] = struct{}{}
	}
}

func (c *idCollector) Summaries() map[string]IDSummary {
	c.mu.Lock()
	defer c.mu.Unlock()

	summaries := make(map[string]IDSummary, len(c.hashes))
	for dimension, seen := range c.hashes {
		if len(seen) == 0 {
			continue
		}
		hashes := slices.Sorted(maps.Keys(seen))

		truncated := false
		if len(hashes) > maxHashesPerDimension {
			hashes = hashes[:maxHashesPerDimension]
			truncated = true
		}
		summaries[dimension] = IDSummary{
			Hashes:       hashes,
			Count:        len(seen),
			Truncated:    truncated,
			CountIsFloor: c.dropped[dimension],
		}
	}
	if len(summaries) == 0 {
		return nil
	}
	return summaries
}

func (c *idCollector) hashID(value string) string {
	mac := hmac.New(sha256.New, []byte(c.salt))
	mac.Write([]byte(value))
	return hex.EncodeToString(mac.Sum(nil))[:hashLength]
}
