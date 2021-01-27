package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("ikman.lk", "www.ikman.lk"),
	)

	c.OnHTML(".title--3s1R8", func(e *colly.HTMLElement) {
		title := e.Attr("class")
		fmt.Printf("Title Found: %q -> %s\n", e.Text, title)
		c.Visit(e.Request.AbsoluteURL(title))
	})
	c.OnHTML(".amount--3NTpl", func(e *colly.HTMLElement) {
		price := e.Attr("class")
		fmt.Printf("Price Found: %q -> %s\n", e.Text, price)
		c.Visit(e.Request.AbsoluteURL(price))
	})
	c.OnHTML(".word-break--2nyVq", func(e *colly.HTMLElement) {
		information := e.Attr("class")
		fmt.Printf("Data Found: %q -> %s\n", e.Text, information)
		c.Visit(e.Request.AbsoluteURL(information))
	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://ikman.lk/en/ad/suzuki-wagon-r-stingray-2019-for-sale-colombo")
}
