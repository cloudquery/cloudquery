package transformations

import (
	"fmt"
	"strings"
)

// queryBuilder is a helper struct to build a query from multiple subqueries
//
// Example:
//
// You want these queries:
//
// - `SELECT month, num, link FROM source_table`
// - `SELECT *, ('cloudquery_sync') AS source FROM source_table`
// - `SELECT * FROM view_3 WHERE year = 2006`
//
// You can add each query separately using `.add()` and then build the final query using `.build()`:
//
// CREATE VIEW month, num, link FROM source_table
// CREATE VIEW view_2 AS SELECT *, ('cloudquery_sync') AS source FROM view_1
// CREATE VIEW view_3 AS SELECT * FROM view_2 WHERE year = 2006
// SELECT * FROM view_3
//
// This is necessary because subsequent queries SELECT on the result of the previous queries.
type queryBuilder struct {
	subqueries []string
}

func newQueryBuilder() *queryBuilder {
	return &queryBuilder{}
}

func (qb *queryBuilder) add(subquery string) {
	source_table := "source_table"
	if len(qb.subqueries) > 0 {
		source_table = fmt.Sprintf("view_%v", len(qb.subqueries))
	}
	subquery = strings.Replace(subquery, "source_table", source_table, 1)
	qb.subqueries = append(qb.subqueries, fmt.Sprintf("CREATE VIEW view_%v AS %v", len(qb.subqueries)+1, subquery))
}

func (qb *queryBuilder) build() []string {
	qb.subqueries = append(qb.subqueries, fmt.Sprintf("SELECT * FROM view_%v", len(qb.subqueries)))
	return qb.subqueries
}
