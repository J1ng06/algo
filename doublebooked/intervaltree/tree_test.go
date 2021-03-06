package intervaltree

import (
	"algo/doublebooked/schedule"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type testNewTreePair struct {
	schedules []schedule.Schedule
	expected  interface{}
}

func TestTreeInsert(t *testing.T) {
	testCases := []testNewTreePair{
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{24, 21},
				schedule.Schedule{20, 24},
			},
			expected: errors.New("Invalid Schedule Input"),
		},
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

		var err error
		for i := 0; i < len(pair.schedules); i++ {
			if _, err = tree.Insert(pair.schedules[i]); err != nil {
				break
			}
		}

		if err != nil {
			if err.Error() != "Invalid Schedule Input [24, 21]" {
				t.Error(
					"[ Testcase: TestDoubleBooked ", i, " ]\n",
					"For Node:     ", fmt.Sprintf("%+v", pair.schedules), "\n",
					"Expected:", fmt.Sprintf("\n%+v", pair.expected.(error).Error()), "\n",
					"Got:     ", fmt.Sprintf("\n%+v", err.Error()), "\n",
				)
			}
			continue
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
