// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gutil_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gutil"
)

func Test_MapCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := gutil.MapCopy(m1)
		m2["k2"] = "v2"

		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], "v1")
		t.Assert(m2["k2"], "v2")
	})
}

func Test_MapContains(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		t.Assert(gutil.MapContains(m1, "k1"), true)
		t.Assert(gutil.MapContains(m1, "K1"), false)
		t.Assert(gutil.MapContains(m1, "k2"), false)
		m2 := g.Map{}
		t.Assert(gutil.MapContains(m2, "k1"), false)
	})
}

func Test_MapDelete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		gutil.MapDelete(m1, "k1")
		gutil.MapDelete(m1, "K1")
		m2 := g.Map{}
		gutil.MapDelete(m2, "k1")
	})
}

func Test_MapMerge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		gutil.MapMerge(m1, m2, m3, nil)
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], "v2")
		t.Assert(m1["k3"], "v3")
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
		gutil.MapMerge(nil)
	})
}

func Test_MapMergeCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		m := gutil.MapMergeCopy(m1, m2, m3, nil)
		t.Assert(m["k1"], "v1")
		t.Assert(m["k2"], "v2")
		t.Assert(m["k3"], "v3")
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapPossibleItemByKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		k, v := gutil.MapPossibleItemByKey(m, "NAME")
		t.Assert(k, "name")
		t.Assert(v, "guo")

		k, v = gutil.MapPossibleItemByKey(m, "nick name")
		t.Assert(k, "NickName")
		t.Assert(v, "john")

		k, v = gutil.MapPossibleItemByKey(m, "none")
		t.Assert(k, "")
		t.Assert(v, nil)
	})
}

func Test_MapContainsPossibleKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		t.Assert(gutil.MapContainsPossibleKey(m, "name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "NAME"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nickname"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick_name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick-name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "nick.name"), true)
		t.Assert(gutil.MapContainsPossibleKey(m, "none"), false)
	})
}

func Test_MapOmitEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"k1": "john",
			"e1": "",
			"e2": 0,
			"e3": nil,
			"k2": "smith",
		}
		gutil.MapOmitEmpty(m)
		t.Assert(len(m), 2)
		t.AssertNE(m["k1"], nil)
		t.AssertNE(m["k2"], nil)
		m1 := g.Map{}
		gutil.MapOmitEmpty(m1)
		t.Assert(len(m1), 0)
	})
}

func Test_MapToSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		s := gutil.MapToSlice(m)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[1], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[2], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[3], g.Slice{"k1", "k2", "v1", "v2"})
		s1 := gutil.MapToSlice(&m)
		t.Assert(len(s1), 4)
		t.AssertIN(s1[0], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[1], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[2], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[3], g.Slice{"k1", "k2", "v1", "v2"})
	})
	gtest.C(t, func(t *gtest.T) {
		m := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		s := gutil.MapToSlice(m)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[1], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[2], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[3], g.Slice{"k1", "k2", "v1", "v2"})
	})
	gtest.C(t, func(t *gtest.T) {
		m := g.MapStrStr{}
		s := gutil.MapToSlice(m)
		t.Assert(len(s), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		s := gutil.MapToSlice(1)
		t.Assert(s, nil)
	})
}
