package main

import (
	_ "day30-3/entity"
	"fmt"
)

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan

		if !ok {
			break
		}

		fmt.Println("读到的数据=", v)
	}

	exitChan <- true
	close(exitChan)
}
func writeData(innChan chan int) {
	for i := 0; i < 50; i++ {
		innChan <- i
		fmt.Println("写数据=", i)
	}
	close(innChan)
}
func main() {

	// goroutine和channel混合使用

	intChan := make(chan int, 50)

	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan

		fmt.Println("ok=", ok)
		if ok {
			break
		}
	}

}
