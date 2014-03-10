package main

import (
	"fmt"
	"net/url"

	"tieba/lib"
)

const (
	u = "http://tieba.baidu.com/f?kw=%CE%DE%CB%F0%D2%F4%C0%D6"
)

func main() {
	// fmt.Println("geting post links...")
	// postLinks := lib.GetpostLink(u)
	// fmt.Printf("got %d links.\n", len(postLinks))

	// fmt.Println("getting posts...")
	// var posts []lib.Post
	// for _, link := range postLinks {
	// 	fmt.Println(">", link.String())
	// 	posts = append(posts, lib.GetPost(link))
	// }
	// fmt.Printf("got %d posts.\n", len(posts))

	// fmt.Println("getting pan info...")
	// var infos []lib.PanInfo
	// for _, post := range posts {
	// 	// fmt.Println(">", post.Url.String())
	// 	infos = append(infos, lib.GetPanInfo(post)...)
	// }

	// for _, info := range infos {
	// 	fmt.Println(info.String())
	// }
	fmt.Println("\u6211")
	post := lib.Post{}
	u, _ := url.Parse("http://pan.baidu.com/s/1sjATQGl")
	post.Links = append(post.Links, u)
	info := lib.GetPanInfo(post)
	fmt.Println(info[0].String())
}
