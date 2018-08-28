package btree

import (
	"algo/doublebooked/schedule"
)

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(schedule schedule.Schedule) {
	if t.Root == nil {
		t.Root = &Node{Schedule: schedule, MaxEnd: schedule.End}
		return
	}
	t.Root.Insert(schedule)
	if t.Root.bal < -1 || t.Root.bal > 1 {
		t.rebalance()
	}
}

func (t *Tree) rebalance() {
	fakeParent := &Node{Left: t.Root, Schedule: schedule.Schedule{}}
	fakeParent.rebalance(t.Root)
	t.Root = fakeParent.Left
}

func (t *Tree) Dump() {
	t.Root.Dump(0, "")
}
