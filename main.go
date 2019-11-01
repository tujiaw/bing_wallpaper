package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get("https://cn.bing.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	HOST := "https://cn.bing.com"
	wallpaperUrl := ""
	wallpaperTitle := ""
	doc.Find("body #bgImgProgLoad").Each(func(i int, s *goquery.Selection) {
		src, exist := s.Attr("data-ultra-definition-src")
		if exist {
			wallpaperUrl = HOST + src
		}
	})

	doc.Find("body #sh_cp").Each(func(i int, s *goquery.Selection) {
		wallpaperTitle, _ = s.Attr("title")
	})

	result := map[string]string{
		"title": wallpaperTitle,
		"url":   wallpaperUrl,
	}

	fmt.Println(result)
}
