package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	args := 12
	expectedResult := 84
	if result := Years(args); result != expectedResult {
		t.Error("Expexted", expectedResult, "Got", result)
	}
}

func TestYearsTwo(t *testing.T) {
	args := 12
	expectedResult := 84
	if result := YearsTwo(args); result != expectedResult {
		t.Error("Expexted", expectedResult, "Got", result)
	}
}

func ExampleYears() {
	args := 12
	fmt.Println(Years(args))
	// Output:
	// 84
}

func ExampleYearsTwo() {
	args := 12
	fmt.Println(YearsTwo(args))
	// Output:
	// 84
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		args := 12
		Years(args)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		args := 12
		YearsTwo(args)
	}
}
