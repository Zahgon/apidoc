// SPDX-License-Identifier: MIT

package lang

// swift 嵌套风格的块注释。会忽略掉内嵌的注释块。
type swiftNestMCommentBlock struct {
	begin  string
	end    string
	prefix []byte // 需要过滤的前缀
	begins []byte
	ends   []byte
	level  int8
}

// prefix 表示每一行的前缀符号，比如：
//
//	/*
//	 *
//	 */
//
// 中的 * 字符
func newSwiftNestMCommentBlock(begin, end, prefix string) blocker {
	_ = "STUB: not implemented"
	return *new(blocker)
}

func (b *swiftNestMCommentBlock) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

func (b *swiftNestMCommentBlock) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
