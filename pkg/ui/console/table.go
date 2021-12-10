package console

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
)

type Table struct {
	writer *tablewriter.Table
}

func (t Table) SetHeaders(headers ...string) {
	t.writer.SetHeader(headers)
}

func (t Table) Append(i ...interface{}) {
	t.writer.Append(cast.ToStringSlice(i))
}

func (t Table) Render() {
	t.writer.Render()
}
