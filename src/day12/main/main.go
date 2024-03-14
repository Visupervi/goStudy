package main

import (
	"fmt"
	"math/rand"
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
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		if n == 99 {
			break
		}
		fmt.Println(n)
	}
	test03()
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
