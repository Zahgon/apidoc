// SPDX-License-Identifier: MIT

package lsp

import (
	"context"
	"log"
	"sync"

	"github.com/issue9/jsonrpc"
	"golang.org/x/text/message"

	"github.com/caixw/apidoc/v7/internal/lsp/protocol"
)

type serverState int

const (
	serverCreated serverState = iota
	serverInitializing
	serverInitialized
	serverShutdown
)

// server LSP 服务实例
type server struct {
	*jsonrpc.Conn
	state        serverState
	trace        string
	stateMux     sync.RWMutex
	workspaceMux sync.RWMutex

	folders []*folder

	clientParams *protocol.InitializeParams
	serverResult *protocol.InitializeResult
	info, erro   *log.Logger
	cancelFunc   context.CancelFunc
}

func newServe(t jsonrpc.Transport, infolog, errlog *log.Logger) *server {
	_ = "STUB: not implemented"
	return nil
}

// workspace

// textDocument

// apidoc 自定义的接口

func (s *server) serve() error { _ = "STUB: not implemented"; return nil }

func (s *server) setState(state serverState) { _ = "STUB: not implemented"; return }

func (s *server) getState() serverState { _ = "STUB: not implemented"; return *new(serverState) }

// $/setTrace
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#setTrace
func (s *server) setTrace(notify bool, in *protocol.SetTraceParams, out *any) error {
	_ = "STUB: not implemented"
	return nil
}

// $/logTrace
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#logTrace
func (s *server) logTrace(message, verbose string) { _ = "STUB: not implemented"; return }

// $/cancelRequest
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#cancelRequest
func (s *server) cancel(notify bool, in *protocol.CancelParams, out *any) error {
	_ = "STUB: not implemented"

	// 所有以 $/ 开头且未处理的服务由此函数处理
	//
	// $ Notifications and Requests
	//
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#dollarRequests
	return nil
}

func (s *server) dollarHandler(notify bool, in, out *any) error {
	_ = "STUB: not implemented"
	return nil
}

// window/logMessage
//
// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#window_logMessage
func (s *server) windowLogMessage(t protocol.MessageType, message string) {
	_ = "STUB: not implemented"
	return
}

func (s *server) printErr(err error) { _ = "STUB: not implemented"; return }

func (s *server) windowLogInfoMessage(key message.Reference, v ...any) {
	_ = "STUB: not implemented"
	return
}

func (s *server) windowLogLogMessage(key message.Reference, v ...any) {
	_ = "STUB: not implemented"
	return
}

func (s *server) windowLogWarnMessage(key message.Reference, v ...any) {
	_ = "STUB: not implemented"
	return
}

func (s *server) windowLogErrorMessage(key message.Reference, v ...any) {
	_ = "STUB: not implemented"
	return
}
