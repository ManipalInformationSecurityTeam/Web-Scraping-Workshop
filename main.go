package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type FlipkartProduct struct {
	Name  string
	Price string
}

func main() {
	var query string
	var pages int
	var list []FlipkartProduct

	fmt.Println("Enter query:")
	fmt.Scanln(&query)
	fmt.Println("Enter max pages to scrape:")
	fmt.Scanln(&pages)
	pages++

	for i := 1; i < pages; i++ {
		url := fmt.Sprintf("https://www.flipkart.com/search?q=%s&page=%d", query, i)
		c := colly.NewCollector(colly.AllowedDomains("www.flipkart.com"))
		// Setting user-agent as Flipkart will display Internal server error if it can't identify the device
		c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:125.0) Gecko/20100101 Firefox/125.0"
		// Squared listing with minimal feature display
		c.OnHTML("div.slAVV4", func(h *colly.HTMLElement) {
			productName := h.ChildAttr("a.wjcEIp", "title")
			price := h.ChildText("div.Nx9bqj")
			list = append(list, FlipkartProduct{Name: productName, Price: price})

		})
		// Long listing with more feature display
		c.OnHTML("div.tUxRFH", func(h *colly.HTMLElement) {
			productName := h.ChildText("div.KzDlHZ")
			price := h.ChildText("div.Nx9bqj._4b5DiR")
			list = append(list, FlipkartProduct{Name: productName, Price: price})

		})
		c.OnRequest(func(r *colly.Request) {
			fmt.Printf("Visiting URL %s\n", r.URL)
			FlipkartHeaders(r)
		})
		c.OnError(func(r *colly.Response, e error) {
			fmt.Println(e)
		})
		c.Visit(url)
	}
	fmt.Println(list)
}

func FlipkartHeaders(req *colly.Request) {
	req.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Headers.Add("Accept-Encoding", "deflate, br")
	req.Headers.Add("Accept-Language", "en-US,en;q=0.5")
	req.Headers.Add("Connection", "keep-alive")
	req.Headers.Add("DNT", "1")
	req.Headers.Add("Host", "www.flipkart.com")
	req.Headers.Add("Referer", "https://www.flipkart.com/")
	req.Headers.Add("Sec-Fetch-Dest", "document")
	req.Headers.Add("Sec-Fetch-Mode", "navigate")
	req.Headers.Add("Sec-Fetch-Site", "same-origin")
	req.Headers.Add("Sec-Fetch-User", "?1")
	req.Headers.Add("Sec-GPC", "1")
	req.Headers.Add("Upgrade-insecure-Requests", "1")
	req.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:125.0) Gecko/20100101 Firefox/125.0")
}
