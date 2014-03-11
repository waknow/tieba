package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

//描述网盘信息的结构体
//包含有网盘url，名称，网页标题等
type PanInfo struct {
	Url     *url.URL
	Name    string
	Title   string
	path    string
	uk      string
	shareId string
	refer   Post
}

//抓取指定url网盘的信息，并返回
func GetPanInfo(u *url.URL) (info []PanInfo) {
	s, err := HttpGet(u.String())
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
	info = PanInfo{
		Url:     link,
		Name:    name,
		Title:   title,
		path:    path,
		uk:      uk,
		shareId: id,
		refer:   p,
	}
	return
}

//扫描一个Post对象中所有网盘的地址的信息，并返回
func GetPanInfos(p Post) (infos []PanInfo) {

	for _, link := range p.Links {
		info := GetPanInfo(link.String())
		infos = append(infos, info)
	}
}

//从网盘的首页获取json，并解码到结构体PanFile中
func getFileFromHomePage(u *url.URL) (files []PanFile) {
	s, err := HttpGet(u.String())
	if err != nil {
		log.Println("Package Pan:", err)
		return
	}

	jsons := panHPJson.FindAllString(s)
	f := PanFile{}
	for _, j := range jsons {
		j = panHPJsoneplacer.ReplaceAllString(j, `"`)
		err = json.Unmarshal([]byte(j), &f)
		if err != nil {
			log.Println("Package Pan:", err)
			return
		}
		files = append(f)
	}
	return
}

//向特定的url请求数据，获取json并解析到PanFile结构体中
func getFileFromRequest(path, uk, id string) (files []PanFile) {
	requeUrl := url.Parse("http://pan.baidu.com/share/list")
	var values url.Values
	values.Set("dir", path)
	values.Set("uk", uk)
	values.Set("shareid", id)
	values.Set("order", time)
	values.Set("desc", 1)
	requeUrl.RawQuery = values.Encode()
}

//从网盘中扫描所有的文件信息，并返回
func scanFile(u *url.URL) (files []PanFile) {

}

func (p *PanInfo) Path() string    { return p.path }
func (p *PanInfo) UK() string      { return p.uk }
func (p *PanInfo) ShareID() string { return p.shareId }
func (p *PanInfo) Refer() Post     { return p.refer }
func (p PanInfo) String() string {
	return fmt.Sprintf("Name:%s\nTitle:%s\nUrl:%s\nRefer:%s\n",
		p.Name,
		p.Title,
		p.Url.String(),
		p.Refer.Url)
}

type PanFile struct {
	Name       string `json:"server_filename"`
	Size       int    `json:"size"`
	Path       string `json:"path"`
	Parent     string `json:"parent_path"`
	IsDir      int    `json:"isdir"`
	ReferTitle string
	ReferUrl   *url.URL
}
