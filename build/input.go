// SPDX-License-Identifier: MIT

package build

import (
	"golang.org/x/text/encoding"

	"github.com/caixw/apidoc/v7/core"
)

// Input 指定输入内容的相关信息。
type Input struct {
	Lang      string   `yaml:"lang"`                // 输入的目标语言，值为 internal/lang 中的 Language.ID
	Dir       core.URI `yaml:"dir"`                 // 源代码目录
	Exts      []string `yaml:"exts,omitempty"`      // 需要扫描的文件扩展名，为空则表示采用默认规则。
	Recursive bool     `yaml:"recursive,omitempty"` // 是否查找 Dir 的子目录
	Encoding  string   `yaml:"encoding,omitempty"`  // 源文件的编码，默认为 UTF-8
	Ignores   []string `yaml:"ignores,omitempty"`   // 忽略的文件或目录，比如 node_modules 等可在此指定

	paths     []core.URI        // 根据 Dir、Exts、Ignores 和 Recursive 生成
	encoding  encoding.Encoding // 根据 Encoding 生成
	sanitized bool
}

func (o *Input) sanitize() error { _ = "STUB: not implemented"; return nil }

// 按 Input 中的规则查找所有符合条件的文件列表并保存至 Input.paths
func (o *Input) recursivePath() error { _ = "STUB: not implemented"; return nil }

func (o *Input) isIgnore(root, path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ParseInputs 分析 opt 中所指定的内容并输出到 blocks
//
// 分析后的内容推送至 blocks 中。
func ParseInputs(blocks chan core.Block, h *core.MessageHandler, opt ...*Input) {
	_ = "STUB: not implemented"
	return
}

// ParseFile 分析 uri 指向的文件并输出到 blocks
func (o *Input) ParseFile(blocks chan core.Block, h *core.MessageHandler, uri core.URI) {
	_ = "STUB: not implemented"
	return
}
