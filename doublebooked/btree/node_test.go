package btree

import (
	"algo/doublebooked/schedule"
	"fmt"
	"reflect"
	"testing"
)

type testMax struct {
	integers []int
	expected interface{}
}

func TestMax(t *testing.T) {
	var testCases = []testMax{
		{
			integers: []int{0},
			expected: 0,
		},
		{
			integers: []int{3, 2, 3},
			expected: 3,
		},
	}

	for i, pair := range testCases {
		result := max(pair.integers...)
		if result != pair.expected {
			t.Error(
				"[ Testcase: TestMax ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%+v", pair.integers), "\n",
				"Expected:", fmt.Sprintf("\n%v", pair.expected), "\n",
				"Got:     ", fmt.Sprintf("\n%v", result), "\n",
			)
		}
	}

}

type testInsertPair struct {
	root      *Node
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
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{1, 2},
				schedule.Schedule{2, 10},
				schedule.Schedule{0, 10},
			},
			expected: &Node{
				Schedule: schedule.Schedule{1, 2},
				MaxEnd:   10,
				Left:     &Node{Schedule: schedule.Schedule{0, 10}, MaxEnd: 10, bal: 0},
				Right:    &Node{Schedule: schedule.Schedule{2, 10}, MaxEnd: 10, bal: 0},
				bal:      0},
		},
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{1, 2},
				schedule.Schedule{0, 10},
				schedule.Schedule{2, 10},
			},
			expected: &Node{
				Schedule: schedule.Schedule{1, 2},
				MaxEnd:   10,
				Left:     &Node{Schedule: schedule.Schedule{0, 10}, MaxEnd: 10, bal: 0},
				Right:    &Node{Schedule: schedule.Schedule{2, 10}, MaxEnd: 10, bal: 0},
				bal:      0},
		},
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{1, 2},
				schedule.Schedule{0, 10},
				schedule.Schedule{-2, 10},
			},
			expected: &Node{
				Schedule: schedule.Schedule{0, 10},
				MaxEnd:   10,
				Right:    &Node{Schedule: schedule.Schedule{1, 2}, MaxEnd: 2, bal: 0},
				Left:     &Node{Schedule: schedule.Schedule{-2, 10}, MaxEnd: 10, bal: 0},
				bal:      0},
		},
	}

	for i, pair := range testCases {

		if pair.root == nil {
			pair.root = &Node{Schedule: pair.schedules[0], MaxEnd: pair.schedules[0].End}
		}

		for j := 1; j < len(pair.schedules); j++ {
			pair.root.Insert(pair.schedules[j])
		}

		if !reflect.DeepEqual(pair.expected, pair.root) {
			t.Error(
				"[ Testcase: TestInsert ", i, " ]\n",
				"For Schedules:     ", fmt.Sprintf("%s", schedule.Flattern(pair.schedules)), "\n",
				"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Node).Dump(0, "")), "\n",
				"Got:     ", fmt.Sprintf("\n%s", pair.root.Dump(0, "")), "\n",
			)
		}
	}
}

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

func TestRotateRightLeft(t *testing.T) {
	var testCases = []testRotatePair{
		{
			node: &Node{
				Schedule: schedule.Schedule{20, 24},
				MaxEnd:   37,
				Left: &Node{
					Schedule: schedule.Schedule{10, 15},
					MaxEnd:   15,
					bal:      0},
				Right: &Node{
					Schedule: schedule.Schedule{30, 35},
					MaxEnd:   37,
					Left: &Node{
						Schedule: schedule.Schedule{20, 28},
						MaxEnd:   28,
						Left: &Node{
							Schedule: schedule.Schedule{20, 25},
							MaxEnd:   25,
							bal:      0},
					},
					Right: &Node{
						Schedule: schedule.Schedule{36, 37},
						MaxEnd:   37,
						bal:      0},
					bal: -1},
				bal: 2},
			expected: &Node{
				Schedule: schedule.Schedule{20, 28},
				MaxEnd:   37,
				Left: &Node{
					Schedule: schedule.Schedule{20, 24},
					MaxEnd:   25,
					Left: &Node{
						Schedule: schedule.Schedule{10, 15},
						MaxEnd:   15,
						bal:      0},
					Right: &Node{
						Schedule: schedule.Schedule{20, 25},
						MaxEnd:   25,
						bal:      0},
					bal: 0},
				Right: &Node{
					Schedule: schedule.Schedule{30, 35},
					MaxEnd:   37,
					Right: &Node{
						Schedule: schedule.Schedule{36, 37},
						MaxEnd:   37,
						bal:      0},
					bal: 1},
				bal: 0,
			},
		},
	}

	for i, pair := range testCases {
		fakeNode := &Node{}
		fakeNode.Left = pair.node
		fakeNode.rotateRightLeft(pair.node)
		if !reflect.DeepEqual(pair.expected, fakeNode.Left) {
			t.Error(
				"[ Testcase: TestRotateRightLeft ", i, " ]\n",
				"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Node).Dump(0, "")), "\n",
				"Got:     ", fmt.Sprintf("\n%s", fakeNode.Left.Dump(0, "")), "\n",
			)
		}
	}

}

func TestRotateLeftRight(t *testing.T) {
	var testCases = []testRotatePair{
		{
			node: &Node{
				Schedule: schedule.Schedule{20, 24},
				MaxEnd:   39,
				Right: &Node{
					Schedule: schedule.Schedule{24, 37},
					MaxEnd:   37,
					bal:      0},
				Left: &Node{
					Schedule: schedule.Schedule{10, 23},
					MaxEnd:   39,
					Right: &Node{
						Schedule: schedule.Schedule{10, 38},
						MaxEnd:   39,
						Right: &Node{
							Schedule: schedule.Schedule{15, 39},
							MaxEnd:   39,
							bal:      0},
						bal: 1,
					},
					Left: &Node{
						Schedule: schedule.Schedule{8, 37},
						MaxEnd:   37,
						bal:      0},
					bal: 1},
				bal: -2},
			expected: &Node{
				Schedule: schedule.Schedule{10, 38},
				MaxEnd:   39,
				Right: &Node{
					Schedule: schedule.Schedule{20, 24},
					MaxEnd:   39,
					Left: &Node{
						Schedule: schedule.Schedule{15, 39},
						MaxEnd:   39,
						bal:      0},
					Right: &Node{
						Schedule: schedule.Schedule{24, 37},
						MaxEnd:   37,
						bal:      0},
					bal: 0},
				Left: &Node{
					Schedule: schedule.Schedule{10, 23},
					MaxEnd:   37,
					Left: &Node{
						Schedule: schedule.Schedule{8, 37},
						MaxEnd:   37,
						bal:      0},
					bal: -1},
				bal: 0,
			},
		},
	}

	for i, pair := range testCases {
		fakeNode := &Node{}
		fakeNode.Left = pair.node
		fakeNode.rotateLeftRight(pair.node)
		if !reflect.DeepEqual(pair.expected, fakeNode.Left) {
			t.Error(
				"[ Testcase: TestRotateLeftRight ", i, " ]\n",
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
		{
			node: &Node{
				Schedule: schedule.Schedule{20, 24},
				MaxEnd:   37,
				Left: &Node{
					Schedule: schedule.Schedule{10, 15},
					MaxEnd:   15,
					bal:      0},
				Right: &Node{
					Schedule: schedule.Schedule{30, 35},
					MaxEnd:   37,
					Left: &Node{
						Schedule: schedule.Schedule{20, 28},
						MaxEnd:   28,
						Left: &Node{
							Schedule: schedule.Schedule{20, 25},
							MaxEnd:   25,
							bal:      0},
					},
					Right: &Node{
						Schedule: schedule.Schedule{36, 37},
						MaxEnd:   37,
						bal:      0},
					bal: -1},
				bal: 2},
			expected: &Node{
				Schedule: schedule.Schedule{20, 28},
				MaxEnd:   37,
				Left: &Node{
					Schedule: schedule.Schedule{20, 24},
					MaxEnd:   25,
					Left: &Node{
						Schedule: schedule.Schedule{10, 15},
						MaxEnd:   15,
						bal:      0},
					Right: &Node{
						Schedule: schedule.Schedule{20, 25},
						MaxEnd:   25,
						bal:      0},
					bal: 0},
				Right: &Node{
					Schedule: schedule.Schedule{30, 35},
					MaxEnd:   37,
					Right: &Node{
						Schedule: schedule.Schedule{36, 37},
						MaxEnd:   37,
						bal:      0},
					bal: 1},
				bal: 0,
			},
		},
		{
			node: &Node{
				Schedule: schedule.Schedule{20, 24},
				MaxEnd:   39,
				Right: &Node{
					Schedule: schedule.Schedule{24, 37},
					MaxEnd:   37,
					bal:      0},
				Left: &Node{
					Schedule: schedule.Schedule{10, 23},
					MaxEnd:   39,
					Right: &Node{
						Schedule: schedule.Schedule{10, 38},
						MaxEnd:   39,
						Right: &Node{
							Schedule: schedule.Schedule{15, 39},
							MaxEnd:   39,
							bal:      0},
						bal: 1,
					},
					Left: &Node{
						Schedule: schedule.Schedule{8, 37},
						MaxEnd:   37,
						bal:      0},
					bal: 1},
				bal: -2},
			expected: &Node{
				Schedule: schedule.Schedule{10, 38},
				MaxEnd:   39,
				Right: &Node{
					Schedule: schedule.Schedule{20, 24},
					MaxEnd:   39,
					Left: &Node{
						Schedule: schedule.Schedule{15, 39},
						MaxEnd:   39,
						bal:      0},
					Right: &Node{
						Schedule: schedule.Schedule{24, 37},
						MaxEnd:   37,
						bal:      0},
					bal: 0},
				Left: &Node{
					Schedule: schedule.Schedule{10, 23},
					MaxEnd:   37,
					Left: &Node{
						Schedule: schedule.Schedule{8, 37},
						MaxEnd:   37,
						bal:      0},
					bal: -1},
				bal: 0,
			},
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
