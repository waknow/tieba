package page

import (
	"fmt"
	"net/url"

	"tieba/lib"
)

type PanInfo struct {
	Url     *url.URL
	Name    string
	Title   string
	Path    string
	Refer   Post
	UK      string
	ShareId string
}

func (p *PanInfo) String() string {
	return fmt.Sprintf("Name:%s\nTitle:%s\nUrl:%s\nRefer:%s\n",
		p.Name,
		p.Title,
		p.Url.String(),
		p.Refer.Url)
}

func GetPanInfo(p Post) (infos []PanInfo) {

	for _, link := range p.Links {
		fmt.Println(link.String())
		s, err := lib.HttpGet(link.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		name := Group(lib.PanSourceNameReg.FindStringSubmatch(s)).Get(1)
		title := Group(lib.TitleReg.FindStringSubmatch(s)).Get(1)
		uk := Group(lib.PanUKReg.FindStringSubmatch(s)).Get(1)
		id := Group(lib.PanIDReg.FindStringSubmatch(s)).Get(1)

		path := Group(lib.PanPathReg.FindStringSubmatch(s)).Get(1)
		path = lib.PanSlashReplacer.ReplaceAllString(path, "/")
		path = lib.PanSlashUReplacer.ReplaceAllString(path, "\\u")
		// fmt.Printf("uk:%s,id:%s,path:%s\n", uk, id, path)
		info := PanInfo{
			Url:     link,
			Name:    name,
			Title:   title,
			Refer:   p,
			Path:    path,
			UK:      uk,
			ShareId: id,
		}
		infos = append(infos, info)
	}

	return
}
