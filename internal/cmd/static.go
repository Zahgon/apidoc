// SPDX-License-Identifier: MIT

package cmd

import (
	"io"

	"github.com/issue9/cmdopt"
)

var (
	staticPort        string
	staticDocs        uri
	staticStylesheet  bool
	staticContentType string
	staticURL         string
	staticPath        uri
)

func initStatic(command *cmdopt.CmdOpt) { _ = "STUB: not implemented"; return }

func static(io.Writer) (err error) { _ = "STUB: not implemented"; return nil }
