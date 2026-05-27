// SPDX-License-Identifier: MIT

package mock

import (
	"encoding/xml"

	"github.com/caixw/apidoc/v7/internal/ast"
)

type xmlValidator struct {
	namespaces []*ast.XMLNamespace
	decoder    *xml.Decoder
}

func validXML(ns []*ast.XMLNamespace, p *ast.Request, content []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// 正常结束

func (v *xmlValidator) validXMLNamespaces(start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

// 默认

func (v *xmlValidator) validXMLName(name xml.Name, p *ast.Param, chkArray bool) bool {
	_ = "STUB: not implemented"
	return false
}

// else goto SPACE

// parent 如果是数组，则是否拿 wrapped 中指示的父元素名称。
func parseXMLWrappedName(p *ast.Param, parent bool) (name string) {
	_ = "STUB: not implemented"
	return ""
}

// index > 0

func (v *xmlValidator) validXMLElement(start xml.StartElement, p *ast.Param, chkArray bool, field string) error {
	_ = "STUB: not implemented"
	return nil
}

// 正常结束

func (v *xmlValidator) validStartElement(start xml.StartElement, p *ast.Param, chkArray bool, field string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildXMLField(field string, p *ast.Param) string { _ = "STUB: not implemented"; return "" }

// 验证 p 描述的类型与 v 是否匹配，如果不匹配返回错误信息。
// field 表示 p 在整个对象中的位置信息。
func validXMLValue(p *ast.Param, field, v string) error { _ = "STUB: not implemented"; return nil }

type xmlBuilder struct {
	start xml.StartElement
	items []*xmlBuilder

	chardata any
	cdata    bool // 表示 chardata 是一个 cdata 数据
}

func buildXML(ns []*ast.XMLNamespace, p *ast.Request, indent string, g *GenOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseXML(ns []*ast.XMLNamespace, p *ast.Param, chkArray, root bool, g *GenOptions) (*xmlBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// end for

// namespace

func buildXMLName(p *ast.Param, chkArray bool) xml.Name {
	_ = "STUB: not implemented"
	return *new(xml.Name)
}

func parseXMLArray(ns []*ast.XMLNamespace, p *ast.Param, parent *xmlBuilder, g *GenOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (builder *xmlBuilder) encode(e *xml.Encoder) error { _ = "STUB: not implemented"; return nil }

func genXMLValue(g *GenOptions, p *ast.Param) any { _ = "STUB: not implemented"; return *new(any) }

// ast.TypeObject:
// 加载的时候已经作语法验证，此处还出错则直接 panic
