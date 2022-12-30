package priorityqueue

import (
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Item struct {
	Cols []schema.CQType // All columns of the row
}

// NewItem creates a new Item to be used with PriorityQueue
func NewItem(cols []schema.CQType) *Item {
	return &Item{
		Cols: cols,
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	table   *schema.Table
	orderBy []destination.OrderByColumn
	items   []*Item
}

func New(table *schema.Table, orderBy []destination.OrderByColumn) *PriorityQueue {
	pq := &PriorityQueue{
		table:   table,
		orderBy: orderBy,
		items:   make([]*Item, 0),
	}
	return pq
}

func (pq PriorityQueue) Len() int { return len(pq.items) }

func (pq PriorityQueue) Less(i, j int) bool {
	for _, orderBy := range pq.orderBy {
		col := pq.table.Columns.Index(orderBy.Name)
		if col == -1 {
			return false
		}
		if !pq.items[i].Cols[col].Equal(pq.items[j].Cols[col]) {
			less := pq.items[i].Cols[col].LessThan(pq.items[j].Cols[col])
			if orderBy.Desc {
				return !less
			}
			return less
		}
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old.items)
	item := old.items[n-1]
	old.items[n-1] = nil // avoid memory leak
	pq.items = old.items[0 : n-1]
	return item
}
