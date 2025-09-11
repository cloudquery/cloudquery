package otel

import (
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/utils"
)

type TableDurations struct {
	data *utils.ConcurrentMap[string, time.Duration]
}

func NewTableDurations() *TableDurations {
	return &TableDurations{
		data: utils.NewConcurrentMap[string, time.Duration](),
	}
}

func (td *TableDurations) Set(table string, duration time.Duration) {
	d, ok := td.data.Get(table)
	if !ok || duration > d {
		// if not ok, add the duration
		// if duration is greater than existing, update it
		// else do nothing (keep the existing duration)

		td.data.Add(table, duration)
	}
}

func (td *TableDurations) GetAll() map[string]time.Duration {
	return td.data.GetAll()
}
