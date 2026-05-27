// SPDX-License-Identifier: MIT

package cmd

import (
	"io"
	"time"

	"github.com/issue9/cmdopt"
)

var (
	lspPort    string
	lspMode    string
	lspHeader  bool
	lspTimeout time.Duration
)

func initLSP(command *cmdopt.CmdOpt) { _ = "STUB: not implemented"; return }

func doLSP(o io.Writer) error { _ = "STUB: not implemented"; return nil }
