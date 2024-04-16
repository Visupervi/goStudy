package channel

import (
	"fmt"
	"sync"
)

// 并发和并行

// 并发说的是单颗核心
// 并行是多颗核心上

// goroutine 协程
// 1. 一个go线程上可以起多个协程
// 2. 协程的特点 有独立的栈空间， 共享程序堆空间，调度由用户控制，协程是轻量级的线程
// 和主线程如何通信

// goroutine的调度模型 MPG M: 操作系统的主线程（物理线程）P：协程执行需要的上下文 G：协程

// MPG 动态

var (
	myMap = make(map[int]int)

	// 声明一个互斥锁
	//

	lock sync.Mutex
)

func test(n int) {
	//for i := 0; i < 10; i++ {
	//	fmt.Println("Hello Goroutine" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	//go test() // go开启一个协程

	// 主线程执行完毕，就不在执行协程
	// 协程执行完不影响主线程执行

	//for i := 0; i < 10; i++ {
	//	fmt.Println("Hello main" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}

	//num := runtime.NumCPU()
	//
	//fmt.Println("num=", num)
	for i := 1; i <= 20; i++ {
		go test(i)
		//time.Sleep(time.Second)
	}
	//time.Sleep(time.Second * 10) // 这种方式会出现资源竞争的问题，资源互斥

	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
	lock.Unlock()
}
