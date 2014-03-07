package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

const (
	u = "http://tieba.baidu.com/f?kw=%CE%DE%CB%F0%D2%F4%C0%D6"
)

type Post struct {
	Url   url.URL
	Title string
	Links []url.URL
}

type PanInfo struct {
	Url     url.URL
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

type Group []string

func (g Group) Get(n int) string {
	strs := []string(g)
	if len(strs) >= n+1 {
		return strs[n]
	}
	return ""
}

var titleReg *regexp.Regexp

func init() {
	titleReg = regexp.MustCompile("<title>(.*)</title>")
}

func httpGet(u string) (s string, err error) {
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	s = string(b)
	return
}

func GetPostLink(u string) (urls []url.URL) {
	scheme := "http"
	host := "tieba.baidu.com"

	reg := regexp.MustCompile("/p/(\\d+)")

	s, err := httpGet(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	links := reg.FindAllString(s, -1)
	for _, link := range links {
		u, _ := url.Parse(link)
		u.Scheme = scheme
		u.Host = host
		urls = append(urls, *u)
	}
	return
}

func GetPost(u url.URL) (post Post) {
	post.Url = u

	panReg := regexp.MustCompile("http://(pan|yun).baidu.com/[0-9a-zA-Z/?&=;]+")
	replacerReg := regexp.MustCompile("&amp;")

	s, err := httpGet(u.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	post.Title = Group(titleReg.FindStringSubmatch(s)).Get(1)
	links := panReg.FindAllString(s, -1)
	for _, link := range links {
		link = replacerReg.ReplaceAllString(link, "&")
		u, _ := url.Parse(link)
		post.Links = append(post.Links, *u)
	}
	return
}

func GetPanInfo(p Post) (infos []PanInfo) {
	nameReg := regexp.MustCompile("<h2.*title=\"(.*?)\"")
	ukReg := regexp.MustCompile("share_uk=\"(\\d+)\"")
	idReg := regexp.MustCompile("share_id=\"(\\d+)\"")

	// raw string: \\/\\u6211\\u7684\\u97f3\\u4e50\\/\\u65e0\\u635f\\u97f3\\u4e50
	pathReg := regexp.MustCompile(`\\"path\\":\\"(.*?)\\"`)
	slashReplacer := regexp.MustCompile(`\\\\/`)  //for \\/
	slashuReplacer := regexp.MustCompile(`\\\\u`) //for \\u

	for _, link := range p.Links {
		fmt.Println(link.String())
		s, err := httpGet(link.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		name := Group(nameReg.FindStringSubmatch(s)).Get(1)
		title := Group(titleReg.FindStringSubmatch(s)).Get(1)
		uk := Group(ukReg.FindStringSubmatch(s)).Get(1)
		id := Group(idReg.FindStringSubmatch(s)).Get(1)

		path := Group(pathReg.FindStringSubmatch(s)).Get(1)
		path = slashReplacer.ReplaceAllString(path, "/")
		path = slashuReplacer.ReplaceAllString(path, "\\u")
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

func main() {
	fmt.Println("geting post links...")
	postLinks := GetPostLink(u)
	fmt.Printf("got %d links.\n", len(postLinks))

	fmt.Println("getting posts...")
	var posts []Post
	for _, link := range postLinks {
		fmt.Println(">", link.String())
		posts = append(posts, GetPost(link))
	}
	fmt.Printf("got %d posts.\n", len(posts))

	fmt.Println("getting pan info...")
	var infos []PanInfo
	for _, post := range posts {
		// fmt.Println(">", post.Url.String())
		infos = append(infos, GetPanInfo(post)...)
	}

	for _, info := range infos {
		fmt.Println(info.String())
	}
	// post := Post{}
	// u, _ := url.Parse("http://pan.baidu.com/s/1gdzIa5d")
	// post.Links = append(post.Links, *u)
	// GetPanInfo(post)
}
