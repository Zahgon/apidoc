// SPDX-License-Identifier: MIT

// Package node 处理 ast 中各个节点的结构信息
//
// struct tag
//
// 标签属性分为 4 个字段，其中前三个是必填的：
//
//	apidoc:"name,node-type,usage-key,omitempty"
//
// name 表示当前标签的名称，或是节点表示的类型；
// node-type 表示当前节点的类型，可以是以下值：
//   - elem 表示这是一个子元素；
//   - attr 表示为一个 XML 属性；
//   - cdata 表示为 CDATA 数据；
//   - content 表示为普通的字符串值；
//   - meta 表示这个字段仅用于描述当前元素的元数据，比如元素的名称等；
//
// usage-key 指定了当前元素的翻译项；
// omitempty 表示当前值为空时，是否可以忽略。
package node

import (
	"reflect"
)

// TagName 结构体标签的名称
const TagName = "apidoc"

// Type 节点类型
type Type int8

// 节点类型的值
const (
	attribute Type = iota
	element
	cdata
	content
	meta // 用于描述节点的一些元数据
)

var stringNodeMap = map[string]Type{
	"attr":    attribute,
	"elem":    element,
	"cdata":   cdata,
	"content": content,
	"meta":    meta,
}

// Node 表示一个 XML 标签节点
type Node struct {
	Attributes     []*Value // 当前标签的属性值列表
	Elements       []*Value // 当前标签的元素列表
	CData, Content *Value   // 当前标签如果没有子元素，则可能有普通的内容或是 CDATA 内容
	Value          Value    // 当前节点本身代表的值
	TypeName       string   // 当前节点的类型名称
}

// New 声明 Node 实例
func New(name string, rv reflect.Value) *Node { _ = "STUB: not implemented"; return nil }

// 顶层元素可能没有 name，此处就和 fieldName 相同

func (n *Node) appendAnonymous(v reflect.Value) { _ = "STUB: not implemented"; return }

func (n *Node) setCData(v *Value) { _ = "STUB: not implemented"; return }

func (n *Node) setContent(v *Value) { _ = "STUB: not implemented"; return }

// Element 查找名称为 name 的节点元素
func (n *Node) Element(name string) (*Value, bool) { _ = "STUB: not implemented"; return nil, false }

// Attribute 查找名称为 name 的节点属性
func (n *Node) Attribute(name string) (*Value, bool) { _ = "STUB: not implemented"; return nil, false }

func (n *Node) findElem(name string, elems []*Value) (*Value, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (n *Node) appendAttr(v *Value) { _ = "STUB: not implemented"; return }

func (n *Node) appendElem(v *Value) { _ = "STUB: not implemented"; return }

// `apidoc:"name,attr,usage,omitempty"`
func parseTag(field reflect.StructField) (string, Type, string, bool) {
	_ = "STUB: not implemented"
	return "", *new(Type), "", false
}

func getTagName(field reflect.StructField, name string) string {
	_ = "STUB: not implemented"
	return ""
}

func getNodeType(v string) Type { _ = "STUB: not implemented"; return *new(Type) }

func getOmitempty(v string) bool { _ = "STUB: not implemented"; return false }
