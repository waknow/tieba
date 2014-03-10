package lib

import (
	"fmt"
	"testing"
)

func TestPathToUnicode(t *testing.T) {
	path := `/abc\u6211\u7684\u97f3\u4e50/\u6211\u7684\u97f3\u4e50123`
	res := PathToUnicode(path)
	if res == "/abc我的音乐/我的音乐123" {
		t.Log("PathToUnicode ok.")
	} else {
		t.Error(fmt.Sprintf("want:%s,got:%s", "/abc我的音乐/我的音乐123", res))
	}
}
