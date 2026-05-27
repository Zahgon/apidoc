// SPDX-License-Identifier: MIT

package build

import (
	"github.com/caixw/apidoc/v7/core"
)

// 获取 path 的绝对路径
//
// 如果 path 是相对路径的，则将其设置为相对于 wd 的路径。
func abs(path, wd core.URI) (uri core.URI, err error) {
	_ = "STUB: not implemented"
	return *new(core.URI), nil
}

// 相对路径

// 获取 path 相对于 wd 的路径
//
// 如果两者不存在关联性，则返回 path 的原始值。
// 返回值仅为普通的路径表示，不会带 scheme 内容。
func rel(path, wd core.URI) (uri core.URI, err error) {
	_ = "STUB: not implemented"
	return *new(core.URI), nil
}
