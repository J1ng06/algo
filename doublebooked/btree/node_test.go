package btree

import (
	"algo/doublebooked/schedule"
	"fmt"
	"reflect"
	"testing"
)

type testRotatePair struct {
	node     *Node
	expected interface{}
}

func TestRotateLeft(t *testing.T) {
	var testCases = []testRotatePair{
		{
			node: &Node{
				Schedule: schedule.Schedule{0, 1},
				MaxEnd:   3,
				Right: &Node{
					Schedule: schedule.Schedule{0, 2},
					MaxEnd:   3,
					Right: &Node{
						Schedule: schedule.Schedule{0, 3},
						MaxEnd:   3,
					},
					bal: 1},
				bal: 2},
			expected: &Node{
				Schedule: schedule.Schedule{0, 2},
				MaxEnd:   3,
				Left:     &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 1, bal: 0},
				Right:    &Node{Schedule: schedule.Schedule{0, 3}, MaxEnd: 3, bal: 0},
				bal:      0},
		},
	}

	for i, pair := range testCases {

		fakeNode := &Node{}
		fakeNode.Left = pair.node
		fakeNode.rotateLeft(pair.node)
		if !reflect.DeepEqual(pair.expected, fakeNode.Left) {
			t.Error(
				"[ Testcase: TestRotateLeft ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%s", pair.node.Dump(0, "")), "\n",
				"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Node).Dump(0, "")), "\n",
				"Got:     ", fmt.Sprintf("\n%s", fakeNode.Left.Dump(0, "")), "\n",
			)
		}
	}

}

func TestRotateRight(t *testing.T) {
	var testCases = []testRotatePair{
		{
			node: &Node{
				Schedule: schedule.Schedule{0, 3},
				MaxEnd:   3,
				Left: &Node{
					Schedule: schedule.Schedule{0, 2},
					MaxEnd:   2,
					Left: &Node{
						Schedule: schedule.Schedule{0, 1},
						MaxEnd:   1,
					},
					bal: -1},
				bal: -2},
			expected: &Node{
				Schedule: schedule.Schedule{0, 2},
				MaxEnd:   3,
				Left:     &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 1, bal: 0},
				Right:    &Node{Schedule: schedule.Schedule{0, 3}, MaxEnd: 3, bal: 0},
				bal:      0},
		},
	}

	for i, pair := range testCases {

		fakeNode := &Node{}
		fakeNode.Left = pair.node
		fakeNode.rotateRight(pair.node)
		if !reflect.DeepEqual(pair.expected, fakeNode.Left) {
			t.Error(
				"[ Testcase: TestRotateRight ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%s", pair.node.Dump(0, "")), "\n",
				"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Node).Dump(0, "")), "\n",
				"Got:     ", fmt.Sprintf("\n%s", fakeNode.Left.Dump(0, "")), "\n",
			)
		}
	}

}

func TestRebalance(t *testing.T) {
	var testCases = []testRotatePair{
		{
			node: &Node{
				Schedule: schedule.Schedule{0, 3},
				MaxEnd:   3,
				Left: &Node{
					Schedule: schedule.Schedule{0, 2},
					MaxEnd:   2,
					Left: &Node{
						Schedule: schedule.Schedule{0, 1},
						MaxEnd:   1,
					},
					bal: -1},
				bal: -2},
			expected: &Node{
				Schedule: schedule.Schedule{0, 2},
				MaxEnd:   3,
				Left:     &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 1, bal: 0},
				Right:    &Node{Schedule: schedule.Schedule{0, 3}, MaxEnd: 3, bal: 0},
				bal:      0},
		},
		{
			node: &Node{
				Schedule: schedule.Schedule{0, 1},
				MaxEnd:   3,
				Right: &Node{
					Schedule: schedule.Schedule{0, 2},
					MaxEnd:   3,
					Right: &Node{
						Schedule: schedule.Schedule{0, 3},
						MaxEnd:   3,
					},
					bal: 1},
				bal: 2},
			expected: &Node{
				Schedule: schedule.Schedule{0, 2},
				MaxEnd:   3,
				Left:     &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 1, bal: 0},
				Right:    &Node{Schedule: schedule.Schedule{0, 3}, MaxEnd: 3, bal: 0},
				bal:      0},
		},
	}

	for i, pair := range testCases {

		fakeNode := &Node{}
		fakeNode.Left = pair.node
		fakeNode.rebalance(pair.node)
		if !reflect.DeepEqual(pair.expected, fakeNode.Left) {
			t.Error(
				"[ Testcase: TestRebalance ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%s", pair.node.Dump(0, "")), "\n",
				"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Node).Dump(0, "")), "\n",
				"Got:     ", fmt.Sprintf("\n%s", fakeNode.Left.Dump(0, "")), "\n",
			)
		}
	}

}

type testInsertPair struct {
	schedules []schedule.Schedule
	expected  interface{}
}

func TestInsert(t *testing.T) {
	var testCases = []testInsertPair{
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{0, 1},
				schedule.Schedule{0, 2},
			},
			expected: &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 2, Right: &Node{Schedule: schedule.Schedule{0, 2}, MaxEnd: 2, bal: 0}, bal: 1},
		},
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{1, 2},
				schedule.Schedule{0, 10},
			},
			expected: &Node{Schedule: schedule.Schedule{1, 2}, MaxEnd: 10, Left: &Node{Schedule: schedule.Schedule{0, 10}, MaxEnd: 10, bal: 0}, bal: -1},
		},
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{1, 2},
				schedule.Schedule{1, 2},
			},
			expected: &Node{Schedule: schedule.Schedule{1, 2}, MaxEnd: 2, bal: 0},
		},
	}

	for i, pair := range testCases {

		result := &Node{Schedule: pair.schedules[0], MaxEnd: pair.schedules[0].End}

		for j := 1; j < len(pair.schedules); j++ {
			result.Insert(pair.schedules[j])
		}

		if !reflect.DeepEqual(pair.expected, result) {
			t.Error(
				"[ Testcase: TestInsert ", i, " ]\n",
				"For Schedules:     ", fmt.Sprintf("%s", schedule.Flattern(pair.schedules)), "\n",
				"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Node).Dump(0, "")), "\n",
				"Got:     ", fmt.Sprintf("\n%s", result.Dump(0, "")), "\n",
			)
		}
	}
}

type testOverlapPair struct {
	node     *Node
	schedule schedule.Schedule
	expected interface{}
}

func TestOverlap(t *testing.T) {
	var testCases = []testOverlapPair{
		{
			node:     &Node{Schedule: schedule.Schedule{0, 2}},
			schedule: schedule.Schedule{0, 1},
			expected: true,
		},
		{
			node:     &Node{Schedule: schedule.Schedule{0, 2}},
			schedule: schedule.Schedule{2, 4},
			expected: false,
		},
		{
			node:     &Node{Schedule: schedule.Schedule{0, 2}},
			schedule: schedule.Schedule{0, 2},
			expected: false,
		},
	}

	for i, pair := range testCases {

		result := pair.node.Overlap(pair.schedule)

		if !reflect.DeepEqual(pair.expected, result) {
			t.Error(
				"[ Testcase: TestInsert ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%s", pair.node.Dump(0, "")), "\n",
				"Expected:", fmt.Sprintf("\n%v", pair.expected), "\n",
				"Got:     ", fmt.Sprintf("\n%v", result), "\n",
			)
		}
	}
}
