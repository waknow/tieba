package lib

import (
	"regexp"
)

var TitleReg *regexp.Regexp

//post
var PostPanLinkReg *regexp.Regexp
var PostMarkerReplacer *regexp.Regexp
var PostLink *regexp.Regexp

//pan
var PanSourceNameReg *regexp.Regexp
var PanUKReg *regexp.Regexp
var PanIDReg *regexp.Regexp
var PanPathReg *regexp.Regexp
var PanSlashReplacer *regexp.Regexp
var PanSlashUReplacer *regexp.Regexp

func init() {
	TitleReg = regexp.MustCompile("<title>(.*)</title>")

	PanSourceNameReg = regexp.MustCompile("<h2.*title=\"(.*?)\"")
	PanUKReg = regexp.MustCompile("share_uk=\"(\\d+)\"")
	PanIDReg = regexp.MustCompile("share_id=\"(\\d+)\"")
	// raw string: \\/\\u6211\\u7684\\u97f3\\u4e50\\/\\u65e0\\u635f\\u97f3\\u4e50
	PanPathReg = regexp.MustCompile(`\\"path\\":\\"(.*?)\\"`)
	PanSlashReplacer = regexp.MustCompile(`\\\\/`)  //for \\/
	PanSlashUReplacer = regexp.MustCompile(`\\\\u`) //for \\u

	PostPanLinkReg = regexp.MustCompile("http://(pan|yun).baidu.com/[0-9a-zA-Z/?&=;]+")
	PostMarkerReplacer = regexp.MustCompile("&amp;")
	PostLink = regexp.MustCompile("/p/(\\d+)")
}
