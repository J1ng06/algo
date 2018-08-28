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
