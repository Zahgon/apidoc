// SPDX-License-Identifier: MIT

package lang

const (
	phpHerodoc int8 = iota + 1
	phpNowdoc
)

type phpDocBlock struct {
	token1  string
	token2  string
	doctype int8
}

// herodoc 和 nowdoc 的实现。
//
// http://php.net/manual/zh/language.types.string.php#language.types.string.syntax.heredoc
func newPHPDocBlock() blocker { _ = "STUB: not implemented"; return *new(blocker) }

func (b *phpDocBlock) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

// <<< 之后直接是换行符，则应该退回 <<< 字符

// l.delim 会带上换行符，需要去掉

func (b *phpDocBlock) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
