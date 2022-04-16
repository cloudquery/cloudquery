package history

import (
	"time"

	"github.com/cloudquery/cq-provider-sdk/database/dsn"
)

const SchemaName = "history"

type Config struct {
	// Retention of data in days, defaults to 7
	Retention int `default:"7" hcl:"retention,optional"`
	// TimeInterval defines how chunks are split by time defaults to one chunk per 24 hours.
	TimeInterval int `default:"24" hcl:"interval,optional"`
	// TimeTruncation truncates fetch time by hour, for example if we fetch with TimeTruncation = 1 at 11:25 the fetch date will truncate to 11:00
	// defaults to 24 hours, which means one set of fetch data per day.
	TimeTruncation int `default:"24" hcl:"truncation,optional"`
}

func (c Config) FetchDate() time.Time {
	return time.Now().UTC().Truncate(time.Duration(c.TimeTruncation) * time.Hour)
}

// TransformDSN sets the search_path of the given DSN to the history schema
func TransformDSN(inputDSN string) (string, error) {
	return dsn.SetDSNElement(inputDSN, map[string]string{"search_path": SchemaName})
}
