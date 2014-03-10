package lib

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func PathToUnicode(path string) (s string) {
	unicodeReg := regexp.MustCompile(`\\u.{4}`)

	s = path
	codes := unicodeReg.FindAllString(path, -1)
	for _, item := range codes {
		res, err := toUnicode(item)
		if err != nil {
			log.Println(err)
			return ""
		}
		s = strings.Replace(s, item, res, -1)
	}
	return s
}

func toUnicode(s string) (res string, err error) {
	strs := strings.Split(s, "\\u")
	var i int64
	for _, str := range strs {
		if str == "" {
			continue
		}

		i, err = strconv.ParseInt(str, 16, 64)
		if err != nil {
			return
		}
		res += fmt.Sprintf("%c", i)
	}
	return
}
