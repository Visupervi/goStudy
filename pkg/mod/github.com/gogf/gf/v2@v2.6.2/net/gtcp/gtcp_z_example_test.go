// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtcp_test

import (
	"fmt"

	"github.com/gogf/gf/v2/net/gtcp"
)

func ExampleGetFreePort() {
	fmt.Println(gtcp.GetFreePort())

	// May Output:
	// 57429 <nil>
}

func ExampleGetFreePorts() {
	fmt.Println(gtcp.GetFreePorts(2))

	// May Output:
	// [57743 57744] <nil>
}
