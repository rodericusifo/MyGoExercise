package word

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUseCount(t *testing.T) {
	args := "twinkle twinkle little star"
	expectedResult := map[string]int{
		"twinkle": 2,
		"little":  1,
		"star":    1,
	}
	if result := UseCount(args); !(reflect.DeepEqual(result, expectedResult)) {
		t.Error("Expexted", expectedResult, "Got", result)
	}
}

func TestCount(t *testing.T) {
	args := "twinkle twinkle little star"
	expectedResult := 4
	if result := Count(args); result != expectedResult {
		t.Error("Expexted", expectedResult, "Got", result)
	}
}

func ExampleCount() {
	args := "twinkle twinkle little star"
	fmt.Println(Count(args))
	// Output:
	// 4
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		args := "twinkle twinkle little star"
		UseCount(args)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		args := "twinkle twinkle little star"
		Count(args)
	}
}
