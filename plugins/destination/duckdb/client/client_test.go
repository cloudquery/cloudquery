package client

import (
	"testing"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func TestPlugin(t *testing.T) {
	if err := types.RegisterAllExtensions(); err != nil {
		t.Fatal(err)
	}

	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("duckdb", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				ConnectionString: "?threads=1",
			},
		},
		destination.PluginTestSuiteTests{
			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
		// not supported in Parquet Writer
		destination.WithTestSourceSkipIntervals(),
		destination.WithTestSourceSkipDurations(),

		// not supported in duckDB for now
		destination.WithTestSourceSkipTimes(),
		destination.WithTestSourceSkipDates(),
		destination.WithTestSourceSkipLargeTypes(),
	)
}

func TestParseConnectionString(t *testing.T) {
	cases := []struct {
		give string
		want parsedConnectionString
	}{
		{give: "md:", want: parsedConnectionString{label: "md"}},
		{give: "motherduck:", want: parsedConnectionString{label: "motherduck"}},
		{give: "md:mydb", want: parsedConnectionString{label: "md", path: "mydb"}},
		{give: "md:mydb?threads=1", want: parsedConnectionString{label: "md", path: "mydb", query: "threads=1"}},
		{give: "md:mydb?threads=1&other=2", want: parsedConnectionString{label: "md", path: "mydb", query: "threads=1&other=2"}},
		{give: "mydb.db", want: parsedConnectionString{path: "mydb.db"}},
		{give: "/absolute/path/to/mydb.db", want: parsedConnectionString{path: "/absolute/path/to/mydb.db"}},
		{give: "../../relative/path/to/mydb.db", want: parsedConnectionString{path: "../../relative/path/to/mydb.db"}},
	}
	for _, c := range cases {
		got := parseConnectionString(c.give)
		if got != c.want {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}
