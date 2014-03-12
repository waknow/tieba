package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

func (p *PanInfo) Path() string    { return p.path }
func (p *PanInfo) UK() string      { return p.uk }
func (p *PanInfo) ShareID() string { return p.shareId }
func (p *PanInfo) Refer() Post     { return p.refer }
func (p PanInfo) String() string {
	return fmt.Sprintf("Name:%s\nTitle:%s\nUrl:%s\nRefer:%s\n",
		p.Name,
		p.Title,
		p.Url.String(),
		p.refer.Url.String())
}

//抓取指定url网盘的信息，并返回
func GetPanInfo(u *url.URL) (info PanInfo) {
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
		Url:     u,
		Name:    name,
		Title:   title,
		path:    path,
		uk:      uk,
		shareId: id,
	}
	return
}

//扫描一个Post对象中所有网盘的地址的信息，并返回
func GetPanInfos(p Post) (infos []PanInfo) {
	for _, link := range p.Links {
		info := GetPanInfo(link)
		info.refer = p
		infos = append(infos, info)
	}
	return
}

type hpJson struct {
	Name   string `json:"server_filename"`
	Size   string `json:"size"`
	Path   string `json:"path"`
	Parent string `json:"parent_path"`
	IsDir  string `json:"isdir"`
}

func (h *hpJson) toPanFile() (p PanFile) {
	p.Name = h.Name
	p.Path = h.Path
	p.Parent = h.Parent
	p.IsDir = (h.IsDir == "1")

	v, err := strconv.ParseInt(h.Size, 10, 32)
	if err != nil {
		log.Println(err)
	} else {
		p.Size = v
	}
	return
}

//从网盘的首页获取json，并解码到结构体PanFile中
func getFileFromHomePage(u *url.URL) (files []PanFile) {
	s, err := HttpGet(u.String())
	if err != nil {
		log.Println("Package Pan:", err)
		return
	}

	jsons := panHPJsonReg.FindAllString(s, -1)
	f := hpJson{}
	for _, j := range jsons {
		j = panHPJsonReplacer.ReplaceAllString(j, `"`)
		err = json.Unmarshal([]byte(j), &f)
		if err != nil {
			log.Println("Package Pan:", err)
			return
		}
		files = append(files, f.toPanFile())
	}
	return
}

type requestJson struct {
	Name   string `json:"server_filename"`
	Size   int64  `json:"size"`
	Path   string `json:"path"`
	Parent string `json:"parent"`
	IsDir  int    `json:"size"`
}

func (r *requestJson) toPanFile() (f PanFile) {
	f.Name = r.Name
	f.Size = r.Size
	f.Path = r.Path
	f.IsDir = (r.IsDir == 1)
	return
}

//向特定的url请求数据，获取json并解析到PanFile结构体中
func getFileFromRequest(path, uk, id string) (files []PanFile) {
	log.Println(path)
	log.Println(uk)
	log.Println(id)
	requeUrl, _ := url.Parse("http://pan.baidu.com/share/list")
	values := url.Values{}
	values.Set("dir", path)
	values.Set("uk", uk)
	values.Set("shareid", id)
	values.Set("order", "time")
	values.Set("desc", "1")
	requeUrl.RawQuery = values.Encode()

	log.Println(requeUrl.String())

	resp, err := http.Get(requeUrl.String())
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var list struct {
		List []requestJson `json:"list"`
	}
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range list.List {
		files = append(files, v.toPanFile())
	}
	return
}

// //从网盘中扫描所有的文件信息，并返回
// func scanFile(u *url.URL) (files []PanFile) {

// }

type PanFile struct {
	Name       string `json:"server_filename"`
	Size       int64  `json:"size"`
	Path       string `json:"path"`
	Parent     string `json:"parent"`
	IsDir      bool
	ReferTitle string
	ReferUrl   *url.URL
}
