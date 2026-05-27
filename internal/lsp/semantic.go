// SPDX-License-Identifier: MIT

package lsp

import (
	"reflect"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
	"github.com/caixw/apidoc/v7/internal/lsp/protocol"
)

type tokenBuilder struct {
	uri              core.URI // 当前 builder 对应的文件地址
	tag, attr, value int      // 可用的 token

	// 每个二级数组长度为 5，表示一组 semanticToken 数据。
	// 数据分别为 绝对行号，当前行的绝对起始位置，长度，以及 token 和 modifier。
	tokens [][]int
}

// textDocument/semanticTokens
func (s *server) textDocumentSemanticTokens(notify bool, in *protocol.SemanticTokensParams, out *protocol.SemanticTokens) error {
	_ = "STUB: not implemented"
	return nil
}

// tag 表示标签名的颜色值；
// attr 表示属性；
// value 表示属性的颜色值；
func semanticTokens(doc *ast.APIDoc, uri core.URI, tag, attr, value int) []int {
	_ = "STUB: not implemented"
	return nil
}

// line 和 start 都为未计算的原始值
func (b *tokenBuilder) append(r core.Range, token int) { _ = "STUB: not implemented"; return }

// 未初始化的段被 node.RealValue 初始化成了零值，其长度必为 0

// 可能存在长度为 0 的，比如 default="" 值的长度为 0

func (b *tokenBuilder) build() []int { _ = "STUB: not implemented"; return nil }

// 同一行，start 取相对值

// sort 排序内容，按从小到大
func (b *tokenBuilder) sort() { _ = "STUB: not implemented"; return }

func (b *tokenBuilder) parse(v reflect.Value) { _ = "STUB: not implemented"; return }

func (b *tokenBuilder) matched(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func (b *tokenBuilder) parseAnonymous(v reflect.Value) { _ = "STUB: not implemented"; return }
