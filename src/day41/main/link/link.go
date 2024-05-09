package link

import "fmt"

type LinkNode struct {
	Next *LinkNode
	Prev *LinkNode
	Data interface{}
}

func InsertListNode(headNode *LinkNode, listNode *LinkNode) {
	p := headNode

	fmt.Println("listNode", listNode)
	for {
		if p.Next == nil {
			break
		}
		p = p.Next
	}
	p.Next = listNode

	//fmt.Println("listNode",)
}

func LinkPrint(headNode *LinkNode) {
	q := headNode
	fmt.Println("headNode=", headNode)
	if q.Next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("Data=%s", q.Next.Data)
		q = q.Next
		if q.Next == nil {
			break
		}
	}
}

// InsertPosLink InsertPosLink在pos位置插入
func InsertPosLink(head *LinkNode, pos int, linkNode *LinkNode) {
	p := head
	q := 0
	for {
		if p.Next == nil {
			p.Next = linkNode
			break
		}
		if q == pos {
			temp := p.Next
			p.Next = linkNode
			p.Next.Next = temp
			break
		}
		p = p.Next
		q++

	}
	p = p.Next
}
