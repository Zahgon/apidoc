// SPDX-License-Identifier: MIT

package mock

import (
	"github.com/caixw/apidoc/v7/internal/ast"
)

// GenOptions 生成随机数据的函数
type GenOptions struct {
	// 返回一个随机的数值
	//
	// 可以是浮点和整数类型。
	Number func(p *ast.Param) any

	// 返回一个随机长度的字符串
	String func(p *ast.Param) string

	// 返回一个随机的布尔值
	Bool func() bool

	// 返回一个随机的数值
	//
	// 该数值被用于声明 slice 长度，所以必须为正整数。
	SliceSize func() int

	// 返回一个介于 [0, max] 之间的数值
	//
	// 该数值被用于从数组中获取其中的某个元素。
	Index func(max int) int
}

func isEnum(p *ast.Param) bool { _ = "STUB: not implemented"; return false }

func (g *GenOptions) generateBool() bool { _ = "STUB: not implemented"; return false }

func (g *GenOptions) generateNumber(p *ast.Param) any { _ = "STUB: not implemented"; return *new(any) }

// 这属于文档定义错误，直接 panic

func (g *GenOptions) generateString(p *ast.Param) string { _ = "STUB: not implemented"; return "" }

func (g *GenOptions) generateSliceSize() int { _ = "STUB: not implemented"; return 0 }
