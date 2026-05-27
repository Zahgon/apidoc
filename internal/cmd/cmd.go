// SPDX-License-Identifier: MIT

// Package cmd 提供子命令的相关功能
package cmd

import (
	"io"
	"os"

	"github.com/issue9/cmdopt"
	"github.com/issue9/term/v3/colors"
	"golang.org/x/text/message"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/locale"
)

// 命令行输出的表格中，每一列为了对齐填补的空格数量。
const tail = 3

var printers = map[core.MessageType]*printer{
	core.Erro: {
		out:    os.Stderr,
		color:  colors.Red,
		prefix: locale.ErrorPrefix,
	},
	core.Warn: {
		out:    os.Stderr,
		color:  colors.Cyan,
		prefix: locale.WarnPrefix,
	},
	core.Info: {
		out:    os.Stdout,
		color:  colors.Default,
		prefix: locale.InfoPrefix,
	},
	core.Succ: {
		out:    os.Stdout,
		color:  colors.Green,
		prefix: locale.SuccessPrefix,
	},
}

type printer struct {
	out    io.Writer
	color  colors.Color
	prefix message.Reference
}

type uri core.URI

func (u uri) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (u *uri) Set(v string) error { _ = "STUB: not implemented"; return nil }

func (u *uri) String() string { _ = "STUB: not implemented"; return "" }

func (u uri) URI() core.URI {
	_ = "STUB: not implemented"

	// Init 初始化 cmdopt.CmdOpt 实例
	return *new(core.URI)
}

func Init(out io.Writer) *cmdopt.CmdOpt { _ = "STUB: not implemented"; return nil }

func messageHandle(msg *core.Message) { _ = "STUB: not implemented"; return }

func (p *printer) print(msg any) { _ = "STUB: not implemented"; return }
