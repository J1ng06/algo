package btree

import (
	"algo/doublebooked/schedule"
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	schedules []schedule.Schedule
	expected  interface{}
}

func TestInsert(t *testing.T) {
	var testCases = []testCase{
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{0, 1},
				schedule.Schedule{0, 2},
			},
			expected: &Node{Schedule: schedule.Schedule{0, 1}, MaxEnd: 2, Right: &Node{Schedule: schedule.Schedule{0, 2}, MaxEnd: 2, bal: 0}, bal: 1},
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
