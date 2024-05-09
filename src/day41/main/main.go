package main

import (
	"day40/main/link"
	"day40/main/queue"
	"fmt"
)

// 稀疏数组的处理方式（棋盘，地图）
// 记录数组中的几行几列，有多少个不同的值
// 把具有不同的值的元素的行列值记录在一个小规模的数组中，从而缩小程序的规模
// sparsearray

func main() {
	//var chessMap [11][11]int
	//var sparseArr []sparsearray.Sparsearray
	//chessMap[1][2] = 1
	//chessMap[2][3] = 2
	//
	//sparse := sparsearray.Sparsearray{
	//	Row: 11,
	//	Col: 11,
	//	Val: 0,
	//}
	//sparseArr = append(sparseArr, sparse)
	//for i1, v1 := range chessMap {
	//	for i2, v2 := range v1 {
	//		fmt.Printf("%d", v2)
	//		if v2 != 0 {
	//			sparse = sparsearray.Sparsearray{
	//				Row: i1,
	//				Col: i2,
	//				Val: v2,
	//			}
	//			sparseArr = append(sparseArr, sparse)
	//		}
	//	}
	//
	//	fmt.Println("")
	//}
	//fmt.Println(sparseArr)
	que := &queue.Que{
		Head:    -1,
		Tail:    -1,
		MaxSize: 5,
	}
	que.Insert(5)
	que.Insert(10)
	que.Insert(11)
	que.PrintQue()
	res, err := que.GetQue()
	if err != nil {
		return
	}

	fmt.Println(res)

	head := &link.LinkNode{}

	linkNode := &link.LinkNode{
		Data: "hero",
	}

	linkNode1 := &link.LinkNode{
		Data: "hero1",
	}
	// 遍历chessMap，找到数据放入结构体中

	link.InsertListNode(head, linkNode)
	link.InsertPosLink(head, 1, linkNode1)

	link.LinkPrint(head)
}
