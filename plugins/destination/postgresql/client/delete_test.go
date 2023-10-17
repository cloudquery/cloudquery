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
			want: `DELETE from "table1" where ( "id" = $1 ) RETURNING *`,
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
			want: `DELETE from "table1" where ( "id1" = $1 AND "id2" = $2 ) RETURNING *`,
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
			want: `DELETE from "table1" where ( "id1" = $1 OR "id2" = $2 ) RETURNING *`,
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
			want: `DELETE from "table1" where ( "id1" = $1 OR "id2" = $2 ) AND ( "id1" = $3 AND "id2" = $4 ) RETURNING *`,
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

func TestGenerateRelationsDelete(t *testing.T) {
	tests := []struct {
		tableRelation message.TableRelation
		want          string
	}{
		{
			tableRelation: message.TableRelation{
				TableName:   "relation_table1",
				ParentTable: "parent_table",
			},
			want: `DELETE from "relation_table1" where "_cq_parent_id" in (select "_cq_id" from "parent_table_CTE")`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.tableRelation.TableName, func(t *testing.T) {
			got := generateRelationsDelete(tt.tableRelation)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}

func TestGenerateDeleteCTE(t *testing.T) {
	tests := []struct {
		name         string
		deleteRecord message.DeleteRecord
		want         string
	}{
		{
			deleteRecord: message.DeleteRecord{
				TableName: "table1",
				WhereClause: message.PredicateGroups{
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
				TableRelations: []message.TableRelation{
					{
						TableName:   "relation_table1",
						ParentTable: "table1",
					},
				},
			},
			want: `WITH "table1_CTE" AS (DELETE from "table1" where ( "id" = $1 ) RETURNING *) , "relation_table1_CTE" AS (DELETE from "relation_table1" where "_cq_parent_id" in (select "_cq_id" from "table1_CTE") RETURNING "_cq_id") Select count(*) from "relation_table1_CTE" UNION ALL Select count(*) from "table1_CTE"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateDeleteCTE(tt.deleteRecord)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("%s", diff)
			}
		})
	}
}
