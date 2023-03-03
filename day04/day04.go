package main

import "fmt"
import "strconv"

var a string
var (
	n1 = 90
	n2 = "Rose"
	n3 = "Jerry"
)

func main() {
	// 数据类型
	// 基本数据类型⬇️

	// 数值型
	// 1、整数类型（int,int8, int16, int32, int64,带符号： uint,uint8, uint16, uint32, uint64,）

	// 类型     有无符号  占用存储空间     表示范围     备注
	// int8     有       1个字节        -128～127
	// int16    有       2个字节        -2^15~ 2^15-1
	// int32    有       4个字节        -2^31~ 2^31-1
	// int64    有       8个字节        -2^63~ 2^63-1

	// 2、浮点型(float32, float64)
	// 类型       占用存储空间     表示范围     备注
	// float3     4个字节        -3.403e38~3.403e38
	// float64    8个字节        -1.798e308~1.798e308

	// 浮点型数据都是带有符号的数据， 浮点数=符号位 + 指数位 + 尾数位
	// 尾数位置可能有损失

	//var num1 float32 = -123.0000901
	//var num2 float64 = -123.0000901
	var c1 byte = 'c'
	var c2 byte = '0'

	// 	如果字符在ascii码值内，可以直接使用 byte
	//  如果大于255可以使用int
	//  go 字符使用utf-8编码 英文一个字节，中文三个字节
	//  在go 中，字符的本质是一个整数，直接输出时对应的utf-8编码的码值
	// 字符类型是可以参加运算的
	// 可以给变量赋值整数，然后按字符输出，会输出相应的unicode字
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	fmt.Printf("c1=%c c2=%c", c1, c2)
	//fmt.Println(num1)
	//fmt.Println('\n')
	//fmt.Println(num2)
	// 浮点型有固定的范围和字段长度，不受os的影响
	// 浮点型默认是float64类型
	// 浮点型常量有两种形式 1、十进制 2、科学计数形式
	// 一般情况下使用float64

	// 字符型 使用byte
	//
	// 布尔型
	// 字符串 string

	// go 没有隐式转换， 如果式数据类型转换的话必须显示操作
	// 派生/复杂数据类型⬇️

	// 指针
	// 数组
	// 结构体
	// 管道
	// 函数
	// 切片
	// 接口
	// map
	/**
	 * 基本数据类型转string
	 * string转基本类型
	 */
	// 方式一
	var num int = 99
	var num2 float64 = 23.56
	//var b bool = true
	var str string
	str = fmt.Sprintf("%d", num)
	fmt.Printf("str type %T str=%v\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%v\n", str, str)

	// 方式二
	str = strconv.FormatInt(int64(99), 10)
	fmt.Printf("str type %T str=%v\n", str, str)

	v := int64(-42)
	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)
	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)

	// string 转成基本数据类型

	str1 := string("false")
	b := bool(true)
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T b=%v\n", b, b)

	str2 := string("12121112212")
	b1 := int64(0)
	b1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("b1 type %T b1=%v", b1, b1)

}
