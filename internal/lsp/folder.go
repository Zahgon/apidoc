// SPDX-License-Identifier: MIT

package lsp

import (
	"sync"

	"github.com/caixw/apidoc/v7/build"
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
	"github.com/caixw/apidoc/v7/internal/lsp/protocol"
)

// 表示项目文件夹
type folder struct {
	protocol.WorkspaceFolder
	doc       *ast.APIDoc
	cfg       *build.Config
	srv       *server
	loadError error // 加载过程中的出错信息
	noConfig  bool
	h         *core.MessageHandler

	parsedMux sync.RWMutex // 解析 doc 时需要的锁

	// 保存着错误和警告的信息
	diagnostics map[core.URI]*protocol.PublishDiagnosticsParams
}

func (f *folder) close() { _ = "STUB: not implemented"; return }

func (f *folder) messageHandler(msg *core.Message) { _ = "STUB: not implemented"; return }

func (s *server) appendFolders(folders ...protocol.WorkspaceFolder) {
	_ = "STUB: not implemented"
	return
}

// 刷新项目
//
// 默认情况下，没有配置文件不会解析项目，但是在 force 为 true 时，会强制解析项目内容。
func (f *folder) refresh(force bool) { _ = "STUB: not implemented"; return }

// 找不到配置文件

// 仅在不强制刷新项目的情况下，才会将错误保存至 f.loadError

func (s *server) findFolder(uri core.URI) *folder { _ = "STUB: not implemented"; return nil }
