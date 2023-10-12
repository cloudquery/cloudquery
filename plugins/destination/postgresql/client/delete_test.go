package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/google/go-cmp/cmp"
)

func TestGenerateInitialDelete(t *testing.T) {
	tests := []struct {
		name        string
		tableName   string
		whereClause message.PredicateGroups
		want        string
	}{
		{
			name:      "single pk",
			tableName: "table1",
			whereClause: message.PredicateGroups{
				{
					GroupingType: "AND",
					Predicates: []message.Predicate{
						{
							Operator: "eq",
							Column:   "id",
						},
					},
				},
			},
			want: `DELETE from "table1" where ( "id" = $1 ) RETURNING "_cq_id"`,
		},
		{
			name:      "multiple pks-AND",
			tableName: "table1",
			whereClause: message.PredicateGroups{
				{
					GroupingType: "AND",
					Predicates: []message.Predicate{
						{
							Operator: "eq",
							Column:   "id1",
						},
						{
							Operator: "eq",
							Column:   "id2",
						},
					},
				},
			},
			want: `DELETE from "table1" where ( "id1" = $1 AND "id2" = $2 ) RETURNING "_cq_id"`,
		},
		{
			name:      "multiple pks-OR",
			tableName: "table1",
			whereClause: message.PredicateGroups{
				{
					GroupingType: "OR",
					Predicates: []message.Predicate{
						{
							Operator: "eq",
							Column:   "id1",
						},
						{
							Operator: "eq",
							Column:   "id2",
						},
					},
				},
			},
			want: `DELETE from "table1" where ( "id1" = $1 OR "id2" = $2 ) RETURNING "_cq_id"`,
		},
		{
			name:      "multiple pks-OR+AND",
			tableName: "table1",
			whereClause: message.PredicateGroups{
				{
					GroupingType: "OR",
					Predicates: []message.Predicate{
						{
							Operator: "eq",
							Column:   "id1",
						},
						{
							Operator: "eq",
							Column:   "id2",
						},
					},
				},
				{
					GroupingType: "AND",
					Predicates: []message.Predicate{
						{
							Operator: "eq",
							Column:   "id1",
						},
						{
							Operator: "eq",
							Column:   "id2",
						},
					},
				},
			},
			want: `DELETE from "table1" where ( "id1" = $1 OR "id2" = $2 ) AND ( "id1" = $3 AND "id2" = $4 ) RETURNING "_cq_id"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateInitialDelete(tt.tableName, tt.whereClause)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("%s", diff)
			}

		})
	}
}
