package avltree

import (
	"algo/doublebooked/schedule"
	"fmt"
	"reflect"
	"strings"
)

type Node struct {
	Schedule schedule.Schedule
	MaxEnd   int
	Left     *Node
	Right    *Node
	bal      int // height(n.Right) - height(n.Left)
}

func max(integers ...int) int {
	max := integers[0]

	for i := 1; len(integers) > 1 && i < len(integers); i++ {
		if integers[i] > max {
			max = integers[i]
		}
	}
	return max
}

func (n *Node) updateMaxEnd() {
	if n.Left != nil && n.Right != nil {
		n.MaxEnd = max(n.Right.MaxEnd, n.Left.MaxEnd, n.Schedule.End)
	} else if n.Left == nil && n.Right != nil {
		n.MaxEnd = max(n.Right.MaxEnd, n.Schedule.End)
	} else if n.Left != nil && n.Right == nil {
		n.MaxEnd = max(n.Left.MaxEnd, n.Schedule.End)
	} else {
		n.MaxEnd = n.Schedule.End
	}
}

func (n *Node) Insert(Schedule schedule.Schedule) bool {

	nSchedule := n.Schedule
	nBal := n.bal

	switch {
	case reflect.DeepEqual(nSchedule, Schedule):
		return false
	case nSchedule.Start > Schedule.Start || (nSchedule.Start == Schedule.Start && nSchedule.End > Schedule.End):

		if n.Left == nil {

			n.Left = &Node{Schedule: Schedule, MaxEnd: Schedule.End}

			n.updateMaxEnd()

			if n.Right == nil {
				nBal = -1
			} else {
				nBal = 0
			}

		} else {

			n.MaxEnd = max(n.MaxEnd, Schedule.End)

			if n.Left.Insert(Schedule) {
				if n.Left.bal < -1 || n.Left.bal > 1 {
					n.rebalance(n.Left)
				} else {
					nBal--
				}
			}
		}
	case nSchedule.Start < Schedule.Start || (nSchedule.Start == Schedule.Start && nSchedule.End < Schedule.End):

		if n.Right == nil {
			n.Right = &Node{Schedule: Schedule, MaxEnd: Schedule.End}

			n.updateMaxEnd()
			if n.Left == nil {
				nBal = 1
			} else {
				nBal = 0
			}

		} else {
			n.MaxEnd = max(n.MaxEnd, Schedule.End)

			if n.Right.Insert(Schedule) {
				if n.Right.bal < -1 || n.Right.bal > 1 {
					n.rebalance(n.Right)
				} else {
					nBal++
				}

			}
		}
	}

	if nBal != n.bal && nBal != 0 {
		n.bal = nBal
		return true
	}
	n.bal = nBal
	return false

}

func (n *Node) rotateLeft(c *Node) {

	fmt.Println("rotateLeft ", c.Schedule, " node n ", n.Schedule)
	r := c.Right
	c.Right = r.Left

	c.updateMaxEnd()

	r.Left = c
	r.updateMaxEnd()

	if c == n.Left {
		n.Left = r
	} else {
		n.Right = r
	}
	c.bal = 0
	r.bal = 0
}

func (n *Node) rotateRight(c *Node) {
	fmt.Println("rotateRight ", c.Schedule, " node n ", n.Schedule)
	l := c.Left
	c.Left = l.Right

	c.updateMaxEnd()

	l.Right = c
	l.updateMaxEnd()

	if c == n.Left {
		n.Left = l
	} else {
		n.Right = l
	}
	c.bal = 0
	l.bal = 0
}

func (n *Node) rotateRightLeft(c *Node) {
	c.Right.Left.bal = 1
	c.rotateRight(c.Right)
	c.Right.bal = 1
	n.rotateLeft(c)
	n.Left.Right.bal = 1
}

func (n *Node) rotateLeftRight(c *Node) {
	// The considerations from rotateRightLeft also apply here.
	c.Left.Right.bal = -1
	c.rotateLeft(c.Left)
	c.Left.bal = -1
	n.rotateRight(c)
	n.Left.Left.bal = -1
}

func (n *Node) rebalance(c *Node) {
	fmt.Println("rebalance c ", c.Schedule, "; node n ", n.Schedule)

	fmt.Println(c.Dump(0, ""))
	switch {
	case c.bal == -2 && c.Left.bal == -1:
		n.rotateRight(c)
	case c.bal == 2 && c.Right.bal == 1:
		n.rotateLeft(c)
	case c.bal == -2 && c.Left.bal == 1:
		n.rotateLeftRight(c)
	case c.bal == 2 && c.Right.bal == -1:
		n.rotateRightLeft(c)
	}
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

func (n *Node) Overlap(schedule schedule.Schedule) bool {

	if reflect.DeepEqual(n.Schedule, schedule) {
		return false
	}

	return n.Schedule.Start < schedule.End && schedule.Start < n.Schedule.End

}

func (n *Node) Dump(i int, lr string) (out string) {
	if n == nil {
		return
	}

	indent := ""
	if i > 0 {
		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}

	out += fmt.Sprintf("%s%s[MaxEnd=%d][bal=%d]\n", indent, n.Schedule.String(), n.MaxEnd, n.bal)
	out += n.Left.Dump(i+1, "L")
	out += n.Right.Dump(i+1, "R")

	return
}
