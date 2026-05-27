// SPDX-License-Identifier: MIT

package build

import (
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/lang"
)

// DetectConfig 检测 wd 内容并生成 Config 实例
//
// wd 只能为本地文件系统；
// recursive 是否检测子目录；
func DetectConfig(wd core.URI, recursive bool) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 检测指定目录下的内容，并为其生成一个合适的 Input 实例。
//
// 检测依据为根据扩展名来做统计，数量最大且被支持的获胜。
func detectInput(dir string, recursive bool) ([]*Input, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type language struct {
	lang.Language
	count int
}

// 根据 exts 计算每个语言对应的文件数量，并按倒序返回
//
// exts 参数为从 detectExts 中获取的返回值
func detectLanguage(exts map[string]int) []*language { _ = "STUB: not implemented"; return nil }

// end for

// 返回 dir 目录下文件类型及对应的文件数量的一个集合。
// recursive 表示是否查找子目录。
func detectExts(dir string, recursive bool) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
