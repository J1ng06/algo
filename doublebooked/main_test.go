package main

import (
	"algo/doublebooked/schedule"
	"fmt"
	"testing"
)

type testDoubleBookedPair struct {
	schedules []schedule.Schedule
	expected  interface{}
}

func TestDoubleBooked(t *testing.T) {
	var testCases = []testDoubleBookedPair{
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{21, 24},
				schedule.Schedule{20, 24},
				schedule.Schedule{20, 24},
			},
			expected: [][2]schedule.Schedule{
				{schedule.Schedule{20, 24}, schedule.Schedule{20, 24}},
				{schedule.Schedule{20, 24}, schedule.Schedule{21, 24}},
			},
		},
	}

	for i, pair := range testCases {
		results := DoubleBooked(pair.schedules)
		if len(pair.expected.([][2]schedule.Schedule)) != len(results) {
			t.Error(
				"[ Testcase: TestDoubleBooked ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%+v", pair.schedules), "\n",
				"Expected:", fmt.Sprintf("\n%+v", pair.expected.([][2]schedule.Schedule)), "\n",
				"Got:     ", fmt.Sprintf("\n%+v", results), "\n",
			)
		}
		resultsMap := make(map[[2]schedule.Schedule]interface{})
		for _, v := range results {
			resultsMap[v] = new(interface{})
		}
		for _, v := range pair.expected.([][2]schedule.Schedule) {
			delete(resultsMap, v)
		}
		if len(resultsMap) > 0 {
			if len(pair.expected.([][2]schedule.Schedule)) != len(results) {
				t.Error(
					"[ Testcase: TestDoubleBooked ", i, " ]\n",
					"For Node:     ", fmt.Sprintf("%+v", pair.schedules), "\n",
					"Expected:", fmt.Sprintf("\n%+v", pair.expected.([][2]schedule.Schedule)), "\n",
					"Got:     ", fmt.Sprintf("\n%+v", results), "\n",
				)
			}

		}
	}
}
