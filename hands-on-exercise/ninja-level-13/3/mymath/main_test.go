package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		args           []int
		expectedResult float64
	}

	tableTest := []test{
		{
			args:           []int{1, 2, 3, 4},
			expectedResult: 2.5,
		},
		{
			args:           []int{1, 2, 3, 4, 5},
			expectedResult: 3,
		},
		{
			args:           []int{1, 2, 3, 4, 5, 6},
			expectedResult: 3.5,
		},
	}

	for _, v := range tableTest {
		if result := CenteredAvg(v.args); result != v.expectedResult {
			t.Error("Expexted", v.expectedResult, "Got", result)
		}
	}
}

func ExampleCenteredAvg() {
	args := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
	}
	for _, v := range args {
		fmt.Println(CenteredAvg(v))
	}
	// Output:
	// 2.5
	// 3
	// 3.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		args := []int{1, 2, 3, 4}
		CenteredAvg(args)
	}
}
