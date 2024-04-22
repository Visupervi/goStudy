package main

import (
	"day33/entity"
	"fmt"
	"reflect"
)

// 反射最佳实践
// 使用反射来遍历结构体字段，调用结构体的方法，并获取结构体标签的值

//func (Value) Method
//func (v Value) Method(i int) Value
//返回v持有值类型的第i个方法的已绑定（到v的持有值的）状态的函数形式的Value封装。
//返回值调用Call方法时不应包含接收者；返回值持有的函数总是使用v的持有者作为接收者（即第一个参数）。
//如果i出界，或者v的持有值是接口类型的零值（nil），会panic。

func main() {
	//num := 3.12
	//
	//test(num)
	//
	//fmt.Println(num)
	//const r = 3.14
	monster := entity.Monster{
		Name:   "黑风怪",
		Age:    300,
		Score:  90.12,
		Gender: "男",
	}

	testStruct(monster)

}

func testStruct(i interface{}) {
	//Type类型用来表示一个go类型。
	//
	//不是所有go类型的Type值都能使用所有方法。请参见每个方法的文档获取使用限制。在调用有分类限定的方法时，应先使用Kind方法获知类型的分类。调用该分类不支持的方法会导致运行时的panic。
	//
	//func TypeOf
	//func TypeOf(i interface{}) Type
	//TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil。

	fmt.Println(reflect.TypeOf(i))

	typ := reflect.TypeOf(i)
	fmt.Println(reflect.ValueOf(i))
	rVal := reflect.ValueOf(i)
	kd := rVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := rVal.NumField() // 获取一共有几个字段

	fmt.Println(num)

	for i := 0; i < num; i++ {

		//	获取得到struct标签， 注意需要通过reflect.Type来获取Tag的值

		fmt.Printf("Filed %d. 值为%v", i, rVal.Field(i))
		tag := typ.Field(i).Tag.Get("json")

		if tag != "" {
			fmt.Printf("Filed %d: tag为%v\n", i, tag)
		}

	}

	// 获取结构体的方法

	methods := rVal.NumMethod()

	fmt.Println("methods=", methods)

	// 使用反射调用方法

	// 方法排序是按照ASCII码的排序来的

	//rVal.Method(1).Call(nil)

	// 声明一个[]reflect.ValueOf的切片

	var params []reflect.Value

	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))

	res := rVal.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

}

func test(i interface{}) {
	//	将interface转成reflect.Value
	rVal := reflect.ValueOf(i)

	fmt.Println(rVal.Kind())
	fmt.Println(rVal.Type())
	iVal := rVal.Interface()

	ii := iVal.(float64)

	fmt.Println(ii)
	//rVal.Elem().SetInt(20)

}
