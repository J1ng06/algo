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
