// SPDX-License-Identifier: MIT

package apidoc

import (
	"net/http"
	"time"

	"github.com/issue9/rands"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/mock"
)

// Range 表示数值的范围
type Range struct {
	Min, Max int
}

func (r *Range) sanitize() *core.Error { _ = "STUB: not implemented"; return nil }

// MockOptions mock 的一些随机设置项
type MockOptions struct {
	Indent    string            // 缩进字符串
	Servers   map[string]string // 为文档中所有 server 以及对应的路由前缀。
	SliceSize Range             // 指定用于生成数组大小范围的数值

	NumberSize  Range // 指定用于生成数值数据的范围
	EnableFloat bool  // 是否允许生成浮点数

	StringSize  Range  // 指定生成随机字符串的长度范围
	StringAlpha []byte // 指定生成字符串可用的字符

	URLDomains        []string // 指定生成 url 类型数据时可用的域名，默认为 example.com
	EmailDomains      []string // 指定生成 email 类型数据时可用的域名，默认为 example.com
	EmailUsernameSize Range    // 指定生成 email 类型数据的用户名长度范围，默认 [3,8]

	ImageBasePrefix string // 图片的基地址

	DateStart time.Time // 指定生成与时间相关的数值时的最小值
	DateEnd   time.Time // 指定生成与时间相关的数值时的最大值
	dateSize  int64     // 根据 DateStart 和 DateEnd 生成
}

var defaultMockOptions = &MockOptions{
	Indent:    "\t",
	SliceSize: Range{Min: 5, Max: 50},

	NumberSize:  Range{Min: 100, Max: 10000},
	EnableFloat: false,

	StringSize:  Range{Min: 50, Max: 1024},
	StringAlpha: rands.AlphaNumber,

	URLDomains:        []string{"https://example.com/"},
	EmailDomains:      []string{"example.com"},
	EmailUsernameSize: Range{Min: 3, Max: 8},

	ImageBasePrefix: "/__images__",

	DateStart: time.Now().Add(-time.Hour * 24 * 365),
	DateEnd:   time.Now().Add(time.Hour * 24 * 3650),
}

func (o *MockOptions) sanitize() *core.Error { _ = "STUB: not implemented"; return nil }

func (o *MockOptions) gen() (*mock.GenOptions, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *MockOptions) integer() int { _ = "STUB: not implemented"; return 0 }

func (o *MockOptions) float() float32 { _ = "STUB: not implemented"; return 0 }

func (o *MockOptions) url() string { _ = "STUB: not implemented"; return "" }

func (o *MockOptions) email() string { _ = "STUB: not implemented"; return "" }

func (o *MockOptions) image() string { _ = "STUB: not implemented"; return "" }

func (o *MockOptions) date() string { _ = "STUB: not implemented"; return "" }

func (o *MockOptions) time() string { _ = "STUB: not implemented"; return "" }

func (o *MockOptions) dateTime() string { _ = "STUB: not implemented"; return "" }

// Mock 根据文档数据生成 Mock 中间件
//
// data 为文档内容；
// o 用于生成 Mock 数据的随机项，如果为 nil，则会采用默认配置项；
func Mock(h *core.MessageHandler, data []byte, o *MockOptions) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

// MockFile 根据文档生成 Mock 中间件
//
// path 为文档路径；
// o 用于生成 Mock 数据的随机项，如果为 nil，则会采用默认配置项；
func MockFile(h *core.MessageHandler, path core.URI, o *MockOptions) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}
