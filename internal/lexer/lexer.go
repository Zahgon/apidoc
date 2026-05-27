// SPDX-License-Identifier: MIT

// Package lexer 提供基本的分词功能
package lexer

import (
	"github.com/caixw/apidoc/v7/core"
)

// Lexer 是对一个文本内容的包装，方便 blocker 等接口操作。
type Lexer struct {
	core.Block
	lastIndex int

	// 分别表示当前和之前的定位，可以在某些可撤消的操作之前保存定位信息到 prev
	current, prev Position
}

// New 声明 Lexer 实例
func New(b core.Block) (*Lexer, error) {
	_ = "STUB: not implemented"
	// 以下代码主要保证内容都是合法的 utf8 编码，
	// 这样后续的操作不用再判断每个 utf8.DecodeRune 的调用返回是否都正常。
	return nil, nil
}

// BlockEndPosition 计算 b 的尾部位置
func BlockEndPosition(b core.Block) (Position, error) {
	_ = "STUB: not implemented"
	return *new(Position), nil
}

func (l *Lexer) AtEOF() bool { _ = "STUB: not implemented"; return false }

// Match 接下来的 n 个字符是否匹配指定的字符串，
// 若匹配，则将指定移向该字符串这后，否则不作任何操作。
//
// NOTE: 可回滚该操作
func (l *Lexer) Match(word string) bool { _ = "STUB: not implemented"; return false }

// Current 返回当前在 data 中的偏移量
func (l *Lexer) Current() Position {
	_ = "STUB: not implemented"

	// Move 移动当前的分析器的位置
	//
	// 执行此操作之后，Rollback 将失效
	//
	// 不会限制 p 的值，如果将 p.Offset 的值设置为大于整个数据的长度，
	// 在下次调用 AtEOF 时会返回 true；如果将 p.Offset 设置为负值，则 panic。
	return *new(Position)
}

func (l *Lexer) Move(p Position) { _ = "STUB: not implemented"; return }

// Spaces 获取之后的所有空格，不包含换行符
//
// NOTE: 可回滚该操作
func (l *Lexer) Spaces(exclude rune) []byte { _ = "STUB: not implemented"; return nil }

// 碰到换行符或是非空字符，则中止

// DelimString 查找 delim 并返回到此字符的所有内容
//
// contain 表示是否包含 delim 本身，如果为 false，则返回内容不包含，且该字符串会退回至输入流中，等待下次被读取。
//
// NOTE: 可回滚此操作
func (l *Lexer) DelimString(delim string, contain bool) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Delim 查找 delim 并返回到此字符的所有内容
//
// NOTE: 可回滚此操作
func (l *Lexer) Delim(delim rune, contain bool) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// DelimFunc 查找并返回当前位置到 f 确定位置的所有内容
//
// contain 表示是否包含字符本身，如果为 false，则返回内容不包含，且该字符会退回至输入流中，等待下次被读取。
//
// NOTE: 可回滚此操作
func (l *Lexer) DelimFunc(f func(r rune) bool, contain bool) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Next 返回之后的 n 个字符，或是直到内容结束
//
// NOTE: 可回滚该操作
func (l *Lexer) Next(n int) []byte { _ = "STUB: not implemented"; return nil }

// All 获取当前定位之后的所有内容
//
// NOTE: 可回滚该操作
func (l *Lexer) All() []byte { _ = "STUB: not implemented"; return nil }

// Rollback 回滚上一次的操作
func (l *Lexer) Rollback() { _ = "STUB: not implemented"; return }

// 回滚
// 清空 prev

// Bytes 返回指定范围的内容
//
// NOTE: 并不会改变定位信息
func (l *Lexer) Bytes(start, end int) []byte { _ = "STUB: not implemented"; return nil }
