package main

import (
	"algo/doublebooked/intervaltree"
	"algo/doublebooked/schedule"
	"algo/doublebooked/utils"
	"errors"
	"fmt"
)

func main() {

	schedules := []schedule.Schedule{
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
		schedule.Schedule{5, 999},
		schedule.Schedule{1, 2},
	}

	if results, err := DoubleBooked(schedules); err != nil {
		for v := range results {
			fmt.Println(v)
		}
	}

}

func DoubleBooked(schedules []schedule.Schedule) (out [][2]schedule.Schedule, err error) {

	if len(schedules) == 0 {
		return nil, errors.New("Invalid Input")
	}
	tree := new(intervaltree.Tree)
	union := make(map[[2]schedule.Schedule]interface{})
	for i := 0; i < len(schedules); i++ {

		var dup bool
		if dup, err = tree.Insert(schedules[i]); err != nil {
			return nil, err
		}

		if dup {
			union[[2]schedule.Schedule{schedules[i], schedules[i]}] = new(interface{})
		}
	}

	for i := 0; i < len(schedules); i++ {
		union = utils.MapUnion(union, intervaltree.Overlap(tree.Root, schedules[i]))
	}

	for k := range union {
		out = append(out, k)
	}
	return
}
