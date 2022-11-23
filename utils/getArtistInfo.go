package utils

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetArtistInfo(url string, link string, chann chan []string) {

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
	// Total Likes
	c.OnHTML(".d-generic-page-head__main-actions .d-like .d-button .d-button-inner .d-button__label", func(h *colly.HTMLElement) {
		artistInfo = append(artistInfo, h.Text)
	})

	artistInfo = append(artistInfo, link)
	lastRel, lastRelDate := getLastRelease(link)
	artistInfo = append(artistInfo, lastRel, lastRelDate)

	if err := c.Visit(url + "/info"); err != nil {
		log.Default().Print(err)
	}

	chann <- artistInfo
}

func getLastRelease(id string) (string, string) {
	lastRelease := ""
	lastReleaseDate := ""

	c := colly.NewCollector()
	// Last release
	c.OnHTML(".page-artist__latest-container .page-artist__latest-side .page-artist__subhead .d-subhead a .d-subhead__title .d-subhead__title-main span", func(h *colly.HTMLElement) {
		lastRelease = h.Text
	})
	// Rel date
	c.OnHTML(".page-artist__latest-album .album__bottom-right .album__year", func(h *colly.HTMLElement) {
		lastReleaseDate = h.Text
	})
	c.Visit("https://music.yandex.ru/artist/" + id)
	return lastRelease, lastReleaseDate
}

func GetSimilarArtistLinks(url string, artistId string, artistLinks []string) []string {
	c := colly.NewCollector()

	var similarLinks []string

	c.OnHTML(".artist__content .artist__name .d-artists a", func(h *colly.HTMLElement) {
		link := strings.TrimPrefix(h.Attr("href"), "/artist/")
		include := includes(artistLinks, link)
		if !include {
			similarLinks = append(similarLinks, link)
		}
	})

	c.Visit(url + "/similar")
	return similarLinks
}

func includes(arr []string, el string) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}
	return false
}
