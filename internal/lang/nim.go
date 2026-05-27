// SPDX-License-Identifier: MIT

package lang

type nimRawString struct {
	escape, begin1, begin2, end string
}

type nimMultipleString struct{}

func newNimRawString() blocker { _ = "STUB: not implemented"; return *new(blocker) }

func (s *nimRawString) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

func (s *nimRawString) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// 转义

// 结束

// end for

func newNimMultipleString() blocker { _ = "STUB: not implemented"; return *new(blocker) }

func (s *nimMultipleString) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

func (s *nimMultipleString) endFunc(l *parser) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// 后面没有内容了

// """ 后只有空格和回车符
