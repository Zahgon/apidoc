// SPDX-License-Identifier: MIT

package cmd

import (
	"io"

	"github.com/issue9/cmdopt"

	"github.com/caixw/apidoc/v7/core"
)

var syntaxDir uri = uri(core.FileURI("./"))

func initSyntax(command *cmdopt.CmdOpt) { _ = "STUB: not implemented"; return }

func syntax(w io.Writer) error { _ = "STUB: not implemented"; return nil }
