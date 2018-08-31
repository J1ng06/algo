package utils

import (
	"fmt"
	"testing"
)

type testMax struct {
	integers []int
	expected interface{}
}

func TestMax(t *testing.T) {
	var testCases = []testMax{
		{
			integers: []int{0},
			expected: 0,
		},
		{
			integers: []int{3, 2, 3},
			expected: 3,
		},
		{
			integers: []int{2, 4},
			expected: 4,
		},
	}

	for i, pair := range testCases {
		result := Max(pair.integers...)
		if result != pair.expected {
			t.Error(
				"[ Testcase: TestMax ", i, " ]\n",
				"For Node:     ", fmt.Sprintf("%+v", pair.integers), "\n",
				"Expected:", fmt.Sprintf("\n%v", pair.expected), "\n",
				"Got:     ", fmt.Sprintf("\n%v", result), "\n",
			)
		}
	}

}
