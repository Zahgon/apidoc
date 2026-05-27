// SPDX-License-Identifier: MIT

package lang

import (
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/lexer"
)

// Parse 分析 data 的内容并输出到到 blocks
func Parse(h *core.MessageHandler, langID string, data core.Block, blocks chan core.Block) {
	_ = "STUB: not implemented"
	return
}

type parser struct {
	*lexer.Lexer
	blocks []blocker
	h      *core.MessageHandler
}

func newParser(h *core.MessageHandler, block core.Block, blocks []blocker) *parser {
	_ = "STUB: not implemented"
	return nil
}

// 从当前位置往后查找，直到找到第一个与 blocks 中某个相匹配的，并返回该 Blocker 。
func (l *parser) block() (blocker, core.Position) {
	_ = "STUB: not implemented"
	return *new(blocker), *new(core.Position)
}

// 分析 l.data 的内容并输出到 blocks
func (l *parser) parse(blocks chan core.Block) { _ = "STUB: not implemented"; return }

// 没有匹配的 block 了

// 没有找到结束标签，那肯定是到文件尾了，可以直接返回。

// 重置 block

// end for
