// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gqueue_test

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/os/gtimer"
)

func ExampleNew() {
	n := 10
	q := gqueue.New()

	// Producer
	for i := 0; i < n; i++ {
		q.Push(i)
	}

	// Close the queue in three seconds.
	gtimer.SetTimeout(context.Background(), time.Second*3, func(ctx context.Context) {
		q.Close()
	})

	// The consumer constantly reads the queue data.
	// If there is no data in the queue, it will block.
	// The queue is read using the queue.C property exposed
	// by the queue object and the selectIO multiplexing syntax
	// example:
	// for {
	//    select {
	//        case v := <-queue.C:
	//            if v != nil {
	//                fmt.Println(v)
	//            } else {
	//                return
	//            }
	//    }
	// }
	for {
		if v := q.Pop(); v != nil {
			fmt.Print(v)
		} else {
			break
		}
	}

	// Output:
	// 0123456789
}

func ExampleQueue_Push() {
	q := gqueue.New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	// Output:
	// 0
	// 1
	// 2
}

func ExampleQueue_Pop() {
	q := gqueue.New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	fmt.Println(q.Pop())
	q.Close()
	fmt.Println(q.Pop())

	// Output:
	// 0
	// <nil>
}

func ExampleQueue_Close() {
	q := gqueue.New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	time.Sleep(time.Millisecond)
	q.Close()

	fmt.Println(q.Len())
	fmt.Println(q.Pop())

	// May Output:
	// 0
	// <nil>
}

func ExampleQueue_Len() {
	q := gqueue.New()

	q.Push(1)
	q.Push(2)

	fmt.Println(q.Len())

	// May Output:
	// 2
}

func ExampleQueue_Size() {
	q := gqueue.New()

	q.Push(1)
	q.Push(2)

	// Size is alias of Len.
	fmt.Println(q.Size())

	// May Output:
	// 2
}
