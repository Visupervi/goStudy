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

func Test_Router_Group_Group(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			r.Response.Write("1")
			r.Middleware.Next()
			r.Response.Write("2")
		})
		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
		group.Group("/order", func(group *ghttp.RouterGroup) {
			group.GET("/list", func(r *ghttp.Request) {
				r.Response.Write("list")
			})
			group.PUT("/update", func(r *ghttp.Request) {
				r.Response.Write("update")
			})
		})
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.GET("/info", func(r *ghttp.Request) {
				r.Response.Write("info")
			})
			group.POST("/edit", func(r *ghttp.Request) {
				r.Response.Write("edit")
			})
			group.DELETE("/drop", func(r *ghttp.Request) {
				r.Response.Write("drop")
			})
		})
		group.Group("/hook", func(group *ghttp.RouterGroup) {
			group.Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
				r.Response.Write("hook any")
			})
			group.Hook("/:name", ghttp.HookBeforeServe, func(r *ghttp.Request) {
				r.Response.Write("hook name")
			})
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "Not Found")
		t.Assert(client.GetContent(ctx, "/api.v2"), "Not Found")
		t.Assert(client.GetContent(ctx, "/api.v2/test"), "1test2")
		t.Assert(client.GetContent(ctx, "/api.v2/hook"), "hook any")
		t.Assert(client.GetContent(ctx, "/api.v2/hook/name"), "hook namehook any")
		t.Assert(client.GetContent(ctx, "/api.v2/hook/name/any"), "hook any")
		t.Assert(client.GetContent(ctx, "/api.v2/order/list"), "1list2")
		t.Assert(client.GetContent(ctx, "/api.v2/order/update"), "Not Found")
		t.Assert(client.PutContent(ctx, "/api.v2/order/update"), "1update2")
		t.Assert(client.GetContent(ctx, "/api.v2/user/drop"), "Not Found")
		t.Assert(client.DeleteContent(ctx, "/api.v2/user/drop"), "1drop2")
		t.Assert(client.GetContent(ctx, "/api.v2/user/edit"), "Not Found")
		t.Assert(client.PostContent(ctx, "/api.v2/user/edit"), "1edit2")
		t.Assert(client.GetContent(ctx, "/api.v2/user/info"), "1info2")
	})
}
