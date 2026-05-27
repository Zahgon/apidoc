// SPDX-License-Identifier: MIT

package openapi

import (
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// 将 doc.APIDoc 转换成 openapi
func convert(doc *ast.APIDoc) (*OpenAPI, error) { _ = "STUB: not implemented"; return nil, nil }

func parsePaths(openapi *OpenAPI, d *ast.APIDoc) *core.Error { _ = "STUB: not implemented"; return nil }

// servers
// 不为 PathItem 设置 servers，直接写在 operation

// 找到对应的 doc.Server.URL 值，之后根据此值从 openapi 中取 Server 对象

// requests

// responses

// end for doc.Apis

func setOperationParams(doc *ast.APIDoc, operation *Operation, api *ast.API) {
	_ = "STUB: not implemented"
	return
}

// 将各个类型的 Request 中的报头都集中到 operation.Parameters

func getDescription(desc *ast.Richtext, summary *ast.Attribute) string {
	_ = "STUB: not implemented"
	return ""
}

func setOperation(path *PathItem, method string) (*Operation, *core.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JSON 输出 JSON 格式数据
func JSON(doc *ast.APIDoc) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// YAML 输出 YAML 格式数据
func YAML(doc *ast.APIDoc) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
