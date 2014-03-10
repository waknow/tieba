package lib

import (
	"fmt"
	"net/url"
)

type PanInfo struct {
	Url     *url.URL
	Name    string
	Title   string
	path    string
	uk      string
	shareId string
	Refer   Post
}

func (p PanInfo) String() string {
	return fmt.Sprintf("Name:%s\nTitle:%s\nUrl:%s\nRefer:%s\n",
		p.Name,
		p.Title,
		p.Url.String(),
		p.Refer.Url)
}

func GetPanInfo(p Post) (infos []PanInfo) {

	for _, link := range p.Links {
		s, err := HttpGet(link.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		name := Group(panSourceNameReg.FindStringSubmatch(s)).Get(1)
		title := Group(titleReg.FindStringSubmatch(s)).Get(1)
		uk := Group(panUKReg.FindStringSubmatch(s)).Get(1)
		id := Group(panIDReg.FindStringSubmatch(s)).Get(1)

		path := Group(panPathReg.FindStringSubmatch(s)).Get(1)
		path = panSlashReplacer.ReplaceAllString(path, "/")
		path = panSlashUReplacer.ReplaceAllString(path, "\\u")
		path = PathToUnicode(path)
		info := PanInfo{
			Url:     link,
			Name:    name,
			Title:   title,
			path:    path,
			uk:      uk,
			shareId: id,
			Refer:   p,
		}
		infos = append(infos, info)
	}

	return
}
