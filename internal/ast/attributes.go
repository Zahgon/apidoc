// SPDX-License-Identifier: MIT

package ast

import (
	"net/http"
	"time"

	"github.com/caixw/apidoc/v7/internal/xmlenc"
)

const dateFormat = time.RFC3339

type (
	// Attribute 表示 XML 属性
	Attribute struct {
		xmlenc.BaseAttribute
		Value    xmlenc.String `apidoc:"-"`
		RootName struct{}      `apidoc:"string,meta,usage-string"`
	}

	// NumberAttribute 表示数值类型的属性
	NumberAttribute struct {
		xmlenc.BaseAttribute
		Value    Number   `apidoc:"-"`
		RootName struct{} `apidoc:"number,meta,usage-number"`
	}

	// BoolAttribute 表示布尔值类型的属性
	BoolAttribute struct {
		xmlenc.BaseAttribute
		Value    Bool     `apidoc:"-"`
		RootName struct{} `apidoc:"bool,meta,usage-bool"`
	}

	// VersionAttribute 表示版本号属性
	VersionAttribute struct {
		xmlenc.BaseAttribute
		Value    xmlenc.String `apidoc:"-"`
		RootName struct{}      `apidoc:"version,meta,usage-version"`
	}

	// DateAttribute 日期属性
	DateAttribute struct {
		xmlenc.BaseAttribute
		Value    Date     `apidoc:"-"`
		RootName struct{} `apidoc:"date,meta,usage-date"`
	}

	// MethodAttribute 表示请求方法
	MethodAttribute Attribute

	// StatusAttribute 状态码的 XML 属性
	StatusAttribute NumberAttribute

	// TypeAttribute 表示方法类型属性
	TypeAttribute struct {
		xmlenc.BaseAttribute
		Value    xmlenc.String `apidoc:"-"`
		RootName struct{}      `apidoc:"type,meta,usage-type"`
	}

	// APIDocVersionAttribute 版本号属性，同时对版本号进行比较
	APIDocVersionAttribute Attribute
)

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (a *Attribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (a *Attribute) EncodeXMLAttr() (string, error) {
	_ = "STUB: not implemented"

	// V 返回当前属性实际表示的值
	return "", nil
}

func (a *Attribute) V() string { _ = "STUB: not implemented"; return "" }

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (num *NumberAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (num *NumberAttribute) EncodeXMLAttr() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IntValue 返回当前属性实际表示的值
func (num *NumberAttribute) IntValue() int { _ = "STUB: not implemented"; return 0 }

// FloatValue 返回当前属性实际表示的值
func (num *NumberAttribute) FloatValue() float64 { _ = "STUB: not implemented"; return 0 }

// IsFloat 当前的数值类型是否为浮点型
func (num *NumberAttribute) IsFloat() bool { _ = "STUB: not implemented"; return false }

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (b *BoolAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (b *BoolAttribute) EncodeXMLAttr() (string, error) { _ = "STUB: not implemented"; return "", nil }

// V 返回当前属性实际表示的值
func (b *BoolAttribute) V() bool { _ = "STUB: not implemented"; return false }

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (a *MethodAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (a *MethodAttribute) EncodeXMLAttr() (string, error) {
	_ = "STUB: not implemented"

	// V 返回当前属性实际表示的值
	return "", nil
}

func (a *MethodAttribute) V() string { _ = "STUB: not implemented"; return "" }

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (a *StatusAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (a *StatusAttribute) EncodeXMLAttr() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// V 返回当前属性实际表示的值
func (a *StatusAttribute) V() int { _ = "STUB: not implemented"; return 0 }

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (a *TypeAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (a *TypeAttribute) EncodeXMLAttr() (string, error) {
	_ = "STUB: not implemented"

	// V 返回当前属性实际表示的值
	return "", nil
}

func (a *TypeAttribute) V() string { _ = "STUB: not implemented"; return "" }

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (a *VersionAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (a *VersionAttribute) EncodeXMLAttr() (string, error) {
	_ = "STUB: not implemented"

	// V 返回当前属性实际表示的值
	return "", nil
}

func (a *VersionAttribute) V() string { _ = "STUB: not implemented"; return "" }

// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
func (d *DateAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (d *DateAttribute) EncodeXMLAttr() (string, error) { _ = "STUB: not implemented"; return "", nil }

// V 返回当前属性实际表示的值
func (d *DateAttribute) V() time.Time {
	_ = "STUB: not implemented"
	return *

	// DecodeXMLAttr AttrDecoder.DecodeXMLAttr
	new(time.Time)
}

func (a *APIDocVersionAttribute) DecodeXMLAttr(p *xmlenc.Parser, attr *xmlenc.Attribute) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeXMLAttr AttrEncoder.EncodeXMLAttr
func (a *APIDocVersionAttribute) EncodeXMLAttr() (string, error) {
	_ = "STUB: not implemented"

	// V 返回当前属性实际表示的值
	return "", nil
}

func (a *APIDocVersionAttribute) V() string { _ = "STUB: not implemented"; return "" }

var validMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodHead,
	http.MethodOptions,
}

func isValidMethod(method string) bool { _ = "STUB: not implemented"; return false }

func isValidStatus(status int) bool { _ = "STUB: not implemented"; return false }

func isValidType(t string) bool { _ = "STUB: not implemented"; return false }

func isValidVersion(v string) bool { _ = "STUB: not implemented"; return false }
