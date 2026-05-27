// SPDX-License-Identifier: MIT

package openapi

import (
	"github.com/caixw/apidoc/v7/core"
	"github.com/caixw/apidoc/v7/internal/ast"
)

// Info 接口文档的基本信息
type Info struct {
	Title          string   `json:"title" yaml:"title"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

// Contact 描述联系方式
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License 授权信息
type License struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

func (info *Info) sanitize() *core.Error { _ = "STUB: not implemented"; return nil }

func (l *License) sanitize() *core.Error { _ = "STUB: not implemented"; return nil }

func newLicense(l *ast.Link) *License { _ = "STUB: not implemented"; return nil }

func newContact(c *ast.Contact) *Contact { _ = "STUB: not implemented"; return nil }

func (c *Contact) sanitize() *core.Error { _ = "STUB: not implemented"; return nil }
