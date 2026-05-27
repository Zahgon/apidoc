// SPDX-License-Identifier: MIT

package cmd

import (
	"io"
	"time"

	"github.com/issue9/cmdopt"

	"github.com/caixw/apidoc/v7"
)

// servers 参数
type (
	servers   map[string]string
	size      apidoc.Range
	slice     []string
	dateRange struct {
		start, end time.Time
	}
)

func (s servers) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (s servers) Set(v string) error { _ = "STUB: not implemented"; return nil }

func (s servers) String() string { _ = "STUB: not implemented"; return "" }

func (r size) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (r *size) Set(v string) (err error) { _ = "STUB: not implemented"; return nil }

func (r *size) String() string { _ = "STUB: not implemented"; return "" }

func (s slice) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (s *slice) Set(v string) (err error) { _ = "STUB: not implemented"; return nil }

func (s *slice) String() string { _ = "STUB: not implemented"; return "" }

func (d dateRange) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (d *dateRange) Set(v string) (err error) { _ = "STUB: not implemented"; return nil }

func (d *dateRange) String() string { _ = "STUB: not implemented"; return "" }

var (
	mockOptions = &apidoc.MockOptions{}

	mockPort         string
	mockServers      = servers{}
	mockStringAlpha  string
	mockPath         = uri("./")
	mockSliceSize    = &size{Min: 5, Max: 10}
	mockNumberSize   = &size{Min: 100, Max: 10000}
	mockStringSize   = &size{Min: 50, Max: 1024}
	mockUsernameSize = &size{Min: 5, Max: 8}
	mockEmailDomains = &slice{"example.com"}
	mockURLDomains   = &slice{"https://example.com"}
	mockDateRange    = &dateRange{}
)

func initMock(command *cmdopt.CmdOpt) { _ = "STUB: not implemented"; return }

func doMock(io.Writer) error { _ = "STUB: not implemented"; return nil }
