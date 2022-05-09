package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go Foo()
	Bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func Foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func Bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}
