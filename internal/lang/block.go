// SPDX-License-Identifier: MIT

package lang

// 接口定义了解析代码块的所有操作
type blocker interface {
	// 确定 l 的当前位置是否匹配 Blocker 的起始位置。
	beginFunc(l *parser) bool

	// 确定 l 的当前位置是否匹配 Blocker 的结束位置
	//
	// data 表示匹配的内容，如果不使用返回的内容，可以返回空值。
	// 比如字符串，只需要返回 true，以确保找到了结束位置，但是 data 可以直接返回 nil。
	//
	// 如果在到达文件末尾都没有找到结束符，则应该返回 nil, false
	endFunc(l *parser) (data []byte, ok bool)
}

type (
	stringBlock struct {
		begin, end, escape string
	}

	singleComment struct {
		begin  string
		begins []byte
	}

	multipleComment struct {
		begin, end           string
		begins, ends, prefix []byte
	}
)

func newString(begin, end, escape string) blocker { _ = "STUB: not implemented"; return *new(blocker) }

func newSingleComment(begin string) blocker { _ = "STUB: not implemented"; return *new(blocker) }

func newMultipleComment(begin, end, prefix string) blocker {
	_ = "STUB: not implemented"
	return *new(blocker)
}

func (b *stringBlock) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

// 从 l 的当前位置开始往后查找，直到找到 b 中定义的 end 字符串，
// 将 l 中的指针移到该位置。
// 正常找到结束符的返回 true，否则返回 false。
//
// 第一个返回参数无用，仅是为了统一函数签名
func (b *stringBlock) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// end for

func (b *singleComment) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

// 从 l 的当前位置往后开始查找连续的相同类型单行代码块。
func (b *singleComment) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// 找不到换行符，直接填充到末尾

// 不是接连着的注释块了，结束当前的匹配

func (b *multipleComment) beginFunc(l *parser) bool { _ = "STUB: not implemented"; return false }

// 从 l 的当前位置一直到定义的 b.End 之间的所有字符。
// 会对每一行应用 filterSymbols 规则。
func (b *multipleComment) endFunc(l *parser) (data []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// 没有找到结束符号，直接到达文件末尾

// 转换单行注释为一个合法的 XML
//
// 主要是去掉了注释符号，比如
//
//	// xx
//
// 会被转换成
//
//	xx
func convertSingleCommentToXML(lines, begin []byte) []byte { _ = "STUB: not implemented"; return nil }

// 零是一个有效的数组下标

// 替换之前字符为空格

// 转换成合法的 XML 格式
//
// 功能与 convertSingleCommentToXML，针对多行注释
func convertMultipleCommentToXML(data, begin, end, chars []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// 替换特殊的符号为空格，使 lines 的内容为一个合法的 xml 文档
func replaceSymbols(lines, chars []byte) []byte { _ = "STUB: not implemented"; return nil }

// 零是一个有效的数组下标

// 替换之前字符为空格
