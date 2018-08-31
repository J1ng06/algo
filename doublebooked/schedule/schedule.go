package schedule

import (
	"fmt"
)

type Schedule struct {
	Start int
	End   int
}

func (s Schedule) String() string {
	return fmt.Sprintf("[%d, %d]", s.Start, s.End)
}

func (s Schedule) IsValid() bool {
	return s.Start < s.End
}

func Flattern(schedule []Schedule) (out string) {

	for i := 0; i < len(schedule); i++ {
		out += fmt.Sprintf("%s ", schedule[i].String())
	}
	return out
}

func SortSchedules(schedules [2]Schedule) (out [2]Schedule) {

	out = [2]Schedule{
		schedules[0],
		schedules[1],
	}

	if schedules[0].Start > schedules[1].Start || schedules[0].Start == schedules[1].Start && schedules[0].End > schedules[1].End {
		out[0] = schedules[1]
		out[1] = schedules[0]
	}

	return
}
