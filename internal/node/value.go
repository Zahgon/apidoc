// SPDX-License-Identifier: MIT

package node

import (
	"reflect"
)

// Value 表示 XML 节点的值的反射表示方式
type Value struct {
	reflect.Value
	Omitempty bool
	Name      string // 节点的名称

	// 当前值可能未初始化，所以保存 usage 的值，
	// 等 Value 初始化之后再赋值给 Base.UsageKey
	Usage string
}

// NewValue 声明 *Value 实例
func NewValue(name string, v reflect.Value, omitempty bool, usage string) *Value {
	_ = "STUB: not implemented"
	return nil
}

// ParseValue 分析 v 并返回 *Value 实例
//
// 与 NewValue 的不同在于，ParseValue 会分析对象字段中是否带有 meta 的结构体标签，
// 如果有才初始化 *Value 对象，否则返回 nil。
func ParseValue(v reflect.Value) *Value { _ = "STUB: not implemented"; return nil }

// IsPrimitive 是否为有效的 Go 原始类型
func IsPrimitive(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// RealType 获取指针指向的类型
func RealType(t reflect.Type) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// RealValue 获取指针指向的值
//
// 如果未初始化，则会对其进行初始化。
func RealValue(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
