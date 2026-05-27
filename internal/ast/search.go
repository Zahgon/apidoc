// SPDX-License-Identifier: MIT

package ast

import (
	"reflect"

	"github.com/caixw/apidoc/v7/core"
)

var searcherType = reflect.TypeOf((*core.Searcher)(nil)).Elem()

// Search 搜索符合条件的对象并返回
//
// 从 doc 中查找符合符合 pos 定位的最小对象，且该对象必须实现了 t 类型。
// 如果不存在则返回 nil。t 必须是一个接口。
func (doc *APIDoc) Search(uri core.URI, pos core.Position, t reflect.Type) (r core.Searcher) {
	_ = "STUB: not implemented"
	return *new(core.Searcher)
}

// apidoc 的 uri 可以与 api 的 uri 不同

func search(v reflect.Value, uri core.URI, pos core.Position, t reflect.Type) (r core.Searcher) {
	_ = "STUB: not implemented"
	return *new(core.Searcher)
}

// 不匹配当前元素，也不需要搜查子元素是否实现 t，则直接返回 nil。
