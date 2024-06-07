// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

type ObjectRest struct{}

func (o *ObjectRest) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *ObjectRest) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *ObjectRest) Get(r *ghttp.Request) {
	r.Response.Write("Object Get")
}

func (o *ObjectRest) Put(r *ghttp.Request) {
	r.Response.Write("Object Put")
}

func (o *ObjectRest) Post(r *ghttp.Request) {
	r.Response.Write("Object Post")
}

func (o *ObjectRest) Delete(r *ghttp.Request) {
	r.Response.Write("Object Delete")
}

func (o *ObjectRest) Patch(r *ghttp.Request) {
	r.Response.Write("Object Patch")
}

func (o *ObjectRest) Options(r *ghttp.Request) {
	r.Response.Write("Object Options")
}

func (o *ObjectRest) Head(r *ghttp.Request) {
	r.Response.Header().Set("head-ok", "1")
}

func Test_Router_ObjectRest(t *testing.T) {
	s := g.Server(guid.S())
	s.BindObjectRest("/", new(ObjectRest))
	s.BindObjectRest("/{.struct}/{.method}", new(ObjectRest))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "1Object Get2")
		t.Assert(client.PutContent(ctx, "/"), "1Object Put2")
		t.Assert(client.PostContent(ctx, "/"), "1Object Post2")
		t.Assert(client.DeleteContent(ctx, "/"), "1Object Delete2")
		t.Assert(client.PatchContent(ctx, "/"), "1Object Patch2")
		t.Assert(client.OptionsContent(ctx, "/"), "1Object Options2")
		resp1, err := client.Head(ctx, "/")
		if err == nil {
			defer resp1.Close()
		}
		t.AssertNil(err)
		t.Assert(resp1.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent(ctx, "/object-rest/get"), "1Object Get2")
		t.Assert(client.PutContent(ctx, "/object-rest/put"), "1Object Put2")
		t.Assert(client.PostContent(ctx, "/object-rest/post"), "1Object Post2")
		t.Assert(client.DeleteContent(ctx, "/object-rest/delete"), "1Object Delete2")
		t.Assert(client.PatchContent(ctx, "/object-rest/patch"), "1Object Patch2")
		t.Assert(client.OptionsContent(ctx, "/object-rest/options"), "1Object Options2")
		resp2, err := client.Head(ctx, "/object-rest/head")
		if err == nil {
			defer resp2.Close()
		}
		t.AssertNil(err)
		t.Assert(resp2.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent(ctx, "/none-exist"), "Not Found")
	})
}
