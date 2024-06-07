// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcron_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gcron"
)

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcron.Add(ctx, "1 1 1 1 1 1", func(ctx context.Context) {

		})
	}
}
