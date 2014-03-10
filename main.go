package main

import (
	"fmt"

	"tieba/page"
)

const (
	u = "http://tieba.baidu.com/f?kw=%CE%DE%CB%F0%D2%F4%C0%D6"
)

func main() {
	fmt.Println("geting post links...")
	postLinks := page.GetPostLink(u)
	fmt.Printf("got %d links.\n", len(postLinks))

	fmt.Println("getting posts...")
	var posts []page.Post
	for _, link := range postLinks {
		fmt.Println(">", link.String())
		posts = append(posts, page.GetPost(link))
	}
	fmt.Printf("got %d posts.\n", len(posts))

	fmt.Println("getting pan info...")
	var infos []page.PanInfo
	for _, post := range posts {
		// fmt.Println(">", post.Url.String())
		infos = append(infos, page.GetPanInfo(post)...)
	}

	for _, info := range infos {
		fmt.Println(info.String())
	}
	// post := Post{}
	// u, _ := url.Parse("http://pan.baidu.com/s/1gdzIa5d")
	// post.Links = append(post.Links, *u)
	// GetPanInfo(post)
}
