// SPDX-License-Identifier: MIT

package build

import (
	"bytes"

	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// 几种输出的类型
const (
	APIDocXML   = "apidoc+xml"
	OpenapiYAML = "openapi+yaml"
	OpenapiJSON = "openapi+json"
)

type marshaler func(*ast.APIDoc) ([]byte, error)

// Output 指定了渲染输出的相关设置项。
type Output struct {
	// 文档的版本号
	//
	// 该值会覆盖文档中 apidoc.version 的值，方便用户通过代码层面进行版本号同步，
	// 该值无法通过配置文件设置，只能由代码进行设置。
	Version string `yaml:"-"`

	// 导出的文件类型格式，默认为 apidoc 的 XML 文件。
	Type string `yaml:"type,omitempty"`

	// 文档的保存路径
	//
	// 仅适用本地路径
	Path core.URI `yaml:"path"`

	// 只输出该标签的文档，若为空，则表示所有。
	Tags []string `yaml:"tags,omitempty"`

	// xslt 文件地址
	//
	// 默认值为 https://apidoc.tools/docs/ 下当前版本的 apidoc.xsl，比如：
	//  https://apidoc.tools/docs/v7/apidoc.xsl
	//
	// NOTE: 仅针对 xml 类型的输出文件
	Style string `yaml:"style,omitempty"`

	// 命名空间的相关设置
	//
	// 当 namespace 为 true 时会在文档中输出以 core.XMLNamespace 作为命名空间的值，
	// 如果还指定了 NamespacePrefix 则会以此值作为前缀值。
	// NamespacePrefix 仅在 Namespace 为 true 时才启作用。
	//
	// NOTE: 仅针对 Type = APIDocXML
	Namespace       bool   `yaml:"namespace,omitempty"`
	NamespacePrefix string `yaml:"namespace-prefix,omitempty"`

	procInst []string  // 保存所有 xml 的指令内容，包括编码信息
	marshal  marshaler // Type 对应的转换函数
	xml      bool      // 是否为 xml 内容
}

func (o *Output) contains(tags ...string) bool { _ = "STUB: not implemented"; return false }

func (o *Output) sanitize() error { _ = "STUB: not implemented"; return nil }

func (o *Output) apidocMarshaler(d *ast.APIDoc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Output) buffer(d *ast.APIDoc) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterDoc(d *ast.APIDoc, o *Output) { _ = "STUB: not implemented"; return }
