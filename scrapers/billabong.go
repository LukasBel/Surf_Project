package scrapers

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Wetsuit struct {
	Type  string
	Price string
}

func WebScraper() ([]Wetsuit, error) {
	wetsuits := []Wetsuit{}
	URL := "https://www.billabong.com/sale-mens-wetsuits/"

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting:", request.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response:", response.StatusCode)
	})

	c.OnHTML("div.product producttile", func(element *colly.HTMLElement) {
		w := Wetsuit{}
		w.Type = element.Text
	})

	err := c.Visit(URL)
	if err != nil {
		return nil, err
	}
	return wetsuits, nil
}
