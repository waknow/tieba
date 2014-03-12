package lib

import (
	"regexp"
)

var titleReg *regexp.Regexp

//post
var postPanLinkReg *regexp.Regexp
var postMarkerReplacer *regexp.Regexp
var postLink *regexp.Regexp

//pan
var panSourceNameReg *regexp.Regexp
var panUKReg *regexp.Regexp
var panIDReg *regexp.Regexp
var panPathReg *regexp.Regexp
var panSlashReplacer *regexp.Regexp
var panSlashUReplacer *regexp.Regexp
var panHPJsonReg *regexp.Regexp
var panHPJsonReplacer *regexp.Regexp

func init() {
	titleReg = regexp.MustCompile("<title>(.*)</title>")

	panSourceNameReg = regexp.MustCompile("<h2.*title=\"(.*?)\"")
	panUKReg = regexp.MustCompile("share_uk=\"(\\d+)\"")
	panIDReg = regexp.MustCompile("share_id=\"(\\d+)\"")
	// raw string: \\/\\u6211\\u7684\\u97f3\\u4e50\\/\\u65e0\\u635f\\u97f3\\u4e50
	panPathReg = regexp.MustCompile(`\\"parent_path\\":\\"(.*?)\\"`)
	panSlashReplacer = regexp.MustCompile(`\\\\/`)  //for \\/
	panSlashUReplacer = regexp.MustCompile(`\\\\u`) //for \\u
	panHPJsonReg = regexp.MustCompile(`{\\"fs_id\\".*?}`)
	panHPJsonReplacer = regexp.MustCompile(`\\"`)

	postPanLinkReg = regexp.MustCompile("http://(pan|yun).baidu.com/[0-9a-zA-Z/?&=;]+")
	postMarkerReplacer = regexp.MustCompile("&amp;")
	postLink = regexp.MustCompile("/p/(\\d+)")
}
