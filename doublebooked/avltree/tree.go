package avltree

import (
	"algo/doublebooked/schedule"
	"fmt"
)

type Tree struct {
	Root *Node
}

func NewTree(schedules []schedule.Schedule) (tree *Tree, err error) {

	tree = new(Tree)
	for i := 0; i < len(schedules); i++ {

		if !schedules[i].IsValid() {
			err = fmt.Errorf("Invalid Schedule Input %s", schedules[i].String())
			return nil, err
		}
		fmt.Println("Insert ----", schedules[i])
		tree.Insert(schedules[i])
		fmt.Println(tree.Dump())

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
	fmt.Println("rebalance from tree.go")
	fakeParent := &Node{Left: t.Root, Schedule: schedule.Schedule{}}
	fakeParent.rebalance(t.Root)
	t.Root = fakeParent.Left
}

func (t *Tree) Dump() string {
	return t.Root.Dump(0, "")
}
