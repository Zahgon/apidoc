// SPDX-License-Identifier: MIT

package lang

// 描述了 pascal/delphi 语言的字符串，在 pascal 中
// 转义字符即引号本身，不适合直接在 block 中定义。
type pascalStringBlock struct {
	symbol string
	escape string
}

func newPascalStringBlock(symbol byte) blocker { _ = "STUB: not implemented"; return *new(blocker) }

func (b *pascalStringBlock) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

func (b *pascalStringBlock) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// 转义

// 结束

// end for
