// SPDX-License-Identifier: MIT

package mock

import (
	"encoding/json"

	"github.com/issue9/errwrap"

	"github.com/caixw/apidoc/v7/internal/ast"
)

type jsonValidator struct {
	param *ast.Param

	// 按顺序表示的状态
	// 可以是 [ 表示在数组中，{ 表示在对象，: 表示下一个值必须是属性，空格表示其它状态
	states []byte

	names []string // 按顺序保存变量名称
}

func validJSON(p *ast.Request, content []byte) error { _ = "STUB: not implemented"; return nil }

func newJSONValidator(r *ast.Request) *jsonValidator { _ = "STUB: not implemented"; return nil }

// 状态有默认值

func (validator *jsonValidator) valid(d *json.Decoder) error { _ = "STUB: not implemented"; return nil }

// 正常结束

// 对应 JSON null

// json string

// 字符串类型的值

// 表示数据为单个值，比如 "str"

// case ']', '}': // 格式错误，由 json.Valid 保证
// case '{' 属性名

// [、]、{、}

// {xx: [] } 类似这种格式，需要同时弹出两个状态

// json bool

// json number

// 只有键值对结束时，才弹出键名

// 如果 t == "" 表示不需要验证类型，比如 null 可以赋值给任何类型
func (validator *jsonValidator) validValue(t string, v any) error {
	_ = "STUB: not implemented"
	return nil
}

// 可能是相对站点的根路径，不作类型检测
// 数值类型都被 json 解释为 float64，无法判断值是浮点还是整数。

// 返回当前的状态
func (validator *jsonValidator) state() byte { _ = "STUB: not implemented"; return 0 }

func (validator *jsonValidator) pushState(state byte) { _ = "STUB: not implemented"; return }

func (validator *jsonValidator) popState() { _ = "STUB: not implemented"; return }

func (validator *jsonValidator) pushName(name string) { _ = "STUB: not implemented"; return }

func (validator *jsonValidator) popName() { _ = "STUB: not implemented"; return }

// 如果 names 为空，返回 validator.param
func (validator *jsonValidator) find() *ast.Param { _ = "STUB: not implemented"; return nil }

type jsonBuilder struct {
	w      *errwrap.Buffer
	deep   int
	indent string // 单次的缩进
}

func buildJSON(p *ast.Request, indent string, g *GenOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (builder *jsonBuilder) encode(p *ast.Param, chkArray bool, g *GenOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (builder *jsonBuilder) writeIndent() *jsonBuilder { _ = "STUB: not implemented"; return nil }

// v 只能是基本类型
func (builder *jsonBuilder) writeValue(v any) *jsonBuilder { _ = "STUB: not implemented"; return nil }
