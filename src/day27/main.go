package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件操作
// 文件的增删改查
// 文件是流的方式来操作的

func main() {
	//读文件
	//file, err := os.Open("/Volumes/life/goStudy/src/day27/file/data.txt")
	//if err != nil {
	//	fmt.Println("open file err", err)
	//}
	//fmt.Printf("file=%v\n", file)
	//defer file.Close()

	//	创建一个reader
	//  分批次读入

	//reader := bufio.NewReader(file) // 默认4096字节

	//for {
	//	str, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	//if err != nil {
	//	//	fmt.Println("读出错", err)
	//	//}
	//
	//	fmt.Print(str)
	//}

	// 一次读入 适合小文件
	// 不需要显示的打开和关闭
	//content, err := ioutil.ReadFile("/Volumes/life/goStudy/src/day27/file/data1.txt")
	//
	//if err != nil {
	//	fmt.Println("出错")
	//}
	//
	//fmt.Printf("%v", string(content))

	// 写文件
	file1, err := os.OpenFile("/Volumes/life/goStudy/src/day27/file/data2.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("出错信息=%v", err)
		return
	}

	str := "hello God\n"
	writer := bufio.NewWriter(file1)
	for i := 0; i < 10; i++ {
		writer.WriteString(str) // 这个时候还没有写到磁盘，这时候内容已经在缓存了
	}

	writer.Flush() // 所以需要flush将缓存的数据真正写入到文件中

}
