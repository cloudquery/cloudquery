package queries

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPartitionKeyAndSortingKeyQueryDoesNotInterpolateNames(t *testing.T) {
	require.Equal(t,
		"SELECT partition_key, sorting_key FROM system.tables WHERE database = ? AND name = ?",
		GetPartitionKeyAndSortingKeyQuery(),
	)
}

func TestGetTTLQueryEscapesIdentifiers(t *testing.T) {
	require.Equal(t,
		"SHOW CREATE TABLE `db\"name`.`table\"name`",
		GetTTLQuery(`db"name`, `table"name`),
	)
	require.Equal(t,
		"SHOW CREATE TABLE `db``name`.`table\\name`",
		GetTTLQuery("db`name", `table\name`),
	)
}
