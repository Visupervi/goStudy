// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*" -benchmem

package gdebug

import (
	"runtime"
	"runtime/debug"
	"testing"
)

func Benchmark_BinVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinVersion()
	}
}

func Benchmark_BinVersionMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinVersionMd5()
	}
}

func Benchmark_RuntimeCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.Caller(0)
	}
}

func Benchmark_RuntimeFuncForPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.FuncForPC(11010101)
	}
}

func Benchmark_callerFromIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callerFromIndex(nil)
	}
}

func Benchmark_Stack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Stack()
	}
}

func Benchmark_StackOfStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		debug.Stack()
	}
}

func Benchmark_StackWithFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StackWithFilter([]string{"test"})
	}
}

func Benchmark_Caller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Caller()
	}
}

func Benchmark_CallerWithFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerWithFilter([]string{"test"})
	}
}

func Benchmark_CallerFilePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFilePath()
	}
}

func Benchmark_CallerDirectory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerDirectory()
	}
}

func Benchmark_CallerFileLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFileLine()
	}
}

func Benchmark_CallerFileLineShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFileLineShort()
	}
}

func Benchmark_CallerFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerFunction()
	}
}

func Benchmark_CallerPackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallerPackage()
	}
}
