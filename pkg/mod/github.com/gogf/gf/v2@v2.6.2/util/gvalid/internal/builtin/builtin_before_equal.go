// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// RuleBeforeEqual implements `before-equal` rule:
// The datetime value should be after or equal to the value of field `field`.
//
// Format: before-equal:field
type RuleBeforeEqual struct{}

func init() {
	Register(RuleBeforeEqual{})
}

func (r RuleBeforeEqual) Name() string {
	return "before-equal"
}

func (r RuleBeforeEqual) Message() string {
	return "The {field} value `{value}` must be before or equal to field {pattern}"
}

func (r RuleBeforeEqual) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.Before(fieldDatetime) || valueDatetime.Equal(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
