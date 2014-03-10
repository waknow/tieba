package lib

import (
	"fmt"
	"net/url"
)

type Post struct {
	Url   *url.URL
	Title string
	Links []*url.URL
}

func GetPostLink(u string) (urls []*url.URL) {
	scheme := "http"
	host := "tieba.baidu.com"

	s, err := HttpGet(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	links := PostLink.FindAllString(s, -1)
	for _, link := range links {
		u, _ := url.Parse(link)
		u.Scheme = scheme
		u.Host = host
		urls = append(urls, u)
	}
	return
}

func GetPost(u *url.URL) (post Post) {
	post.Url = u

	s, err := HttpGet(u.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	post.Title = Group(TitleReg.FindStringSubmatch(s)).Get(1)
	links := PostPanLinkReg.FindAllString(s, -1)
	for _, link := range links {
		link = PostMarkerReplacer.ReplaceAllString(link, "&")
		panUrl, _ := url.Parse(link)
		post.Links = append(post.Links, panUrl)
	}
	return
}
