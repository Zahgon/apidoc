// SPDX-License-Identifier: MIT

package lexer

import (
	"github.com/caixw/apidoc/v7/core"
)

// Position 描述 Lexer 中的定位信息
type Position struct {
	core.Position

	// 表示的是字节的偏移量，
	// 而 Position.Character 表示的是当前行`字符`的偏移量
	Offset int
}

// Equal 判断与 v 是否相等
func (p Position) Equal(v Position) bool { _ = "STUB: not implemented"; return false }

// AddRune 向后移动一个字符
func (p Position) AddRune(r rune) Position { _ = "STUB: not implemented"; return *new(Position) }

// SubRune 向前移动一个字符
func (p Position) SubRune(r rune) Position { _ = "STUB: not implemented"; return *new(Position) }

// 将 p 的定位回滚一个文字 r
//
// 要求 size 的值必须与 r 的字节长度相等。
func (p Position) sub(r rune, size int) Position { _ = "STUB: not implemented"; return *new(Position) }

// add 将 p 的定位往后移动一个字符 r
//
// 要求 size 的值必须与 r 的字节长度相等。
func (p Position) add(r rune, size int) Position { _ = "STUB: not implemented"; return *new(Position) }
