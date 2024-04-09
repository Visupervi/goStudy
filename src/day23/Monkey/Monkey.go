package Monkey

import "fmt"

type Monkey struct {
	Name string
}

func (monkey *Monkey) Climbing() {
	fmt.Printf("%v生来就会爬树。。。\n", monkey.Name)
}

type LittleMonkey struct {
	Monkey
}

// Flying   当A结构体继承了B结构体， 那么A结构体就自动的继承了B结构体的字段和方法，并且可以直接使用
// 当A需要扩充功能，同时不希望破坏继承关系，则可以去实现某个接口即可，因此我们认为实现接口是对继承的补充
func (littleMonkey *LittleMonkey) Flying() {
	fmt.Println(littleMonkey.Name, "我要飞\n")
}

func (littleMonkey *LittleMonkey) Swimming() {
	fmt.Println(littleMonkey.Name, "我要游泳\n")
}

type Graduate struct {
}

func (g *Graduate) LearningEnglish() {
	fmt.Println("大学生学习英语")
}

type FootBool struct {
}

func (fb *FootBool) LearningEnglish() {
	fmt.Println("足球运动员学习英语")
}
