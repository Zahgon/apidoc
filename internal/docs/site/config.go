// SPDX-License-Identifier: MIT

package site

import (
	"reflect"
)

func (d *doc) newConfig() error { _ = "STUB: not implemented"; return nil }

func (d *doc) buildConfigItem(parent string, f reflect.StructField) error {
	_ = "STUB: not implemented"
	return nil
}

// 调用方需要保证 t.Kind() 为 reflect.Struct
func (d *doc) buildConfigObject(parent string, t reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func isPrimitive(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func parseTag(f reflect.StructField) (string, bool) { _ = "STUB: not implemented"; return "", false }

func getName(n string, f reflect.StructField) string { _ = "STUB: not implemented"; return "" }
