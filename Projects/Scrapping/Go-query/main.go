package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	blogTitles, err := GetLatestBlogTitles("https://www.friends2support.org/inner/news/listDonatedBloodUsers.aspx")

	if err != nil {
		log.Println(err)
	}

	file, err := os.Create("data.txt")

	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	_, err2 := file.WriteString(blogTitles)

	if err2 != nil {
		println(err)
	}

	fmt.Println("DATA Extracted")
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

	doc.Find("#dgLastBloodDonated").Each(func(i int, s *goquery.Selection) {
		// titles += "-" + s.Text() + "\n"
		s.Find("tbody").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
         fmt.Println(i);
         check :=  fmt.Sprintf("#dgLastBloodDonated_lblPlace_%d" , i);
				 titles += s.Find(check).Text();
			})
		})
	})

	return titles, nil
}
