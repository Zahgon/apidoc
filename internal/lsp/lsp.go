// SPDX-License-Identifier: MIT

// Package lsp 提供 language server protocol 服务
package lsp

import (
	"log"
	"time"

	"github.com/issue9/jsonrpc"
)

// Version lsp 的版本
const Version = "3.16.0"

// Serve 执行 LSP 服务
//
// t 表示服务的类型，可以是 stdio、udp、tcp 和 unix。
func Serve(header bool, t string, addr string, timeout time.Duration, infolog, errlog *log.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func serveStdio(header bool, infolog, errlog *log.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func serveUDP(header bool, addr string, timeout time.Duration, infolog, errlog *log.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// t 可以是 tcp 和 unix
func serveTCP(header bool, t string, addr string, timeout time.Duration, infolog, errlog *log.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func serve(t jsonrpc.Transport, infolog, errlog *log.Logger) error {
	_ = "STUB: not implemented"
	return nil
}
