package main

import (
	"algo/doublebooked/avltree"
	"algo/doublebooked/schedule"
)

func main() {
	// schedules := []schedule.Schedule{
	// 	schedule.Schedule{20, 21},
	// 	schedule.Schedule{9, 10},
	// 	schedule.Schedule{21, 25},
	// 	schedule.Schedule{8, 9},
	// 	schedule.Schedule{12, 13},
	// 	schedule.Schedule{14, 15},
	// 	schedule.Schedule{1, 2},
	// }

	// schedules := []schedule.Schedule{
	// 	schedule.Schedule{20, 24},
	// 	schedule.Schedule{30, 39},
	// 	schedule.Schedule{10, 15},
	// 	schedule.Schedule{36, 37},
	// 	schedule.Schedule{24, 28},
	// 	schedule.Schedule{20, 25},
	// }

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
		schedule.Schedule{1, 2},
	}

	tree, _ := avltree.NewTree(schedules)
	tree.Dump()
	// for i := 0; i < len(schedules); i++ {
	// 	fmt.Printf("-------- Overlap Schedules for: %s --------\n", schedules[i].String())
	// 	result := btree.Overlap(tree.Root, schedules[i])
	// 	for j := 0; j < len(result); j++ {
	// 		fmt.Println(result[j])
	// 	}
	// }
}
