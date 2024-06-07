// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gins_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/gins"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	ctx           = context.Background()
	configContent = gfile.GetContents(
		gtest.DataPath("config", "config.toml"),
	)
)

func Test_Config1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(configContent, "")
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gins.Config(), nil)
	})
}

func Test_Config2(t *testing.T) {
	// relative path
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(dirPath)
		t.AssertNil(err)
		defer gfile.Remove(dirPath)

		name := "config.toml"
		err = gfile.PutContents(gfile.Join(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config().GetAdapter().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().GetAdapter().(*gcfg.AdapterFile).Clear()

		t.Assert(gins.Config().MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config().MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	// for gfsnotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	// relative path, config folder
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.Temp(gtime.TimestampNanoStr())
		t.AssertNil(gfile.Mkdir(dirPath))
		defer gfile.Remove(dirPath)

		name := "config/config.toml"
		err = gfile.PutContents(gfile.Join(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config().GetAdapter().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config().GetAdapter().(*gcfg.AdapterFile).Clear()

		t.Assert(gins.Config().MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config().MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)

		// for gfsnotify callbacks to refresh cache of config file
		time.Sleep(500 * time.Millisecond)
	})
}

func Test_Config3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(dirPath)
		t.AssertNil(err)
		defer gfile.Remove(dirPath)

		name := "test.toml"
		err = gfile.PutContents(gfile.Join(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config("test").GetAdapter().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config("test").GetAdapter().(*gcfg.AdapterFile).Clear()
		gins.Config("test").GetAdapter().(*gcfg.AdapterFile).SetFileName("test.toml")

		t.Assert(gins.Config("test").MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	// for gfsnotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		var err error
		dirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(dirPath)
		t.AssertNil(err)
		defer gfile.Remove(dirPath)

		name := "config/test.toml"
		err = gfile.PutContents(gfile.Join(dirPath, name), configContent)
		t.AssertNil(err)

		err = gins.Config("test").GetAdapter().(*gcfg.AdapterFile).AddPath(dirPath)
		t.AssertNil(err)

		defer gins.Config("test").GetAdapter().(*gcfg.AdapterFile).Clear()
		gins.Config("test").GetAdapter().(*gcfg.AdapterFile).SetFileName("test.toml")

		t.Assert(gins.Config("test").MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	// for gfsnotify callbacks to refresh cache of config file for next unit testing case.
	time.Sleep(500 * time.Millisecond)
}

func Test_Config4(t *testing.T) {
	// absolute path
	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d`, gfile.Temp(), gtime.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := gfile.PutContents(file, configContent)
		t.AssertNil(err)
		defer gfile.Remove(file)
		defer gins.Config().GetAdapter().(*gcfg.AdapterFile).Clear()

		t.Assert(gins.Config().GetAdapter().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config().MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config().MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d/config`, gfile.Temp(), gtime.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := gfile.PutContents(file, configContent)
		t.AssertNil(err)
		defer gfile.Remove(file)
		defer gins.Config().GetAdapter().(*gcfg.AdapterFile).Clear()
		t.Assert(gins.Config().GetAdapter().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config().MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config().MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config().MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d`, gfile.Temp(), gtime.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := gfile.PutContents(file, configContent)
		t.AssertNil(err)
		defer gfile.Remove(file)
		defer gins.Config("test").GetAdapter().(*gcfg.AdapterFile).Clear()
		gins.Config("test").GetAdapter().(*gcfg.AdapterFile).SetFileName("test.toml")
		t.Assert(gins.Config("test").GetAdapter().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config("test").MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
	time.Sleep(500 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d/config`, gfile.Temp(), gtime.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := gfile.PutContents(file, configContent)
		t.AssertNil(err)
		defer gfile.Remove(file)
		defer gins.Config().GetAdapter().(*gcfg.AdapterFile).Clear()
		gins.Config("test").GetAdapter().(*gcfg.AdapterFile).SetFileName("test.toml")
		t.Assert(gins.Config("test").GetAdapter().(*gcfg.AdapterFile).AddPath(path), nil)
		t.Assert(gins.Config("test").MustGet(ctx, "test"), "v=1")
		t.Assert(gins.Config("test").MustGet(ctx, "database.default.1.host"), "127.0.0.1")
		t.Assert(gins.Config("test").MustGet(ctx, "redis.disk"), `{"address":"127.0.0.1:6379","db":1}`)
	})
}

func Test_Basic2(t *testing.T) {
	config := `log-path = "logs"`
	gtest.C(t, func(t *gtest.T) {
		var (
			path = gcfg.DefaultConfigFileName
			err  = gfile.PutContents(path, config)
		)
		t.AssertNil(err)
		defer func() {
			_ = gfile.Remove(path)
		}()

		t.Assert(gins.Config().MustGet(ctx, "log-path"), "logs")
	})
}
