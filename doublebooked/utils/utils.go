package utils

import "algo/doublebooked/schedule"

func Max(integers ...int) int {
	max := integers[0]

	for i := 1; len(integers) > 1 && i < len(integers); i++ {
		if integers[i] > max {
			max = integers[i]
		}
	}
	return max
}

func MapUnion(maps ...map[[2]schedule.Schedule]interface{}) (union map[[2]schedule.Schedule]interface{}) {

	union = make(map[[2]schedule.Schedule]interface{})

	for i := range maps {
		for k, v := range maps[i] {
			union[k] = v
		}
	}

	return

}
