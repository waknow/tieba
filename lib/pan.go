package lib

import (
	"fmt"
	"net/url"
)

type PanInfo struct {
	Url     *url.URL
	Name    string
	Title   string
	Path    string
	UK      string
	ShareId string
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

		name := Group(PanSourceNameReg.FindStringSubmatch(s)).Get(1)
		fmt.Println(name)
		title := Group(TitleReg.FindStringSubmatch(s)).Get(1)
		uk := Group(PanUKReg.FindStringSubmatch(s)).Get(1)
		id := Group(PanIDReg.FindStringSubmatch(s)).Get(1)

		path := Group(PanPathReg.FindStringSubmatch(s)).Get(1)
		path = PanSlashReplacer.ReplaceAllString(path, "/")
		path = PanSlashUReplacer.ReplaceAllString(path, "\\u")
		path = PathToUnicode(path)
		info := PanInfo{
			Url:     link,
			Name:    name,
			Title:   title,
			Path:    path,
			UK:      uk,
			ShareId: id,
			Refer:   p,
		}
		infos = append(infos, info)
	}

	return
}
