package utils

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetArtistInfo(url string, isSimilar bool) ([]string, []string) {

	c := colly.NewCollector()
	var artistInfo []string

	// Nickname
	c.OnHTML(".page-artist__title", func(h *colly.HTMLElement) {
		artistInfo = append(artistInfo, h.Text)
	})

	// Listnens and likes
	c.OnHTML(".artist-trends__total-count", func(h *colly.HTMLElement) {
		artistInfo = append(artistInfo, h.Text)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url + "/info")

	if !isSimilar {
		similarLinks := getSimilarArtistLinks(url+"/similar", c)
		return artistInfo, similarLinks
	}
	var blankLinks = []string{}
	return artistInfo, blankLinks
}
func getSimilarArtistLinks(url string, c *colly.Collector) []string {

	var similarLinks []string

	c.OnHTML(".artist__content .artist__name .d-artists a", func(h *colly.HTMLElement) {
		link := strings.TrimPrefix(h.Attr("href"), "/artist/")
		similarLinks = append(similarLinks, link)
	})

	c.Visit(url)
	return similarLinks
}
