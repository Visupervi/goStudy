package main

import (
	"bufio"
	"day28/entity"
	"fmt"
	"io"
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
	//file1, err := os.OpenFile("/Volumes/life/goStudy/src/day28/file/data.txt", os.O_WRONLY|os.O_APPEND, 0666)
	//
	//if err != nil {
	//	fmt.Printf("出错信息=%v", err)
	//	return
	//}
	//
	//str := "hello God\n"
	//writer := bufio.NewWriter(file1)
	//for i := 0; i < 10; i++ {
	//	writer.WriteString(str) // 这个时候还没有写到磁盘，这时候内容已经在缓存了
	//}
	//
	//writer.Flush() // 所以需要flush将缓存的数据真正写入到文件中
	//content, err := ioutil.ReadFile("/Volumes/life/goStudy/src/day28/file/data.txt")
	//
	//if err != nil {
	//	fmt.Println("读取文件出错")
	//}
	//
	//err = ioutil.WriteFile("/Volumes/life/goStudy/src/day28/file/data1.txt", content, 0666)
	//fileCopy("/Volumes/life/download/dream_TradingCard (1).jpg", "/Volumes/life/goStudy/src/day27/file/abc.jpg")

	//getCount("/Volumes/life/goStudy/src/day28/file/data1.txt")

	fmt.Println("命令行参数", os.Args)

	for _, arg := range os.Args {
		fmt.Printf("arg=%v\n", arg)
	}
}

// 判断文件是否存在
func pathIsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

// 拷贝文件 "/Volumes/life/download/dream_TradingCard (1).jpg"

func fileCopy(srcFile string, detFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {

	}

	reader := bufio.NewReader(src)
	defer src.Close()

	det, err := os.OpenFile(detFile, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {

	}

	writer := bufio.NewWriter(det)

	defer det.Close()

	return io.Copy(writer, reader)

}

// 统计文件中字符的个数

func getCount(path string) {

	count := entity.CharNum{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("出错啦")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		for _, v := range str {
			//fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v > '0' && v < '9':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}

		}
	}
	fmt.Printf("字符的个数=%v,数字的个数=%v,空格的个数=%v,其他字符个数=%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
