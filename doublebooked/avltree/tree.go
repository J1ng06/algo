package avltree

import (
	"algo/doublebooked/schedule"
	"fmt"
)

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(schedule schedule.Schedule) (err error, dup bool) {

	if !schedule.IsValid() {
		err = fmt.Errorf("Invalid Schedule Input %s", schedule)
		return err, false
	}
	if t.Root == nil {
		t.Root = &Node{Schedule: schedule, MaxEnd: schedule.End}
		return
	}
	_, dup = t.Root.Insert(schedule)
	if t.Root.bal < -1 || t.Root.bal > 1 {
		t.rebalance()
	}

	return nil, dup
}

func (t *Tree) rebalance() {
	fakeParent := &Node{Left: t.Root, Schedule: schedule.Schedule{}}
	fakeParent.rebalance(t.Root)
	t.Root = fakeParent.Left
}

func (t *Tree) Dump() string {
	return t.Root.Dump(0, "")
}
