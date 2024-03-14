package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	// == 区分大小写
	// equals不区分大小写
	// Index
	// // 字符串函数
	fmt.Printf("b=%v\n", strings.Contains("abcccccc", "a"))
	fmt.Printf("c=%v\n", strings.Count("abcccccc", "a"))
	fmt.Printf("d=%v\n", strings.Index("abcccccc", "a"))
	fmt.Printf("d=%v\n", strings.LastIndex("abcccsaccca", "a"))
	fmt.Printf("e=%v\n", strings.Replace("abcccsaccca", "a", "dd", -1))
	fmt.Printf("s=%v\n", strings.Split("abcccsaccca", ""))
	fmt.Printf("A=%v\n", strings.ToLower("AAAAAA"))
	fmt.Printf("sa=%v\n", strings.ToUpper("abcccsaccca"))
	fmt.Printf("trim=%v\n", strings.TrimSpace(" abcccsaccca     "))
	fmt.Printf("trim=%v\n", strings.TrimRight(" abcccsaccca!", "!"))
	fmt.Printf("prefix=%v\n", strings.HasSuffix(" abcccsaccca.jpg", ".jpg"))

	// 时间日期
	fmt.Printf("time=%v\n", time.Now().UnixNano())
	fmt.Printf("time=%v\n", time.Now().Unix())
	fmt.Printf("YEAR=%v\n", time.Now().Year())
	fmt.Printf("month=%v\n", int(time.Now().Month()))
	fmt.Printf("day=%v\n", time.Now().Day())
	fmt.Printf("day=%v\n", time.Now().Format("2006-01-02 15:04:05")) // 固定， 必须这样写
	fmt.Printf("day=%v\n", time.Now().UnixNano())                    // 固定， 必须这样写
	//for {
	//	rand.Seed(time.Now().UnixNano())
	//	n := rand.Intn(100) + 1
	//	if n == 99 {
	//		break
	//	}
	//	fmt.Println(n)
	//}
	//test03()
	//textError()
	//textNew()
	//test()
	arrTest()

}

// 查看函数运行时间

func test03() {
	start := time.Now().Unix()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix() - start
	fmt.Println("函数执行花费时间=", end)
}

// new

// 用来分配内存， 主要用来分配值类型， 返回的是指针

func textNew() {
	num := new(int)

	fmt.Printf("num的类型%T, num的值%v, num的地址%v\n, num指向的值=%v\n", num, num, &num, *num)
}

// 错误处理
func textError() {
	defer func() {
		err := recover()
		if err != nil {
			println("err=", err)
		}
	}()

	num := 10
	num1 := 0

	res := num / num1

	println("输出结果", res)
	//defer recover()
}

// 自定义错误处理

func readConf(name string) (error error) {
	if name == "init.conf" {
		return nil
	} else {
		return errors.New("文件名称不能为空。。。")
	}
}

func test() {
	err := readConf("init")

	if err != nil {
		panic(err)
	}

	fmt.Println("继续执行。。。")
}

// 数组和切片

// 数组的定义

func arrTest() {
	arr := []float64{3, 5, 1, 3.4, 2, 50}
	var sum float64 = 0.0
	for _, i2 := range arr {
		//println("i = ", i)
		//println("i2 = ", i2)
		sum += i2
	}

	fmt.Println("总=", sum)
	fmt.Println("平均", fmt.Sprintf("%.2f", sum/float64(len(arr))))

}
