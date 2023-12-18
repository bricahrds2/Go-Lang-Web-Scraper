package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Data struct {
	Title     string
	OrigPrice string
	DescPrice string
	Img       string
}

func main() {

	data := Data{}

	dataPage := make([]Data, 0, 1)

	c := colly.NewCollector(
		colly.AllowedDomains("invasion.club"),
	)

	c.OnRequest(func(h *colly.Request) {
		h.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Printf("Visiting %s\n", h.URL)
	})

	c.OnError(func(h *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	c.OnHTML("title", func(p *colly.HTMLElement) {
		fmt.Println(p.Text)
		//data.Title = p.Text
	})

	c.OnHTML("div.inner", func(p *colly.HTMLElement) {
		selection := p.DOM

		img := p.ChildAttr("img", "src")
		title := selection.Find("div.collection-product-title div.title").Text()
		oprice := p.ChildText("div.striped")
		dprice := clearDesc(selection.Find("div.collection-product-price div.price").Last().Text())

		//fmt.Printf("%s: %s\n", title, dprice)
		data.Title = title
		data.DescPrice = dprice
		data.Img = img
		data.OrigPrice = oprice

		dataPage = append(dataPage, data)
		data = Data{}

	})
	c.OnResponse(func(h *colly.Response) {
		fmt.Println(h.StatusCode)
	})

	//c.OnScraped(func(h *colly.Response) {
	//	dataPage = append(dataPage, data)
	//data = Data{}
	//})

	c.OnHTML("div.pagination", func(h *colly.HTMLElement) {

		next_page := h.Request.AbsoluteURL(h.ChildAttr("a", "href"))
		//next := next_page.Request.AbsoluteURL((h.Attr("href")))
		//next_page := h.Request.AbsoluteURL((h.Attr("href"))
		//fmt.Println(h.Text)
		c.Visit(next_page)
	})

	c.Visit("https://invasion.club/collections/tops")

	//Prints the JSON File in the Terminal

	//enc := json.NewEncoder(os.Stdout)
	//enc.SetIndent("", " ")
	//enc.Encode(dataPage)

	content, err := json.Marshal(dataPage)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("page-products.json", content, 0644)
}

func clearDesc(s string) string {
	return strings.TrimSpace(s)
}
