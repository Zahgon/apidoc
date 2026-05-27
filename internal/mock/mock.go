// SPDX-License-Identifier: MIT

// Package mock 根据 doc 生成 mock 数据
package mock

import (
	"net/http"

	"github.com/issue9/mux/v7/examples/std"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

type mock struct {
	msgHandler *core.MessageHandler
	doc        *ast.APIDoc
	router     *std.Router
	h          http.Handler
	servers    map[string]string
	indent     string
	gen        *GenOptions
}

// New 声明 Mock 对象
//
// msg 用于处理各类输出消息，仅在 ServeHTTP 中的消息才输出到 msg；
// d doc.APIDoc 实例，调用方需要保证该数据类型的正确性；
// indent 缩进字符串；
// servers 用于指定 d.Servers 中每一个服务对应的路由前缀；
// gen 生成随机数据的函数；
func New(msg *core.MessageHandler, d *ast.APIDoc, indent, imageURL string, servers map[string]string, gen *GenOptions) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

// Load 从本地或是远程加载文档内容
func Load(h *core.MessageHandler, path core.URI, indent, imageURL string, servers map[string]string, gen *GenOptions) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

// 加载并验证

func (m *mock) parse() { _ = "STUB: not implemented"; return }

func (m *mock) ServeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (m *mock) getImage(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func isValidRFC3339Date(val string) bool { _ = "STUB: not implemented"; return false }

func isValidRFC3339Time(val string) bool { _ = "STUB: not implemented"; return false }

func isValidRFC3339DateTime(val string) bool { _ = "STUB: not implemented"; return false }
