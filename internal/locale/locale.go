// SPDX-License-Identifier: MIT

// Package locale 提供了一个本地化翻译服务。
package locale

import (
	"github.com/issue9/localeutil"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// DefaultLocaleID 默认的本地化语言 ID
//
// 当未调用相关函数设置 ID，或是设置为一个不支持的 ID 时，
// 系统最终会采用此 ID。
const DefaultLocaleID = "cmn-Hans"

var (
	// 保证有个初始化的值，部分包的测试功能依赖此变量
	localeTag     = language.MustParse(DefaultLocaleID)
	localePrinter = message.NewPrinter(localeTag)

	tags = []language.Tag{}
)

type LocaleStringer = localeutil.LocaleStringer

type Error struct {
	l LocaleStringer
}

func setMessages(id string, messages map[string]string) { _ = "STUB: not implemented"; return }

// 保证 DefaultLocaleID 为第一个数组元素

func (err *Error) Error() string { _ = "STUB: not implemented"; return "" }

// SetTag 切换本地化环境
func SetTag(tag language.Tag) { _ = "STUB: not implemented"; return }

// Tag 获取当前的本地化 ID
func Tag() language.Tag {
	_ = "STUB: not implemented"

	// Tags 所有支持语言的列表
	return *new(language.Tag)
}

func Tags() []language.Tag { _ = "STUB: not implemented"; return nil }

// Sprintf 类似 fmt.Sprintf，与特定的本地化绑定。
func Sprintf(key message.Reference, v ...any) string { _ = "STUB: not implemented"; return "" }

// New 声明新的 Locale 对象
func New(key message.Reference, v ...any) LocaleStringer {
	_ = "STUB: not implemented"
	return *new(LocaleStringer)
}

// NewError 返回本地化的错误对象
func NewError(key message.Reference, v ...any) error { _ = "STUB: not implemented"; return nil }

// Translate 功能与 Sprintf 类似，但是可以指定本地化 ID 值。
func Translate(localeID string, key message.Reference, v ...any) string {
	_ = "STUB: not implemented"
	return ""
}
