// SPDX-License-Identifier: MIT

package lsp

import (
	"github.com/caixw/apidoc/v7/internal/lsp/protocol"
)

// initialize
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#initialize
func (s *server) initialize(notify bool, in *protocol.InitializeParams, out *protocol.InitializeResult) error {
	_ = "STUB: not implemented"
	return nil
}

// 输出错误信息，但是不中断执行

// initialized
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#initialized
func (s *server) initialized(bool, *protocol.InitializedParams, *any) error {
	_ = "STUB: not implemented"
	return nil
}

// shutdown
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#shutdown
func (s *server) shutdown(bool, *any, *any) error { _ = "STUB: not implemented"; return nil }

// exit
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#exit
func (s *server) exit(bool, *any, *any) error { _ = "STUB: not implemented"; return nil }
