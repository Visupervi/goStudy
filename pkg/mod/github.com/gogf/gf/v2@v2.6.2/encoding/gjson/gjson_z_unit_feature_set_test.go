// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gjson_test

import (
	"bytes"
	"testing"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
)

func Test_Set1(t *testing.T) {
	e := []byte(`{"k1":{"k11":[1,2,3]},"k2":"v2"}`)
	p := gjson.New(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.Set("k1.k11", []int{1, 2, 3})
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, []byte(`{"k1":{"k11":[1,2,3]},"k2":"v2"}`)) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		e := `[[null,1]]`
		p := gjson.New([]string{"a"})
		p.Set("0.1", 1)
		s := p.MustToJsonString()
		t.Assert(s, e)
	})
}

func Test_Set3(t *testing.T) {
	e := []byte(`{"kv":{"k1":"v1"}}`)
	p := gjson.New([]string{"a"})
	p.Set("kv", map[string]string{
		"k1": "v1",
	})
	if c, err := p.ToJson(); err == nil {
		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set4(t *testing.T) {
	e := []byte(`["a",[{"k1":"v1"}]]`)
	p := gjson.New([]string{"a"})
	p.Set("1.0", map[string]string{
		"k1": "v1",
	})
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set5(t *testing.T) {
	e := []byte(`[[[[[[[[[[[[[[[[[[[[[1,2,3]]]]]]]]]]]]]]]]]]]]]`)
	p := gjson.New([]string{"a"})
	p.Set("0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0", []int{1, 2, 3})
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set6(t *testing.T) {
	e := []byte(`["a",[1,2,3]]`)
	p := gjson.New([]string{"a"})
	p.Set("1", []int{1, 2, 3})
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set7(t *testing.T) {
	e := []byte(`{"0":[null,[1,2,3]],"k1":"v1","k2":"v2"}`)
	p := gjson.New(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.Set("0.1", []int{1, 2, 3})
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set8(t *testing.T) {
	e := []byte(`{"0":[[[[[[null,[1,2,3]]]]]]],"k1":"v1","k2":"v2"}`)
	p := gjson.New(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.Set("0.0.0.0.0.0.1", []int{1, 2, 3})
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set9(t *testing.T) {
	e := []byte(`{"k1":[null,[1,2,3]],"k2":"v2"}`)
	p := gjson.New(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	p.Set("k1.1", []int{1, 2, 3})
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set10(t *testing.T) {
	e := []byte(`{"a":{"b":{"c":1}}}`)
	p := gjson.New(nil)
	p.Set("a.b.c", 1)
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set11(t *testing.T) {
	e := []byte(`{"a":{"b":{}}}`)
	p, _ := gjson.LoadContent([]byte(`{"a":{"b":{"c":1}}}`))
	p.Remove("a.b.c")
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set12(t *testing.T) {
	e := []byte(`[0,1]`)
	p := gjson.New(nil)
	p.Set("0", 0)
	p.Set("1", 1)
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set13(t *testing.T) {
	e := []byte(`{"array":[0,1]}`)
	p := gjson.New(nil)
	p.Set("array.0", 0)
	p.Set("array.1", 1)
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set14(t *testing.T) {
	e := []byte(`{"f":{"a":1}}`)
	p := gjson.New(nil)
	p.Set("f", "m")
	p.Set("f.a", 1)
	if c, err := p.ToJson(); err == nil {

		if !bytes.Equal(c, e) {
			t.Error("expect:", string(e))
		}
	} else {
		t.Error(err)
	}
}

func Test_Set15(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)

		t.Assert(j.Set("root.0.k1", "v1"), nil)
		t.Assert(j.Set("root.1.k2", "v2"), nil)
		t.Assert(j.Set("k", "v"), nil)

		s, err := j.ToJsonString()
		t.AssertNil(err)
		t.Assert(
			gstr.Contains(s, `"root":[{"k1":"v1"},{"k2":"v2"}`) ||
				gstr.Contains(s, `"root":[{"k2":"v2"},{"k1":"v1"}`),
			true,
		)
		t.Assert(
			gstr.Contains(s, `{"k":"v"`) ||
				gstr.Contains(s, `"k":"v"}`),
			true,
		)
	})
}

func Test_Set16(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)

		t.Assert(j.Set("processors.0.set.0value", "1"), nil)
		t.Assert(j.Set("processors.0.set.0field", "2"), nil)
		t.Assert(j.Set("description", "3"), nil)

		s, err := j.ToJsonString()
		t.AssertNil(err)
		t.Assert(
			gstr.Contains(s, `"processors":[{"set":{"0field":"2","0value":"1"}}]`) ||
				gstr.Contains(s, `"processors":[{"set":{"0value":"1","0field":"2"}}]`),
			true,
		)
		t.Assert(
			gstr.Contains(s, `{"description":"3"`) || gstr.Contains(s, `"description":"3"}`),
			true,
		)
	})
}

func Test_Set17(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)

		t.Assert(j.Set("0.k1", "v1"), nil)
		t.Assert(j.Set("1.k2", "v2"), nil)
		// overwrite the previous slice.
		t.Assert(j.Set("k", "v"), nil)

		s, err := j.ToJsonString()
		t.AssertNil(err)
		t.Assert(s, `{"k":"v"}`)
	})
}

func Test_Set18(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)

		t.Assert(j.Set("0.1.k1", "v1"), nil)
		t.Assert(j.Set("0.2.k2", "v2"), nil)
		s, err := j.ToJsonString()
		t.AssertNil(err)
		t.Assert(s, `[[null,{"k1":"v1"},{"k2":"v2"}]]`)
	})
}

func Test_Set19(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)

		t.Assert(j.Set("0.1.1.k1", "v1"), nil)
		t.Assert(j.Set("0.2.1.k2", "v2"), nil)
		s, err := j.ToJsonString()
		t.AssertNil(err)
		t.Assert(s, `[[null,[null,{"k1":"v1"}],[null,{"k2":"v2"}]]]`)
	})
}

func Test_Set20(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)

		t.Assert(j.Set("k1", "v1"), nil)
		t.Assert(j.Set("k2", g.Slice{1, 2, 3}), nil)
		t.Assert(j.Set("k2.1", 20), nil)
		t.Assert(j.Set("k2.2", g.Map{"k3": "v3"}), nil)
		s, err := j.ToJsonString()
		t.AssertNil(err)
		t.Assert(gstr.InArray(
			g.SliceStr{
				`{"k1":"v1","k2":[1,20,{"k3":"v3"}]}`,
				`{"k2":[1,20,{"k3":"v3"}],"k1":"v1"}`,
			},
			s,
		), true)
	})
}

func Test_Set_GArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)
		arr := garray.New().Append("test")
		t.AssertNil(j.Set("arr", arr))
		t.Assert(j.Get("arr").Array(), g.Slice{"test"})
	})
}

func Test_Set_WithEmptyStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(&struct{}{})
		t.AssertNil(j.Set("aa", "123"))
		t.Assert(j.MustToJsonString(), `{"aa":"123"}`)
	})
}
