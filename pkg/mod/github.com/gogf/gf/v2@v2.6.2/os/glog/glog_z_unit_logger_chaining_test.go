// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
)

func Test_To(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		To(w).Error(ctx, 1, 2, 3)
		To(w).Errorf(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.Count(w.String(), defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(w.String(), "1 2 3"), 2)
	})
}

func Test_Path(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
	})
}

func Test_Cat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cat := "category"
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Cat(cat).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Cat(cat).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, cat, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
	})
}

func Test_Level(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Level(LEVEL_PROD).Stdout(false).Debug(ctx, 1, 2, 3)
		Path(path).File(file).Level(LEVEL_PROD).Stdout(false).Debug(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 0)
		t.Assert(gstr.Count(content, "1 2 3"), 0)
	})
}

func Test_Skip(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Skip(10).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
		//t.Assert(gstr.Count(content, "Stack"), 1)
	})
}

func Test_Stack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Stack(false).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
		//t.Assert(gstr.Count(content, "Stack"), 1)
	})
}

func Test_StackWithFilter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).StackWithFilter("none").Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(ctx, content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		//t.Assert(gstr.Count(content, "Stack"), 1)

	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).StackWithFilter("/gf/").Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(ctx, content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		//t.Assert(gstr.Count(content, "Stack"), 0)
	})
}

func Test_Header(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Header(true).Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Header(false).Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 0)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})
}

func Test_Line(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Line(true).Stdout(false).Debug(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		//t.Assert(gstr.Count(content, ".go"), 1)
		//t.Assert(gstr.Contains(content, gfile.Separator), true)
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Line(false).Stdout(false).Debug(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		//t.Assert(gstr.Count(content, ".go"), 1)
		//t.Assert(gstr.Contains(content, gfile.Separator), false)
	})
}

func Test_Async(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Async().Stdout(false).Debug(ctx, 1, 2, 3)
		time.Sleep(1000 * time.Millisecond)

		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})

	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Async(false).Stdout(false).Debug(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})
}
