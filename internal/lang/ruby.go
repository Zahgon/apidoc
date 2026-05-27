// SPDX-License-Identifier: MIT

package lang

// 表示超始和结束符号必须占满一行的情况
type rubyMultipleComment struct {
	begin, end           string
	begins, ends, prefix []byte
}

func newRubyMultipleComment(begin, end, prefix string) blocker {
	_ = "STUB: not implemented"
	return *new(blocker)
}

func (b *rubyMultipleComment) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

// 从 l 的当前位置一直到定义的 b.End 之间的所有字符。
// 会对每一行应用 filterSymbols 规则。
func (b *rubyMultipleComment) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// 没有找到结束符号，直接到达文件末尾
