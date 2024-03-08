package main

import (
	"fmt"
	"time"
)

func main() {
	//////////////////////////////////// append 函数 ////////////////////////////////////
	appendFunc()

	//////////////////////////////////// copy 函数 ////////////////////////////////////
	copyFunc()

	//////////////////////////////////// close 函数 ////////////////////////////////////
	closeFunc()

	//////////////////////////////////// delete 函数 ////////////////////////////////////
	deleteFunc()

	//////////////////////////////////// new 函数 ////////////////////////////////////
	newFunc()

	//////////////////////////////////// panic recover 函数 ////////////////////////////////////
	panicRecoverFunc()
}

func panicRecoverFunc() {
	defer func() {
		fmt.Println("panicRecoverFunc execute defer")
	}()

	go func() {
		defer func() {
			fmt.Println("goroutine execute defer")
			// 获取panic进行处理
			if x := recover(); x != nil {
				fmt.Println("recover from panic:", x)
			}
		}()

		panic("test go panic")
	}()
	time.Sleep(time.Second)
	fmt.Println("panicRecoverFunc complete")
}

func newFunc() {
	type S struct {
		a int
		b string
	}
	s := new(S)
	fmt.Printf("new var s is %v of type %T\n", s, s)
}

func deleteFunc() {
	m1 := make(map[string]int)
	m1["k1"] = 1
	m1["k2"] = 2
	fmt.Println("map m1: ", m1)

	delete(m1, "k1")
	delete(m1, "k3")
	delete(m1, "")
	fmt.Println("map m1: ", m1)

	var m2 map[string]int
	delete(m2, "k")
	fmt.Println("map m2: ", m2)
}

func closeFunc() {
	// 创建
	ch1 := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
		close(ch1)
		fmt.Println("channel ch1 closed")
	}()

	for v := range ch1 {
		time.Sleep(time.Second)
		fmt.Println("received value: ", v)
	}
	fmt.Println("channel ch1 is already closed")

	v, ok := <-ch1
	fmt.Println("receive data from closed channel, ok is false: ", v, ok)

	// 重复关闭closed的channel会导致程序错误
	// close(ch1) // panic: close of closed channel

	// 关闭nil的channel 会导致报错
	// var ch2 chan int
	// close(ch2) // panic: close of nil channel

	// 创建一个只读channel，并尝试关闭它，会导致报错
	// readCh := make(<-chan int)
	// close(readCh) // 编译报错
}

func copyFunc() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	var d1 []int = make([]int, len(s1)-2)
	var d2 []int = make([]int, len(s1), len(s1)+2)
	copy(d1, s1)
	copy(d2, s1)
	fmt.Printf("s1: len=%v, cap=%v, values=%v\n", len(s1), cap(s1), s1)
	fmt.Printf("d1: len=%v, cap=%v, values=%v\n", len(d1), cap(d1), d1)
	fmt.Printf("d2: len=%v, cap=%v, values=%v\n", len(d2), cap(d2), d2)

	d1[0] = 99
	d2[0] = 999
	fmt.Printf("s1: len=%v, cap=%v, values=%v\n", len(s1), cap(s1), s1)
	fmt.Printf("d1: len=%v, cap=%v, values=%v\n", len(d1), cap(d1), d1)
	fmt.Printf("d2: len=%v, cap=%v, values=%v\n", len(d2), cap(d2), d2)
}

func appendFunc() {
	var s []int
	s1 := append(s, 0, 1, 2)
	s2 := append(s1, 3, 4, 5, 6, 7, 8)
	fmt.Printf("s: len=%v, cap=%v, values=%v\n", len(s), cap(s), s)
	fmt.Printf("s1: len=%v, cap=%v, values=%v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2: len=%v, cap=%v, values=%v\n", len(s2), cap(s2), s2)
}
