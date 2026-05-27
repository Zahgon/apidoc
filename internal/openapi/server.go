// SPDX-License-Identifier: MIT

package openapi

import (
	"strings"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// 去掉 URL 中的 {} 模板参数。使其符合 is.URL 的判断规则
var urlreplace = strings.NewReplacer("{", "", "}", "")

// Server 服务器描述信息
type Server struct {
	URL         string                     `json:"url" yaml:"url"`
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// ServerVariable Server 中 URL 模板中对应的参数变量值
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default" yaml:"default"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

func newServer(srv *ast.Server) *Server { _ = "STUB: not implemented"; return nil }

func (srv *Server) sanitize() *core.Error { _ = "STUB: not implemented"; return nil }

// 可以是 / 未必是一个 URL

func (v *ServerVariable) sanitize() *core.Error { _ = "STUB: not implemented"; return nil }
