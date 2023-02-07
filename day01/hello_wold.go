package main

import "fmt"

var a string

func main() {
	//fmt.Printf("tom\tjack\n")
	//fmt.Printf("helle\nworld")
	//var v int = Sum(1,2)
	//fmt.Println(v)
	//a ="G"
	//n()
	//m()
	//n()
	//var str string = "This is an example of a string"
	//fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "TH")
	//fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//var str string ="Hi, I'm Marc, Hi."
	//fmt.Printf("The position of the first instance of \"Hi\" is:")
	//fmt.Printf("%d\n", strings.Index(str, "Hi"))
	//fmt.Printf("The position of the last instance of \"Hi\" is: ")
	//fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))
	//fmt.Printf("The position of \"Burger\" is: ")
	//fmt.Printf("%d\n", strings.Index(str, "Burger"))
	//str := "The quick brown fox jumps over the lazy dog"
	//sl := strings.Fields(str)
	//fmt.Printf("Splitted in slice: %v\n", sl)
	//for _, val := range sl {
	//	fmt.Printf("%s - ", val) }
	//fmt.Println()
	//for _, item := range sl{
	//	fmt.Printf("%s -item ", item)
	//}
	//str2 := "GO1|The ABC of Go|25"
	//sl2 := strings.Split(str2, "|")
	//fmt.Printf("Splitted in slice: %v\n", sl2)
	//for _, val := range sl2 {
	//	fmt.Printf("%s - ", val) }
	//fmt.Println()
	//str3 := strings.Join(sl2,";")
	//fmt.Printf("sl2 joined by ;: %s\n", str3)
	//t := time.Now()
	//fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	//fmt.Printf("%d\n", time.Now())
	//var sum int = 0
	//for i := 0; i < 15; i++ {
	//	sum += i
	//	fmt.Printf("This is the %d iteration\n", i)
	//	fmt.Printf("sum is the %d iteration\n", sum)
	//}
	//
	for i := 0; i < 15; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%c", 'G')
		}
		fmt.Printf("\n")
	}

	//for i:= 1; i < 101; i++ {
	//	switch {
	//	case i%3 == 0 && i != 0 && i%5 != 0:
	//		fmt.Printf("%s\n", "Fizz")
	//	case i%3 == 0 && i%5==0 && i != 0:
	//		fmt.Printf("%s\n", "FizzBuzz")
	//	case i % 5 == 0 && i%3 != 0 && i != 0:
	//		fmt.Printf("%s\n", "Buzz")
	//	default:
	//		fmt.Printf("%d\n", i)
	//	}
	//
	//}
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}
	//s := Stack{1, 2, 3, 4}
	//i := s.Pop()
	//fmt.Println(i)

}
func Sum(a, b int) int {
	return a + b
}
func n() {
	print(a)
}
func m() {
	a := "O"
	print(a)
}

type Stack []int

func (st Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
	return 0
}
