package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	blogTitles, err := GetLatestBlogTitles("https://erp.cgc.ac.in/Account/Login")

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Blog Titles")
	fmt.Println(blogTitles)

}

func GetLatestBlogTitles(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return "", err
	}

	titles := ""

	doc.Find(".row").Each(func(i int, s *goquery.Selection) {
		titles += "-" + s.Text() + "\n"
	})

	return titles, nil
}
