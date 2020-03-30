package main

import (
	"fmt"
	//go语言版本的jquery
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var wg sync.WaitGroup

func main() {
	strUrl := "https://www.17k.com"

	doc, err := goquery.NewDocument(strUrl)
	if err != nil {
		fmt.Errorf("connect error:%#v", err)
		os.Exit(-1)
	}

	doc.Find(".list").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Find("a").Attr("href")
		fmt.Printf("开始下载影集图片:%v\n", src)
		if exists {
			wg.Add(1)
			go func(src string) {
				defer wg.Done()

				fmt.Printf("成功下载图片到:111\n")
			}(src)
		}
	})

	wg.Wait()
}


// test