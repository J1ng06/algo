package btree

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

func (n *Node) Insert(Schedule schedule.Schedule) bool {

	nSchedule := n.Schedule

	switch {
	case reflect.DeepEqual(nSchedule, Schedule):
		return false
	case nSchedule.Start > Schedule.Start || (nSchedule.Start == Schedule.Start && nSchedule.End > Schedule.End):

		if n.Left == nil {

			n.Left = &Node{Schedule: Schedule, MaxEnd: Schedule.End}

			if n.Left.MaxEnd > n.MaxEnd {
				n.MaxEnd = n.Left.MaxEnd
			}

			if n.Right == nil {
				n.bal = -1
				return true
			}
			n.bal = 0
			return false

		}

		if n.MaxEnd < Schedule.End {
			n.MaxEnd = Schedule.End
		}
		if n.Left.Insert(Schedule) {
			if n.Left.bal < -1 || n.Left.bal > 1 {
				n.rebalance(n.Left)
			} else {
				n.bal--
			}
		} else {
			return false
		}
	case nSchedule.Start < Schedule.Start || (nSchedule.Start == Schedule.Start && nSchedule.End < Schedule.End):

		if n.Right == nil {
			n.Right = &Node{Schedule: Schedule, MaxEnd: Schedule.End}

			if n.Right.MaxEnd > n.MaxEnd {
				n.MaxEnd = n.Right.MaxEnd
			}

			if n.Left == nil {
				n.bal = 1
				return true
			}
			n.bal = 0
			return false

		}
		if n.MaxEnd < Schedule.End {
			n.MaxEnd = Schedule.End
		}
		if n.Right.Insert(Schedule) {
			if n.Right.bal < -1 || n.Right.bal > 1 {
				n.rebalance(n.Right)
			} else {
				n.bal++
			}
		} else {
			return false
		}
	}

	if n.bal != 0 {
		return true
	}

	return false
}

func (n *Node) rotateLeft(c *Node) {

	fmt.Println("rotateLeft ")
	r := c.Right
	c.Right = r.Left
	if c.Right == nil {
		if c.Left != nil && c.MaxEnd <= c.Left.MaxEnd {
			c.MaxEnd = c.Left.MaxEnd
		} else {
			c.MaxEnd = c.Schedule.End
		}
	} else if c.Right.MaxEnd > c.Schedule.End {
		c.MaxEnd = c.Right.MaxEnd
	} else if c.Left.MaxEnd > c.Schedule.End {
		c.MaxEnd = c.Left.MaxEnd
	} else {
		c.MaxEnd = c.Schedule.End
	}

	r.Left = c
	if r.MaxEnd < c.MaxEnd {
		r.MaxEnd = c.MaxEnd
	}

	if c == n.Left {
		n.Left = r
	} else {
		n.Right = r
	}
	c.bal = 0
	r.bal = 0
}

func (n *Node) rotateRight(c *Node) {
	fmt.Println("rotateRight ")
	l := c.Left
	c.Left = l.Right
	if c.Left == nil {
		if c.Right != nil && c.MaxEnd <= c.Right.MaxEnd {
			c.MaxEnd = c.Right.MaxEnd
		} else {
			c.MaxEnd = c.Schedule.End
		}
	} else if c.Left.MaxEnd > c.Schedule.End {
		c.MaxEnd = c.Left.MaxEnd
	} else if c.Right.MaxEnd > c.Schedule.End {
		c.MaxEnd = c.Right.MaxEnd
	} else {
		c.MaxEnd = c.Schedule.End
	}

	l.Right = c
	if l.MaxEnd < c.MaxEnd {
		l.MaxEnd = c.MaxEnd
	}
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
	c.Left.Right.bal = -1 // The considerations from rotateRightLeft also apply here.
	c.rotateLeft(c.Left)
	c.Left.bal = -1
	n.rotateRight(c)
	n.Left.Left.bal = -1
}

func (n *Node) rebalance(c *Node) {
	fmt.Println("rebalance " + c.Schedule.String())

	c.Dump(0, "")
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

func (n *Node) Overlap(schedule schedule.Schedule) bool {

	if reflect.DeepEqual(n.Schedule, schedule) {
		return false
	}

	return n.Schedule.Start < schedule.End && schedule.Start < n.Schedule.End

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

func (n *Node) Dump(i int, lr string) {
	if n == nil {
		return
	}

	indent := ""
	if i > 0 {
		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}
	fmt.Printf("%s%s[MaxEnd=%d][bal=%d]\n", indent, n.Schedule.String(), n.MaxEnd, n.bal)
	n.Left.Dump(i+1, "L")
	n.Right.Dump(i+1, "R")
}
