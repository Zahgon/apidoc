// SPDX-License-Identifier: MIT

package lsp

import (
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
	"github.com/caixw/apidoc/v7/internal/lsp/protocol"
)

// textDocument/didChange
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_didChange
func (s *server) textDocumentDidChange(notify bool, in *protocol.DidChangeTextDocumentParams, out *any) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *folder) parseBlock(block core.Block) { _ = "STUB: not implemented"; return }

// 无需解析

func deleteURI(doc *ast.APIDoc, uri core.URI) (deleted bool) {
	_ = "STUB: not implemented"
	return false
}

// textDocument/publishDiagnostics
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_publishDiagnostics
func (s *server) textDocumentPublishDiagnostics(f *folder) { _ = "STUB: not implemented"; return }

// 清空所有的诊断信息
func (f *folder) clearDiagnostics() { _ = "STUB: not implemented"; return }

// textDocument/foldingRange
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_foldingRange
func (s *server) textDocumentFoldingRange(notify bool, in *protocol.FoldingRangeParams, out *[]protocol.FoldingRange) error {
	_ = "STUB: not implemented"
	return nil
}

// textDocument/completion
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_completion
func (s *server) textDocumentCompletion(notify bool, in *protocol.CompletionParams, out *protocol.CompletionList) error {
	_ = "STUB: not implemented"
	// TODO
	return nil
}
