package btree

import (
	"algo/doublebooked/schedule"
	"fmt"
)

type Tree struct {
	Root *Node
}

func NewTree(schedules []schedule.Schedule) (tree *Tree, err error) {

	for i := 0; i < len(schedules); i++ {

		tree = new(Tree)

		if !schedules[i].IsValid() {
			err = fmt.Errorf("Invalid Schedule Input %s", schedules[i].String())
			return nil, err
		}

		tree.Insert(schedules[i])

	}

	return

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

func Overlap(n *Node, schedule schedule.Schedule) (out []schedule.Schedule) {

	if n.Overlap(schedule) {
		out = append(out, n.Schedule)
	}

	if n.Left != nil && n.Left.MaxEnd > schedule.Start {
		out = append(out, Overlap(n.Left, schedule)...)
	}
	if n.Right != nil && n.Schedule.Start < schedule.End {
		out = append(out, Overlap(n.Right, schedule)...)
	}

	return

}
