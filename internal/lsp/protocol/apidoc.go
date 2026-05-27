// SPDX-License-Identifier: MIT

package protocol

import (
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// APIDocDetectParams apidoc/detect 的请求参数
type APIDocDetectParams struct {
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Recursive 是否检测子目录的内容
	Recursive bool `json:"recursive,omitempty"`
}

// APIDocDetectResult apidoc/detect 的返回参数
type APIDocDetectResult struct {
	Error string `json:"error,omitempty"` // 如果生成配置文件有误，返回此字段。
}

// APIDocOutline 传递给客户端的文档摘要
//
// 这不是一个标准的 LSP 数据结构，由 apidoc 自定义，
// 用户由服务端向客户端发送当前的文档结构信息。
type APIDocOutline struct {
	WorkspaceFolder

	Err      string `json:"err,omitempty"`      // 表示项目解析出问题，此值不为空，则除去 WorkspaceFolder 和 Err 之外的字段都是无意义的。
	NoConfig bool   `json:"noConfig,omitempty"` // 没有配置文件的相关信息

	Location core.Location   `json:"location,omitempty"`
	Title    string          `json:"title,omitempty"`
	Version  string          `json:"version,omitempty"`
	Tags     []*APIDocTag    `json:"tags,omitempty"`
	Servers  []*APIDocServer `json:"servers,omitempty"`

	APIs []*API `json:"apis,omitempty"`
}

// APIDocTag 文档支持的标签属性
type APIDocTag struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// APIDocServer 文档支持的服务器
type APIDocServer struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// API 描述单个 API 的信息
type API struct {
	Location   core.Location `json:"location"`
	Method     string        `json:"method"`
	Path       string        `json:"path"`
	Tags       []string      `json:"tags,omitempty"`
	Servers    []string      `json:"servers,omitempty"`
	Deprecated string        `json:"deprecated,omitempty"`
	Summary    string        `json:"summary,omitempty"`
}

// BuildAPIDocOutline 根据 ast.APIDoc 构建 APIDoc
//
// 如果 doc 不是一个有效的文档内容，比如是零值，则返回 nil。
// 如果是 doc.APIs 中的某一个元素的 path 未必须，则会忽略此记录的显示。
func BuildAPIDocOutline(f WorkspaceFolder, doc *ast.APIDoc) *APIDocOutline {
	_ = "STUB: not implemented"
	return nil
}

func (o *APIDocOutline) appendAPI(api *ast.API) { _ = "STUB: not implemented"; return }

// 获取 API 的路由地址，如果为空使用 ？代替
