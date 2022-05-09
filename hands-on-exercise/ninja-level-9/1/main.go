package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	wg.Add(2)

	go Foo()
	go Bar()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}

func Foo() {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	fmt.Println("This is Foo")
}

func Bar() {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	fmt.Println("This is Bar")
}
