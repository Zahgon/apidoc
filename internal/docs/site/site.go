// SPDX-License-Identifier: MIT

// Package site 用于生成网站内容
//
// 包括网站的基本信息，以及文档的翻译内容等。
package site

import (
	"github.com/caixw/apidoc/v7/core"
)

const (
	siteFilename = "site.xml" // 配置文件的文件名
	docBasename  = "locale."  // 翻译文档文件名的前缀部分，一般格式为 docBasename.{locale}.xml
)

type site struct {
	XMLName struct{} `xml:"site"`

	Name      string     `xml:"name"`
	Version   string     `xml:"version"`
	Repo      string     `xml:"repo"`
	URL       string     `xml:"url"`
	Languages []language `xml:"languages>language"`
	Locales   []loc      `xml:"locales>locale"`
}

type language struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type loc struct {
	ID    string `xml:"id,attr"`
	Href  string `xml:"href,attr"`
	Title string `xml:"title,attr"`
	Doc   string `xml:"doc,attr"`
}

type doc struct {
	XMLName  struct{}   `xml:"locale"`
	Spec     []*spec    `xml:"spec>type"`
	Commands []*command `xml:"commands>command"`
	Config   []*item    `xml:"config>item"`
}

type spec struct {
	Name  string   `xml:"name,attr,omitempty"`
	Usage innerXML `xml:"usage,omitempty"`
	Items []*item  `xml:"item,omitempty"`
}

type innerXML struct {
	Text string `xml:",innerxml"`
}

type item struct {
	Name     string `xml:"name,attr"` // 变量名
	Type     string `xml:"type,attr"` // 变量的类型
	Array    bool   `xml:"array,attr"`
	Required bool   `xml:"required,attr"`
	Usage    string `xml:",innerxml"`
}

type command struct {
	Name  string `xml:"name,attr"`
	Usage string `xml:",innerxml"`
}

// Write 输出站点中所有需要自动生成的内容
func Write(target core.URI) error { _ = "STUB: not implemented"; return nil }

func writeXML(uri core.URI, v any, indent string) error { _ = "STUB: not implemented"; return nil }

// 统一代码风格，文件末尾加一空行。

func gen() (*site, map[string]*doc, error) { _ = "STUB: not implemented"; return nil, nil, nil }

func genDoc() (*doc, error) { _ = "STUB: not implemented"; return nil, nil }

func buildDocFilename(id string) string { _ = "STUB: not implemented"; return "" }
