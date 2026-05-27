// SPDX-License-Identifier: MIT

package xmlenc

import (
	"golang.org/x/text/message"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/lexer"
)

const (
	cdataStart  = "<![CDATA["
	cdataEnd    = "]]>"
	cdataEscape = "]]]]><![CDATA[>"
)

// Parser 代码块的解析器
type Parser struct {
	*lexer.Lexer
	*core.MessageHandler
}

// NewParser 声明新的 Parser 实例
func NewParser(h *core.MessageHandler, b core.Block) (*Parser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Token 返回下一个 token 对象
//
// token 可能的类型为 *StartElement、*EndElement、*Instruction、*Attribute、*CData、*Comment 和 *String。
// 其中 *String 用于表示 XML 元素的内容。
//
// loc 表示返回的 token 所占的范围；
// 当返回 nil, {}, io.EOF 时，表示已经结束
func (p *Parser) Token() (token any, loc core.Location, err error) {
	_ = "STUB: not implemented"
	return *new(any), *new(core.Location), nil
}

// 记录元素的开始位置

// 当前字符是内容的一部分，返回给 parseContent 解析

func (p *Parser) parseContent() (*String, core.Location, error) {
	_ = "STUB: not implemented"
	return nil, *new(core.Location), nil
}

func (p *Parser) parseComment(pos lexer.Position) (*Comment, core.Location, error) {
	_ = "STUB: not implemented"
	return nil, *new(core.Location), nil
}

// 跳过 --> 三个字符

func (p *Parser) parseStartElement(pos lexer.Position) (*StartElement, core.Location, error) {
	_ = "STUB: not implemented"
	// 跳过空格
	return nil, *new(core.Location), nil
}

func parseName(name []byte, uri core.URI, start, end core.Position) Name {
	_ = "STUB: not implemented"
	return *new(Name)
}

// pos 表示当前元素的起始位置，包含了 < 元素
func (p *Parser) parseEndElement(pos lexer.Position) (*EndElement, core.Location, error) {
	_ = "STUB: not implemented"
	// 名称开始的定位，传递过来的 pos 表示的 < 起始位置
	return nil, *new(core.Location), nil
}

// 去掉 > 符号

// pos 表示当前元素的起始位置，包含了 < 元素
func (p *Parser) parseCData(pos lexer.Position) (*CData, core.Location, error) {
	_ = "STUB: not implemented"
	return nil, *new(core.Location), nil
}

// 回滚两个字符，用于匹配转义内容

// 将 ]]> 从流中去掉

// pos 表示当前元素的起始位置，包含了 < 元素
func (p *Parser) parseInstruction(pos lexer.Position) (*Instruction, core.Location, error) {
	_ = "STUB: not implemented"
	return nil, *new(core.Location), nil
}

func (p *Parser) parseAttributes() (attrs []*Attribute, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parseAttribute() (*Attribute, error) {
	_ = "STUB: not implemented"
	// 忽略空格
	return nil, nil
}

// 不包含 " 符号

func (p *Parser) getName() ([]byte, core.Range) {
	_ = "STUB: not implemented"
	return nil, *new(core.Range)
}

// 找到与 start 相对应的结束符号位置
//
// 如果找不到对应的结束符号，则会向 p.h 输出一条错误信息，然后将定位至原始位置，并不返回错误。
// 这样可以保证最大限度地解析 xml 内容，不会因为一些非致命的错误而中断整个解析。
func (p *Parser) endElement(start *StartElement) error { _ = "STUB: not implemented"; return nil }

// newError 生成 *core.Error 对象，其中的 URI 来自于 p.Location.URI。
func (p *Parser) newError(start, end core.Position, field string, key message.Reference, v ...any) *core.Error {
	_ = "STUB: not implemented"
	return nil
}
