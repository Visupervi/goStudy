package _interface

// IUsb 不能包含任何变量
// 只包含方法，但是不实现
// 接口可以继承接口，
// 可以把任何数据类型赋值给空接口
// 在实现接口的时候只看方法有没有实现

type IUsb interface {
	Start()
	Stop()
}

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test02()
}
