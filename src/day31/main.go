package main

import (
	_ "day31/entity"
	"fmt"
)

// 管道堵塞
// 只有写没有读会发生堵塞
// 如果写的快读得慢，不会发生堵塞
// 要求统计1-200000之间的素数
// 分配给四个协程去完成

func main() {

	numChan := make(chan int, 8000) // 需要计算的值
	resChan := make(chan map[int]bool, 8000)
	exitChan := make(chan bool, 8)

	// 先用一个协程写入数据

	go func() {
		for i := 1; i <= 8000; i++ {
			numChan <- i
		}

		close(numChan)
	}()

	// 然后起8个协程计算是否为素数
	for i := 0; i < 8; i++ {
		go readData(numChan, resChan, exitChan)
	}

	// 这里主线程还要操作， 当从这个管道取出8个，说明全部执行完毕，所以下一步可以关闭exitChan这个
	//go func() {
	for i := 0; i < 8; i++ {
		<-exitChan
	}

	close(exitChan)
	//}()

	close(resChan)

	//for v := range resChan {
	//	for k, v1 := range v {
	//		fmt.Printf("数字[%v]是不是素数%v\n", k, v1)
	//	}
	//}

	ch := make(chan int, 10)
	exitChan1 := make(chan struct{}, 2)

	go send(ch, exitChan1)
	go recv(ch, exitChan1)

	var total = 0

	for _ = range exitChan1 {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束。。。")

	intChan := make(chan int, 10)
	strChan := make(chan string, 5)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 在实际的开发过程中很难确定什么时候关闭管道

	// 所以使用select可以有效解决

	select {
	// 这里intChan没有关闭也不会出现deadlock,而是会自动跳转到下一个case
	case v := <-intChan:
		fmt.Printf("v%d", v)
	case v := <-strChan:
		fmt.Printf("v%s", v)
	default:
		fmt.Println("自定义数据")
		return
	}

}

func readData(numChan chan int, resChan chan map[int]bool, exitChan chan bool) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		resChan <- map[int]bool{num: flag}

		// 这里还不能关闭resChan这个管道，因为这时候可能其他协程在写数据，所以昨天第一次写的会出现写错误。
		//resChan <- true
	}

	exitChan <- true
}

// channel 使用的注意事项

// channel可以声明为只读或者只写

// 在默认的情况下是双向管道

//只写
//var chan2 chan<- int

//只读
// var chan3 <-chan int

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}

	exitChan <- a
}

func recv(ch <-chan int, exitChan chan struct{}) {

	for {
		v, ok := <-ch

		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a

}

// 使用select可以解决从管道取数据的阻塞问题
// goroutine使用recover，解决协程中出现的panic,导致程序崩溃的问题

//func selectTest(ch chan int){
//
//}
