package main

import (
	"algo/doublebooked/avltree"
	"algo/doublebooked/schedule"
	"algo/doublebooked/utils"
	"fmt"
	"log"
)

func main() {

	// schedules := []schedule.Schedule{
	// 	schedule.Schedule{20, 24},
	// 	schedule.Schedule{30, 35},
	// 	schedule.Schedule{10, 15},
	// 	schedule.Schedule{36, 37},
	// 	schedule.Schedule{20, 28},
	// 	schedule.Schedule{20, 25},
	// 	schedule.Schedule{8, 9},
	// 	schedule.Schedule{13, 14},
	// 	schedule.Schedule{14, 15},
	// }

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

	// schedules := []schedule.Schedule{
	// 	schedule.Schedule{21, 24},
	// 	schedule.Schedule{20, 24},
	// 	schedule.Schedule{20, 24},
	// }

	for _, v := range DoubleBooked(schedules) {
		fmt.Println(v)
	}

}

func DoubleBooked(schedules []schedule.Schedule) (out [][2]schedule.Schedule) {

	if len(schedules) == 0 {
		log.Fatal("Invalid Input")
	}
	tree := new(avltree.Tree)
	union := make(map[[2]schedule.Schedule]interface{})
	for i := 0; i < len(schedules); i++ {
		err, dup := tree.Insert(schedules[i])
		if err != nil {
			log.Fatal("Failed to build tree ", err.Error())
		}

		if dup {
			union[[2]schedule.Schedule{schedules[i], schedules[i]}] = new(interface{})
		}
	}

	for i := 0; i < len(schedules); i++ {
		union = utils.MapUnion(union, avltree.Overlap(tree.Root, schedules[i]))
	}

	for k := range union {
		out = append(out, k)
	}
	return
}
