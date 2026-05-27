// SPDX-License-Identifier: MIT

package ast

import (
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/xmlenc"
)

// ParseBlocks 从多个 core.Block 实例中解析文档内容
//
// g 必须是一个阻塞函数，直到所有代码块都写入参数之后，才能返回。
func (doc *APIDoc) ParseBlocks(h *core.MessageHandler, g func(chan core.Block)) {
	_ = "STUB: not implemented"
	return
}

// Parse 将注释块的内容添加到当前文档
func (doc *APIDoc) Parse(h *core.MessageHandler, b core.Block) { _ = "STUB: not implemented"; return }

// apidoc 已经初始化，检测依赖于 apidoc 的字段

// 多个 apidoc 标签

// api 进入 doc 的顺序是未知的，进行排序可以保证文档的顺序一致。

// 简单预判是否是一个合规的 apidoc 内容
func isValid(b core.Block) bool { _ = "STUB: not implemented"; return false }

// 去除空格之后，必须保证以 < 开头，且不能以 </ 开关。

// 获取根标签的名称
func getTagName(p *xmlenc.Parser) string { _ = "STUB: not implemented"; return "" }

// 获取第一个元素名称就出错，说明不是一个合则的 XML，直接忽略。

// 表示是一个非法的 XML，忽略！

// 其它标签忽略

func (doc *APIDoc) sortAPIs() { _ = "STUB: not implemented"; return }
