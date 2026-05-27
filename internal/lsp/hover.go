// SPDX-License-Identifier: MIT

package lsp

import (
	"reflect"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/lsp/protocol"
)

type usager interface {
	core.Searcher
	Usage() string
}

var usagerType = reflect.TypeOf((*usager)(nil)).Elem()

// textDocument/hover
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#textDocument_hover
func (s *server) textDocumentHover(notify bool, in *protocol.HoverParams, out *protocol.Hover) error {
	_ = "STUB: not implemented"
	return nil
}

// 非项目文件，不应该出错
