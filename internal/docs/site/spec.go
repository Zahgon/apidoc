// SPDX-License-Identifier: MIT

package site

import (
	"reflect"

	"github.com/caixw/apidoc/v7/internal/node"
)

func (d *doc) newSpec(v any) error { _ = "STUB: not implemented"; return nil }

func (d *doc) dumpToTypes(n *node.Node) error { _ = "STUB: not implemented"; return nil }

// 保证子元素在后显示

func appendItem(t *spec, name string, v reflect.Value, usageKey string, req bool) {
	_ = "STUB: not implemented"
	return
}

func (d *doc) typeExists(typeName string) bool { _ = "STUB: not implemented"; return false }
