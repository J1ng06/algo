package avltree

import (
	"algo/doublebooked/schedule"
	"fmt"
	"reflect"
	"testing"
)

type testNewTreePair struct {
	schedules []schedule.Schedule
	expected  interface{}
}

// func TestNewTree(t *testing.T) {
// 	testCases := []testNewTreePair{
// 		{
// 			schedules: []schedule.Schedule{
// 				schedule.Schedule{0, 1},
// 				schedule.Schedule{0, 2},
// 			},
// 			expected: &Tree{Root: &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 2, Right: &Node{Schedule: schedule.Schedule{0, 2}, MaxEnd: 2, bal: 0}, bal: 1}},
// 		},
// 		{
// 			schedules: []schedule.Schedule{
// 				schedule.Schedule{1, 1},
// 				schedule.Schedule{0, 2},
// 			},
// 			expected: fmt.Errorf("Invalid Schedule Input [1, 1]"),
// 		},
// 	}
// 	for i, pair := range testCases {

// 		result, err := NewTree(pair.schedules)
// 		if err != nil {
// 			if !reflect.DeepEqual(pair.expected, err) {
// 				t.Error(
// 					"[ Testcase: TestRotateLeft ", i, " ]\n",
// 					"For Schedules:     ", fmt.Sprintf("%+v", pair.schedules), "\n",
// 					"Expected:", fmt.Sprintf("\n%s", pair.expected.(error).Error()), "\n",
// 					"Got:     ", fmt.Sprintf("\n%s", err.Error()), "\n",
// 				)
// 			}
// 		} else {
// 			if !reflect.DeepEqual(pair.expected, result) {
// 				t.Error(
// 					"[ Testcase: TestNewTree ", i, " ]\n",
// 					"For Schedules:     ", fmt.Sprintf("%+v", pair.schedules), "\n",
// 					"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Tree).Dump()), "\n",
// 					"Got:     ", fmt.Sprintf("\n%s", result.Dump()), "\n",
// 				)
// 			}
// 		}
// 	}

// }

func TestTreeInsert(t *testing.T) {
	testCases := []testNewTreePair{
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{0, 1},
				schedule.Schedule{0, 2},
				schedule.Schedule{0, 3},
			},
			expected: &Tree{Root: &Node{Schedule: schedule.Schedule{0, 2}, MaxEnd: 3,
				Left:  &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 1, bal: 0},
				Right: &Node{Schedule: schedule.Schedule{0, 3}, MaxEnd: 3, bal: 0}, bal: 0}},
		},
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{20, 24},
				schedule.Schedule{30, 35},
				schedule.Schedule{10, 15},
				schedule.Schedule{36, 1000},
				schedule.Schedule{20, 28},
				schedule.Schedule{20, 25},
				schedule.Schedule{8, 9},
				schedule.Schedule{13, 14},
				schedule.Schedule{14, 15},
				schedule.Schedule{5, 999},
				schedule.Schedule{1, 2},
			},
			expected: &Tree{Root: &Node{
				Schedule: schedule.Schedule{13, 14}, MaxEnd: 1000, bal: 0,
				Left: &Node{Schedule: schedule.Schedule{8, 9}, MaxEnd: 999, bal: -1,
					Left: &Node{Schedule: schedule.Schedule{5, 999}, MaxEnd: 999, bal: -1,
						Left: &Node{Schedule: schedule.Schedule{1, 2}, MaxEnd: 2, bal: 0},
					},
					Right: &Node{Schedule: schedule.Schedule{10, 15}, MaxEnd: 15, bal: 0},
				},
				Right: &Node{Schedule: schedule.Schedule{20, 28}, MaxEnd: 1000, bal: 0,
					Left: &Node{Schedule: schedule.Schedule{20, 24}, MaxEnd: 25, bal: 0,
						Left:  &Node{Schedule: schedule.Schedule{14, 15}, MaxEnd: 15, bal: 0},
						Right: &Node{Schedule: schedule.Schedule{20, 25}, MaxEnd: 25, bal: 0},
					},
					Right: &Node{Schedule: schedule.Schedule{30, 35}, MaxEnd: 1000, bal: 1,
						Right: &Node{Schedule: schedule.Schedule{36, 1000}, MaxEnd: 1000, bal: 0},
					},
				},
			},
			},
		},
	}
	for i, pair := range testCases {

		tree := new(Tree)

		for i := 0; i < len(pair.schedules); i++ {
			tree.Insert(pair.schedules[i])
		}

		if !reflect.DeepEqual(pair.expected, tree) {
			t.Error(
				"[ Testcase: TestTreeInsert ", i, " ]\n",
				"For Schedules:     ", fmt.Sprintf("%+v", pair.schedules), "\n",
				"Expected:", fmt.Sprintf("\n%s", pair.expected.(*Tree).Dump()), "\n",
				"Got:     ", fmt.Sprintf("\n%s", tree.Dump()), "\n",
			)
		}

	}

}
