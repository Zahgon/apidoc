// SPDX-License-Identifier: MIT

// Package asttest 提供了一个合法的 ast.APIDoc 对象
package asttest

import (
	"github.com/issue9/assert/v3"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// Filename 文档的文件名
const Filename = "index.xml"

// Get 返回 doc.APIDoc 对象
//
// 同时当前目录下的 index.xml 文件与此返回对象内容是相同的。
func Get() *ast.APIDoc { _ = "STUB: not implemented"; return nil }

// XML 获取 Get 返回对象的 XML 编码
func XML(a *assert.Assertion) []byte { _ = "STUB: not implemented"; return nil }

// URI 返回测试文件基于 URI 的表示方式
func URI(a *assert.Assertion) core.URI { _ = "STUB: not implemented"; return *new(core.URI) }

// Path 返回测试文件的绝对路径
//
// NOTE: 该文件与 Get() 对象的内容是相同的。
func Path(a *assert.Assertion) string { _ = "STUB: not implemented"; return "" }

// Dir 返回测试文件所在的目录
func Dir(a *assert.Assertion) string { _ = "STUB: not implemented"; return "" }

func pp(a *assert.Assertion, p string) string { _ = "STUB: not implemented"; return "" }
