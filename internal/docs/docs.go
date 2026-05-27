// SPDX-License-Identifier: MIT

// Package docs 打包文档内容
package docs

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/issue9/source"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// FileHeader 表示输出文件的文件头内容
const FileHeader = "该文件由工具自动生成，请勿手动修改！"

// 默认页面
const indexPage = "index.xml"

// 指定在 Handler 中，folder 不为空时，可以访问的文件列表。
//
// 可以以前缀的方式指定，比如：v5/ 表示以 v5/ 开头的所有文件。
var styles = []string{
	"icon.svg",
	ast.MajorVersion + "/",
}

var docsDir = core.FileURI(source.CurrentPath("../../docs"))

// Dir 指向 /docs 的路径
func Dir() core.URI {
	_ = "STUB: not implemented"

	// StylesheetURL 生成 apidoc.xsl 文件的 URL 地址
	//
	// 相对于 docs 目录
	return *new(core.URI)
}

func StylesheetURL(prefix string) string { _ = "STUB: not implemented"; return "" }

// Handler 返回文件服务中间件
//
// 如果 folder 为空，表示采用内嵌的数据作为文件服务；
// stylesheet 是否只返回最基本的样式表相关文件。
func Handler(folder core.URI, stylesheet bool, erro *log.Logger) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func fsHandler(fsys fs.FS, stylesheet bool, erro *log.Logger) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func remoteHandler(url core.URI, stylesheet bool, erro *log.Logger) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func errStatus(w http.ResponseWriter, status int) { _ = "STUB: not implemented"; return }

func errStatusWithError(w http.ResponseWriter, err error, l *log.Logger) {
	_ = "STUB: not implemented"
	return
}

func isStylesheetFile(filename string) bool { _ = "STUB: not implemented"; return false }
