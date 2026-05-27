// SPDX-License-Identifier: MIT

package cmd

import (
	"io"

	"github.com/issue9/cmdopt"
)

var versionKind string

func initVersion(command *cmdopt.CmdOpt) { _ = "STUB: not implemented"; return }

func version(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// all 与 default 采取相同的输出
