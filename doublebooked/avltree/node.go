package avltree

import (
	"algo/doublebooked/schedule"
	"algo/doublebooked/utils"
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

func (n *Node) updateMaxEnd() {
	if n.Left != nil && n.Right != nil {
		n.MaxEnd = utils.Max(n.Right.MaxEnd, n.Left.MaxEnd, n.Schedule.End)
	} else if n.Left == nil && n.Right != nil {
		n.MaxEnd = utils.Max(n.Right.MaxEnd, n.Schedule.End)
	} else if n.Left != nil && n.Right == nil {
		n.MaxEnd = utils.Max(n.Left.MaxEnd, n.Schedule.End)
	} else {
		n.MaxEnd = n.Schedule.End
	}
}

func (n *Node) Insert(s schedule.Schedule) (updateBal bool, dup bool) {

	nSchedule := n.Schedule
	nBal := n.bal

	switch {
	case reflect.DeepEqual(nSchedule, s):
		return false, true
	case nSchedule.Start > s.Start || (nSchedule.Start == s.Start && nSchedule.End > s.End):

		if n.Left == nil {

			n.Left = &Node{Schedule: s, MaxEnd: s.End}

			n.updateMaxEnd()

			if n.Right == nil {
				nBal = -1
			} else {
				nBal = 0
			}

		} else {

			n.MaxEnd = utils.Max(n.MaxEnd, s.End)

			if updateBal, dup = n.Left.Insert(s); updateBal && !dup {
				if n.Left.bal < -1 || n.Left.bal > 1 {
					n.rebalance(n.Left)
				} else {
					nBal--
				}
			}
		}
	case nSchedule.Start < s.Start || (nSchedule.Start == s.Start && nSchedule.End < s.End):

		if n.Right == nil {
			n.Right = &Node{Schedule: s, MaxEnd: s.End}

			n.updateMaxEnd()
			if n.Left == nil {
				nBal = 1
			} else {
				nBal = 0
			}

		} else {
			n.MaxEnd = utils.Max(n.MaxEnd, s.End)

			if updateBal, dup = n.Right.Insert(s); updateBal && !dup {
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
		return true, dup
	}
	n.bal = nBal
	return false, dup

}

func (n *Node) rotateLeft(c *Node) {

	//fmt.Println("rotateLeft ", c.Schedule, " node n ", n.Schedule)
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

	//fmt.Println("rotateRight ", c.Schedule, " node n ", n.Schedule)
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
	//fmt.Println("rebalance c ", c.Schedule, "; node n ", n.Schedule)

	//fmt.Println(c.Dump(0, ""))
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

func (n *Node) overlap(s schedule.Schedule) ([2]schedule.Schedule, bool) {

	if reflect.DeepEqual(n.Schedule, s) {
		return [2]schedule.Schedule{}, false
	}

	out := schedule.SortSchedules([2]schedule.Schedule{n.Schedule, s})
	if out[0].Start < out[1].End && out[1].Start < out[0].End {
		return out, true
	}

	return out, false

}

func Overlap(n *Node, s schedule.Schedule) (out map[[2]schedule.Schedule]interface{}) {

	out = make(map[[2]schedule.Schedule]interface{})

	if pair, ok := n.overlap(s); ok {
		out[pair] = new(interface{})
	}

	if n.Left != nil && n.Left.MaxEnd > s.Start {
		out = utils.MapUnion(out, Overlap(n.Left, s))
	}
	if n.Right != nil && n.Schedule.Start < s.End {
		out = utils.MapUnion(out, Overlap(n.Right, s))
	}

	return

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
