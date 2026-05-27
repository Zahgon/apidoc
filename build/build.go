// SPDX-License-Identifier: MIT

// Package build 提供构建文档的相关功能
package build

import (
	"bytes"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// Build 解析文档并输出文档内容
//
// 如果是配置文件有问题，则直接返回错误信息，文档错误则输出至 h 对象。
func Build(h *core.MessageHandler, o *Output, i ...*Input) error {
	_ = "STUB: not implemented"
	return nil
}

// Buffer 生成文档内容并返回
//
// 如果是配置文件有问题，则直接返回错误信息，文档错误则输出至 h 对象。
func Buffer(h *core.MessageHandler, o *Output, i ...*Input) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckSyntax 测试文档语法
//
// 如果是配置文件有问题，则直接返回错误信息，文档错误则输出至 h 对象。
func CheckSyntax(h *core.MessageHandler, i ...*Input) error { _ = "STUB: not implemented"; return nil }

func parse(h *core.MessageHandler, i ...*Input) (*ast.APIDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
