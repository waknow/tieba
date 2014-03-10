package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(u string) (s string, err error) {
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
