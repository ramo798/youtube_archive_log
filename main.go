package main

import (
	"log"

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
	outside.Each(func(index int, s *goquery.Selection) {
		// なんかする
	})
	innerSeceltion := outside.Find("iframe")
	innerSeceltion.Each(func(index int, s *goquery.Selection) {
		iframe_url, _ := s.Attr("src")
		log.Println(iframe_url)
	})

}
