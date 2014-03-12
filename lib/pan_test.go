package lib

import (
	"fmt"
	"net/url"
	"testing"
)

const (
	RawUrl = "http://pan.baidu.com/s/1sj6kkeL"
)

func TestGetPanInfo(t *testing.T) {
	u, _ := url.Parse(RawUrl)
	info := GetPanInfo(u)
	if info.UK() == "" || info.ShareID() == "" {
		t.Error("can't get useful info", info.UK(), info.ShareID())
	} else {
		t.Log("ok")
	}
}

func TestGetFIleFromHomepage(t *testing.T) {
	u, _ := url.Parse(RawUrl)
	fs := getFileFromHomePage(u)
	if len(fs) == 0 {
		t.Error("can't get files from the home page!")
	} else {
		t.Log(fs)
	}
}

func TestGetFileFromRequest(t *testing.T) {
	u, _ := url.Parse(RawUrl)
	info := GetPanInfo(u)
	var path string
	if info.Path() != "/" {
		path = info.Path() + "/" + info.Name
	} else {
		path = "/" + info.Name
	}

	fs := getFileFromRequest(path, info.uk, info.shareId)
	if len(fs) == 0 {
		t.Error("can't get files from the request!")
	} else {
		for _, f := range fs {
			fmt.Println(f.Name)
		}
		t.Log(fs)
	}
}
