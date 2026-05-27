// SPDX-License-Identifier: MIT

package core

import (
	"golang.org/x/text/encoding"
)

// 目前 URI 支持的协议
const (
	SchemeFile  = "file"
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"

	separator = "://"
)

// URI 定义 [URI]
//
//	  foo://example.com:8042/over/there?name=ferret#nose
//	  \_/   \______________/\_________/ \_________/ \__/
//	   |           |            |            |        |
//	scheme     authority       path        query   fragment
//	   |   _____________________|__
//	  / \ /                        \
//	  urn:example:animal:ferret:nose
//
// 如果是本地相对路径，也可以直接使用 `./path/file` 的形式表示，
// 不需要指定协议。
//
// NOTE: 并非完整的 URI 实现，仅作为了 file:// 和 http:// 支持，
// 也提供对 windows 路径的支持。
//
// [URI]: http://tools.ietf.org/html/rfc3986
type URI string

// FileURI 根据本地文件路径构建 URI 实例
//
// 如果已经存在协议，则不作任何改变返回。
func FileURI(path string) URI { _ = "STUB: not implemented"; return *new(URI) }

func (uri *URI) UnmarshalJSON(v []byte) error { _ = "STUB: not implemented"; return nil }

// File 返回 file:// 协议关联的文件路径
func (uri URI) File() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (uri URI) String() string {
	_ = "STUB: not implemented"

	// Append 追加 path 至 URI 生成新的 URI
	return ""
}

func (uri URI) Append(path string) URI { _ = "STUB: not implemented"; return *new(URI) }

func isPathSeparator(b byte) bool { _ = "STUB: not implemented"; return false }

// Exists 判断 uri 指向的内容是否存在
//
// 如果是非本地文件，通过 http 的状态码是否为 400 以内加以判断。
func (uri URI) Exists() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ReadAll 以 enc 编码读取 uri 的内容
//
// 目前仅支持 file、http 和 https 协议
func (uri URI) ReadAll(enc encoding.Encoding) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteAll 写入内容至 uri
func (uri URI) WriteAll(data []byte) error { _ = "STUB: not implemented"; return nil }

// Parse 分析 uri，获取其各个部分的内容
func (uri URI) Parse() (schema, path string) { _ = "STUB: not implemented"; return "", "" }

func remoteFileIsExists(url string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// 以指定的编码方式读取本地文件内容
func readLocalFile(path string, enc encoding.Encoding) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 以指定的编码方式读取远程文件内容
func readRemoteFile(url string, enc encoding.Encoding) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
