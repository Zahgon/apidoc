// SPDX-License-Identifier: MIT

package core

import (
	"golang.org/x/text/message"
)

// HTTPError 表示 HTTP 状态码的错误
type HTTPError struct {
	error
	Code int
}

// Error 用于描述 apidoc 中的大部分错误信息
//
// 无论是配置文件的错误，还是文档的语法错误，都将返回此错误。
type Error struct {
	Err      error                // 具体的错误信息
	Location Location             // 错误的详细定位
	Field    string               // 出错的字段
	Types    []ErrorType          // 该错误的类型
	Related  []RelatedInformation // 与此错误关联的一些信息
}

// RelatedInformation 错误信息的关联内容
//
// 兼容 LSP 中的 DiagnosticRelatedInformation 的相关定义
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16/#diagnostic
type RelatedInformation struct {
	Location Location `json:"location"`
	Message  string   `json:"message"`
}

// ErrorType 语法错误的类型
type ErrorType int

// 语法错误类型可用的枚举值
const (
	ErrorTypeDeprecated ErrorType = iota + 1
	ErrorTypeUnused
)

// NewHTTPError 声明 HTTPError 实例
func NewHTTPError(code int, key message.Reference, v ...any) *HTTPError {
	_ = "STUB: not implemented"
	return nil
}

func (err *Error) Error() string { _ = "STUB: not implemented"; return "" }

// ErrMessage = "%s 位次于 %s:%d"

// Unwrap 实现 errors.Unwrap 接口
func (err *Error) Unwrap() error {
	_ = "STUB: not implemented"

	// Is 实现 errors.Is 接口
	return nil
}

func (err *Error) Is(target error) bool { _ = "STUB: not implemented"; return false }

// Relate 添加关联的错误信息
func (err *Error) Relate(loc Location, msg string) *Error { _ = "STUB: not implemented"; return nil }

// WithField 为语法错误修改或添加具体的错误字段
func (err *Error) WithField(field string) *Error { _ = "STUB: not implemented"; return nil }

// WithLocation  为语法错误添加定位信息
func (err *Error) WithLocation(loc Location) *Error { _ = "STUB: not implemented"; return nil }

// AddTypes 为语法错误添加错误类型
func (err *Error) AddTypes(t ...ErrorType) *Error { _ = "STUB: not implemented"; return nil }

// NewError 返回 *Error 实例
func NewError(key message.Reference, v ...any) *Error { _ = "STUB: not implemented"; return nil }

// WithError 采用 err 实例 *Error 实例
func WithError(err error) *Error { _ = "STUB: not implemented"; return nil }

// NewError 在当前位置生成语法错误信息
//
// 其中的 msg 和 val 会被转换成本地化的内容保存。
func (l Location) NewError(key message.Reference, v ...any) *Error {
	_ = "STUB: not implemented"
	return nil
}

// WithError 在当前位置生成语法错误信息
//
// 若 err 本身就是 *Error 类型，则会更新其 location 和 Field 两个字段的信息。
func (l Location) WithError(err error) *Error { _ = "STUB: not implemented"; return nil }
