// SPDX-License-Identifier: MIT

package ast

import (
	"github.com/caixw/apidoc/v7/internal/xmlenc"
)

// Sanitize token.Sanitizer
func (api *API) Sanitize(p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// 报头不能为 object

// 对 Servers 和 Tags 查重

// Sanitize token.Sanitizer
func (e *Enum) Sanitize(p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// Sanitize token.Sanitizer
func (p *Path) Sanitize(pp *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// 路径参数和查询参数不能为 object

func parsePath(path string) (params map[string]struct{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 没有结束符号

// Sanitize token.Sanitizer
func (r *Request) Sanitize(p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// 报头不能为 object

// Sanitize token.Sanitizer
func (p *Param) Sanitize(pp *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// 检测 enums 中的类型是否符合 t 的标准，比如 Number 要求枚举值也都是数值
func chkEnumsType(t *TypeAttribute, enums []*Enum, p *xmlenc.Parser) error {
	_ = "STUB: not implemented"
	return nil
}

func checkDuplicateEnum(enums []*Enum, p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

func checkDuplicateItems(items []*Param, p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

func checkXML(isArray, hasItems bool, xml *XML, p *xmlenc.Parser) error {
	_ = "STUB: not implemented"
	return nil
}

// Sanitize 检测内容是否合法
func (doc *APIDoc) Sanitize(p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// 保证单文件的文档能正常解析

// Sanitize 检测内容是否合法
func (ns *XMLNamespace) Sanitize(p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

func (doc *APIDoc) checkXMLNamespaces(p *xmlenc.Parser) error {
	_ = "STUB: not implemented"
	return nil
}

// 按 URN 查重

// 按 prefix 查重

func (doc *APIDoc) findTag(tag string) *Tag { _ = "STUB: not implemented"; return nil }

func (doc *APIDoc) findServer(srv string) *Server { _ = "STUB: not implemented"; return nil }

func (api *API) sanitizeTags(p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// 检测当前 api 是否与 apidoc.APIs 中存在相同的值
func (api *API) checkDup(p *xmlenc.Parser) { _ = "STUB: not implemented"; return }

// 默认服务器

// 判断是否拥有相同的 server 字段
