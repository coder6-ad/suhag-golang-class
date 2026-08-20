package main

import (
	"fmt"
	"time"
)
func sudip(){fmt.Println("hello")}
func ayush(){fmt.Println("world")}

func chup_ram() {
	fmt.Println("kati bole ko")
}

func test_time() {
	fmt.Println("Hello world")
}

func main() {
	fmt.Println(time.Second)
	fmt.Println(time.Nanosecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Millisecond)


	time.Sleep(10 * time.Second)
	test_time()

	fmt.Println("Write now before")
	go chup_ram()
	fmt.Println("Tespachi")
	start := time.Now()
	go sudip()
	fmt.Println("world")
	go ayush()
	fmt.Println("hello")
	end :=t.Since(start)
}
