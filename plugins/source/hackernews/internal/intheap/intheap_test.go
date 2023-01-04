package intheap

import (
	"container/heap"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	got := make([]int, 0, 4)
	for h.Len() > 0 {
		got = append(got, heap.Pop(h).(int))
	}
	want := []int{1, 2, 3, 5}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("heap mismatch (-got +want):\n%s", diff)
	}
}
