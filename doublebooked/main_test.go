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
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{20, 21},
				schedule.Schedule{9, 10},
				schedule.Schedule{21, 25},
				schedule.Schedule{8, 9},
				schedule.Schedule{12, 13},
				schedule.Schedule{14, 15},
				schedule.Schedule{1, 2},
			},
			expected: [][2]schedule.Schedule{},
		},
		{
			schedules: []schedule.Schedule{
				schedule.Schedule{20, 24},
				schedule.Schedule{30, 39},
				schedule.Schedule{10, 15},
				schedule.Schedule{36, 37},
				schedule.Schedule{24, 28},
				schedule.Schedule{20, 25},
			},
			expected: [][2]schedule.Schedule{
				{schedule.Schedule{20, 24}, schedule.Schedule{20, 25}},
				{schedule.Schedule{20, 25}, schedule.Schedule{24, 28}},
				{schedule.Schedule{30, 39}, schedule.Schedule{36, 37}},
			},
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
				schedule.Schedule{5, 999},
				schedule.Schedule{1, 2},
			},
			expected: [][2]schedule.Schedule{
				{schedule.Schedule{5, 999}, schedule.Schedule{5, 999}},
				{schedule.Schedule{5, 999}, schedule.Schedule{20, 25}},
				{schedule.Schedule{5, 999}, schedule.Schedule{20, 24}},
				{schedule.Schedule{5, 999}, schedule.Schedule{30, 35}},
				{schedule.Schedule{5, 999}, schedule.Schedule{10, 15}},
				{schedule.Schedule{5, 999}, schedule.Schedule{20, 28}},
				{schedule.Schedule{5, 999}, schedule.Schedule{36, 1000}},
				{schedule.Schedule{5, 999}, schedule.Schedule{13, 14}},
				{schedule.Schedule{5, 999}, schedule.Schedule{14, 15}},
				{schedule.Schedule{5, 999}, schedule.Schedule{8, 9}},

				{schedule.Schedule{20, 24}, schedule.Schedule{20, 28}},
				{schedule.Schedule{20, 24}, schedule.Schedule{20, 25}},

				{schedule.Schedule{20, 25}, schedule.Schedule{20, 28}},

				{schedule.Schedule{10, 15}, schedule.Schedule{13, 14}},
				{schedule.Schedule{10, 15}, schedule.Schedule{14, 15}},
			},
		},
	}

	for i, pair := range testCases {
		results := DoubleBooked(pair.schedules)
		resultsMap := make(map[[2]schedule.Schedule]interface{})
		for _, v := range results {
			resultsMap[v] = new(interface{})
		}
		for _, v := range pair.expected.([][2]schedule.Schedule) {
			delete(resultsMap, v)
		}
		if len(pair.expected.([][2]schedule.Schedule)) != len(results) || len(resultsMap) > 0 {
			t.Error(
				"[ Testcase: TestDoubleBooked ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%+v", pair.schedules), "\n",
				"Expected:", fmt.Sprintf("\n%+v", pair.expected.([][2]schedule.Schedule)), "\n",
				"Got:     ", fmt.Sprintf("\n%+v", results), "\n",
			)
		}
	}
}
