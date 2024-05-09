package queue

import (
	"errors"
	"fmt"
)

type Que struct {
	Head int
	Tail int
	//Data interface{}
	MaxSize int
	Arr     [5]interface{}
}

func (q *Que) Insert(data interface{}) (err error) {
	// 先判断队列是否已满
	if q.Tail == q.MaxSize-1 {
		return errors.New("队列已满")
	}
	q.Tail++
	q.Arr[q.Tail] = data
	return
}

func (q *Que) PrintQue() {
	for i := q.Head + 1; i <= q.Tail; i++ {
		fmt.Println(q.Arr[i])
	}
}

func (q *Que) GetQue() (val interface{}, err error) {
	if q.Tail == q.Head {
		return val, errors.New("队列为空")
	}

	q.Head++
	val = q.Arr[q.Head]
	return val, err
}

func (q *Que) InsertCircleQue(data interface{}) (err error) {
	// 先判断队列是否已满
	if ((q.Tail + 1) % q.MaxSize) == q.Head {
		return errors.New("队列已满")
	}
	q.Arr[q.Tail] = data
	q.Tail++
	return
}

func (q *Que) PrintCircleQue() {
	for i := q.Head + 1; i <= q.Tail; i++ {
		fmt.Println(q.Arr[i])
	}
}

func (q *Que) GetCircleQue() (val interface{}, err error) {
	if q.Tail == q.Head {
		return val, errors.New("队列为空")
	}

	val = q.Arr[q.Head]
	q.Head++
	return val, err
}

// 统计数据（tail + maxsize - head） % maxsize

func (q *Que) CircleSize() int {
	return (q.Tail + q.MaxSize - q.Head) % q.MaxSize
}
