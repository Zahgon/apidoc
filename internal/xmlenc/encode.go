// SPDX-License-Identifier: MIT

package xmlenc

import (
	"encoding/xml"
	"reflect"

	"github.com/caixw/apidoc/v7/internal/node"
)

// Encoder 将元素内容编码成 XML 内容
type Encoder interface {
	// EncodeXML 仅需要返回元素内容的 XML 编码，不需要包含本身的标签和属性。
	EncodeXML() (string, error)
}

// AttrEncoder 将属性值编码成符合 XML 规范的值
type AttrEncoder interface {
	// EncodeXMLAttr 仅需要返回属性的 XML 表示，不需要包含属性值的引号字符。
	EncodeXMLAttr() (string, error)
}

var (
	attrEncoderType = reflect.TypeOf((*AttrEncoder)(nil)).Elem()
	encoderType     = reflect.TypeOf((*Encoder)(nil)).Elem()
)

// Encode 将 v 转换成 XML 内容
//
// namespace 指定 XML 的命名空间；
// prefix 命名空间的前缀，如果为空表示使用默认命名空间；
func Encode(indent string, v any, namespace, prefix string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encode(n *node.Node, e *xml.Encoder, namespace, prefix string, root bool) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeElements(n *node.Node, e *xml.Encoder, start xml.StartElement, namespace, prefix string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func encodeElement(e *xml.Encoder, v *node.Value, namespace, prefix string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func buildStartElement(n *node.Node, namespace, prefix string, root bool) (xml.StartElement, error) {
	_ = "STUB: not implemented"
	return *new(xml.StartElement), nil
}

// +1 表示 xmlns 的字段

func buildXMLName(name, prefix string) xml.Name { _ = "STUB: not implemented"; return *new(xml.Name) }

func getAttributeValue(elem reflect.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// 获取 cdata 和 content 节点的的内容
func getContentValue(elem reflect.Value) (string, error) { _ = "STUB: not implemented"; return "", nil }

func isOmitempty(v *node.Value) bool { _ = "STUB: not implemented"; return false }
