package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	target_url := "https://www.youtube.com/watch?v=wh0qATKFpOo"

	// req, err := http.NewRequest("GET", target_url, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// doc, err := goquery.NewDocumentFromReader(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(doc.Html())

	// log.Println(exists)

	// s := doc.Find("#secondary")

	// fmt.Printf("lines: %d\n", s.Size())
	// h, err := goquery.OuterHtml(s)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(h)

	// defer resp.Body.Close()
	// if resp.StatusCode != 200 {
	// 	log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	// }

	outside, err := goquery.NewDocument(target_url)
	if err != nil {
		log.Fatal(err)
	}

	innerSeceltion := outside.Find("script")
	innerSeceltion.Each(func(index int, s *goquery.Selection) {
		// iframe_url, _ := s.Attr("src")
		// log.Println(iframe_url)
		// log.Println(s.Html())
		t, err := s.Html()
		if err != nil {
			log.Fatal(err)
		}
		reg := "\r\n|\n"
		arr1 := regexp.MustCompile(reg).Split(t, -1)
		// log.Println(len(arr1))
		// fmt.Printf("%s\n", arr1[1])

		// fmt.Println(len(arr1))
		for _, s := range arr1 {
			if strings.Contains(s, "ytInitialData") {
				fmt.Printf("%s\n", s[0:100])
				fmt.Println("-----------------")
			}
		}

	})

}
