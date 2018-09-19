package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	utf8Str := "你好，2018！"
	gbkStr, _ := ioutil.ReadAll(
		transform.NewReader(
			bytes.NewReader(
				[]byte(utf8Str),
			), simplifiedchinese.GBK.NewEncoder(),
		),
	)
	fmt.Printf("utf8:%v;len:%d\n", utf8Str, len(utf8Str))
	fmt.Printf("utf8 str:%s\n", utf8Str)
	fmt.Printf("gbk:%v;len:%d\n", gbkStr, len(gbkStr))
	fmt.Printf("gbk str:%s\n", gbkStr)
}
