// SPDX-License-Identifier: MIT

package mock

import (
	"net/http"

	"github.com/issue9/qheader"

	"github.com/caixw/apidoc/v7/internal/ast"
)

func (m *mock) buildAPI(api *ast.API) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// GET、OPTIONS 之类的可能没有 body

func validRequest(ns []*ast.XMLNamespace, requests []*ast.Request, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// 用户提交的 content-type 必须是明确的值

func (m *mock) renderResponse(api *ast.API, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// 仅在 api.Responses 无法匹配任何内容的时候，才从 doc.Responses 中查找内容

// 此时状态码已经输出

// 需要保证 ct 的值不能为空
func findRequestByContentType(requests []*ast.Request, ct string) *ast.Request {
	_ = "STUB: not implemented"
	return nil
}

// accepts 必须是已经按权重进行排序的。
func findResponseByAccept(mimetypes []*ast.Element, requests []*ast.Request, accepts []*qheader.Item) (*ast.Request, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// 表示 requests 中 mimetype 值为空的第一个子项

// 从 requests 中查找是否有符合 accepts 的内容

// 如果存在 none，则从 doc.mimetypes 中查找是否有与 none.Mimetype 相匹配的

// 查看 ct 是否有与 accepts 相匹配的项，必须保证 ct 的值不为空。
func matchContentType(ct string, accepts []*qheader.Item) bool {
	_ = "STUB: not implemented"
	return false
}

// 处理 serveHTTP 中的错误
func (m *mock) handleError(w http.ResponseWriter, r *http.Request, field string, err error) {
	_ = "STUB: not implemented"
	// 这并不是一个真实存在的 URI
	return
}

func validQueries(queries []*ast.Param, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// 默认的 form 格式

// 验证单个参数，仅支持对 query、header 等简单类型的参数验证
func validSimpleParam(p *ast.Param, name, val string) error { _ = "STUB: not implemented"; return nil }

// 字符串的默认值可以为 “”

func (m *mock) buildResponse(p *ast.Request, r *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
