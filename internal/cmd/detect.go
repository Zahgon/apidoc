// SPDX-License-Identifier: MIT

package cmd

import (
	"io"

	"github.com/issue9/cmdopt"
)

var (
	detectRecursive bool
	detectWrite     bool
	detectDir       = uri("./")
)

func initDetect(command *cmdopt.CmdOpt) { _ = "STUB: not implemented"; return }

func detect(w io.Writer) error { _ = "STUB: not implemented"; return nil }
