package priorityqueue

import (
	"container/heap"
	"sort"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TestPriorityQueue(t *testing.T) {
	table := &schema.Table{
		Name: "test",
		Columns: []schema.Column{
			{
				Name: "id",
				Type: schema.TypeInt,
			},
			{
				Name: "version",
				Type: schema.TypeInt,
			},
		},
	}
	orderBy := []destination.OrderByColumn{
		{
			Name: "id",
			Desc: false,
		},
		{
			Name: "version",
			Desc: true,
		},
	}
	pq := New(table, orderBy)
	heap.Push(pq, NewItem([]schema.CQType{
		&schema.Int8{Int: 1},
		&schema.Int8{Int: 1},
	}))
	if pq.Len() != 1 {
		t.Fatalf("after first push, pq.Len() = %d, want 1", pq.Len())
	}
	if pq.items[0].Cols[0].(*schema.Int8).Int != 1 {
		t.Fatalf("after first push, pq.items[0].Cols[0].(*schema.Int8).Int = %d, want 1", pq.items[0].Cols[0].(*schema.Int8).Int)
	}

	heap.Push(pq, NewItem([]schema.CQType{
		&schema.Int8{Int: 2},
		&schema.Int8{Int: 1},
	}))
	if pq.Len() != 2 {
		t.Fatalf("after second push, pq.Len() = %d, want 2", pq.Len())
	}

	heap.Push(pq, NewItem([]schema.CQType{
		&schema.Int8{Int: 2},
		&schema.Int8{Int: 2},
	}))
	if pq.Len() != 3 {
		t.Fatalf("after third push, pq.Len() = %d, want 3", pq.Len())
	}

	items := []*Item{
		heap.Pop(pq).(*Item),
		heap.Pop(pq).(*Item),
		heap.Pop(pq).(*Item),
	}
	// items will be popped in reverse priority order, so we reverse the slice
	reverseSlice(items)

	if items[0].Cols[0].(*schema.Int8).Int != 1 {
		t.Fatalf("after pop, items[0] id = %d, want 1", items[0].Cols[0].(*schema.Int8).Int)
	}
	if items[1].Cols[0].(*schema.Int8).Int != 2 {
		t.Fatalf("after pop, items[1] id = %d, want 2", items[1].Cols[0].(*schema.Int8).Int)
	}
	if items[1].Cols[1].(*schema.Int8).Int != 2 {
		t.Fatalf("after pop, items[1] version = %d, want 2", items[1].Cols[1].(*schema.Int8).Int)
	}
	if items[2].Cols[0].(*schema.Int8).Int != 2 {
		t.Fatalf("after pop, items[2] id = %d, want 2", items[2].Cols[0].(*schema.Int8).Int)
	}
	if items[2].Cols[1].(*schema.Int8).Int != 1 {
		t.Fatalf("after pop, items[2] version = %d, want 1", items[2].Cols[1].(*schema.Int8).Int)
	}
}

func reverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
