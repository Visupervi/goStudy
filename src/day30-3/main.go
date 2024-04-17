package main

import (
	"day21/Hero"
	_ "day30-3/entity"
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

// 通过使用channel解决互斥问题

// channel的特点
// channel本质就是一个队列，先进先出。
// 数据是先进先出
// 线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
// channel有类型的， 一个string的channel只能存放string类型

// channel的基本介绍
// channel的基本用法
//
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
	//for i := 1; i <= 20; i++ {
	//	go test(i)
	//	//time.Sleep(time.Second)
	//}
	////time.Sleep(time.Second * 10) // 这种方式会出现资源竞争的问题，资源互斥
	//
	//lock.Lock()
	//for k, v := range myMap {
	//	fmt.Printf("map[%d]=%d\n", k, v)
	//}
	//lock.Unlock()

	// channel的基本用法

	//var intChan chan int
	//var mapChan chan map[int]string
	//var charNumChan chan entity.CharNum
	//var charNumChan2 chan *entity.CharNum
	//
	//fmt.Println(intChan)
	//fmt.Println(mapChan)
	//fmt.Println(charNumChan)
	//fmt.Println(charNumChan2)

	// 说明

	// channel 是引用类型
	// channel必须初始化才能写入数据，即make后才能使用
	// 管道是有类型的，intChan只能写入整数int

	var intChan chan int

	intChan = make(chan int, 3)

	fmt.Println(intChan)
	// 向管道写入数据

	intChan <- 12
	num := 985
	intChan <- 211
	intChan <- num
	//  注意⚠️ 当我们给管道写入数据时，不能超过其容量
	// channel只能存放指定的数据类型
	// channel的数据放满后就不能在放入了
	// 如果从channel取出数据后，可继续放入
	// 在没有使用协程的情况下， 如果channel数据取完了，再取就会报dead lock

	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))

	var num2 int

	num2 = <-intChan

	fmt.Println(num2)
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	//fmt.Println(intChan)

	mapChan := make(chan map[string]string, 10)

	m1 := make(map[string]string, 20)

	m1["city1"] = "北京"
	m1["city2"] = "上海"

	m2 := make(map[string]string, 20)

	m2["hero1"] = "武松"
	m2["hero2"] = "宋江"

	mapChan <- m1
	mapChan <- m2

	//map11 := <-mapChan
	//map22 := <-mapChan

	//fmt.Println(map11)
	//fmt.Println(map22)

	// 存放任意类型数据channel
	allChan := make(chan interface{}, 10)

	hero1 := Hero.Hero{Name: "宋江", Age: 12}
	hero2 := Hero.Hero{Name: "武松", Age: 18}

	allChan <- hero1
	allChan <- hero2
	allChan <- 10
	allChan <- "str"

	h1 := <-allChan

	// 使用断言去判断再获取

	if h, ok := h1.(Hero.Hero); ok {
		fmt.Println(h.Name)
	}
	// 如果直接获取Name会报错，所以使用断言来解决这个问题
	h11 := h1.(Hero.Hero)
	fmt.Println(h11.Age)

	// channel的关闭

	// 使用内置函数close可以关闭channel， 当channel关闭后，就不能再向channel写入数据，但仍然可以从channel读取数据

	// channel遍历
	// channel支持for-range方式进行遍历，但要注意两个细节
	// 在遍历时，如果channel没有关闭，则会出现deadlock错误
	// 在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

	intChan1 := make(chan int, 3)

	intChan1 <- 100
	intChan1 <- 200

	close(intChan1) // 关闭后只能读不能写

	//intChan1 <- 3000

	n1 := <-intChan1

	fmt.Println(n1)

}
