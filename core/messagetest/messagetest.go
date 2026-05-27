// SPDX-License-Identifier: MIT

// Package messagetest 提供测试生成 message 相关的测试工具
package messagetest

import "github.com/caixw/apidoc/v7/core"

type Result struct {
	Errors, Warns, Infos, Successes []any
	Handler                         *core.MessageHandler
}

// NewMessageHandler 返回一个用于测试的 core.MessageHandler 实例
func NewMessageHandler() *Result { _ = "STUB: not implemented"; return nil }
