package schedule

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	schedule Schedule
	expected interface{}
}

func TestString(t *testing.T) {
	var testCases = []testCase{
		{
			schedule: Schedule{1, 2},
			expected: "[1, 2]",
		},
	}

	for i, pair := range testCases {

		result := testCases[i].schedule.String()
		if !reflect.DeepEqual(pair.expected, result) {
			t.Error(
				"[ Testcase: TestString ", i, " ]\n",
				"[ For:     ", fmt.Sprintf("%+v", pair.schedule.String()), "]\n",
				"[ Expected:", fmt.Sprintf("%+v", pair.expected), "]\n",
				"[ Got:     ", fmt.Sprintf("%+v", result), "]\n",
			)
		}
	}
}

func TestIsValid(t *testing.T) {
	var testCases = []testCase{
		{
			schedule: Schedule{1, 2},
			expected: true,
		},
		{
			schedule: Schedule{3, 2},
			expected: false,
		},
	}

	for i, pair := range testCases {

		result := testCases[i].schedule.IsValid()
		if !reflect.DeepEqual(pair.expected, result) {
			t.Error(
				"[ Testcase: TestIsValid ", i, " ]\n",
				"[ For:     ", fmt.Sprintf("%+v", pair.schedule.String()), "]\n",
				"[ Expected:", fmt.Sprintf("%+v", pair.expected), "]\n",
				"[ Got:     ", fmt.Sprintf("%+v", result), "]\n",
			)
		}
	}
}
