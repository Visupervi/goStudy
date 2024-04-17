package main

import (
	_ "day30-3/entity"
	"fmt"
)

//golang 1. 启动一个协程，将1-2000的数放入到一个channel中，比如numChan 2. 启动八个协程，从numChan 取出数(比如n)，并计算1+...+n的值
func main() {

	// goroutine和channel混合使用
	numChan := make(chan int, 2000) // 需要计算的值

	resChan := make(chan map[int]int, 2000)
	exitChan := make(chan bool, 8)
	//resulChan := make(chan bool, 8)
	go func() {
		for i := 1; i <= 2000; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for i := 0; i < 8; i++ {
		go readData(numChan, resChan, exitChan)
	}

	for i := 0; i < 8; i++ {
		<-exitChan
	}
	close(exitChan)

	close(resChan)
	for v := range resChan {
		for k, v1 := range v {
			fmt.Printf("K=[%v]=%v\n", k, v1)
		}
	}

}

func readData(numChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		v, ok := <-numChan
		if !ok {
			break
		}

		res := 0

		for i := 1; i <= v; i++ {
			res += i
		}

		resChan <- map[int]int{v: res}
	}
	exitChan <- true
}
