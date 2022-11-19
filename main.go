package main

import (
	"github.com/sklvv/yandex-music-parser/data"
	"github.com/sklvv/yandex-music-parser/utils"
)

func main() {
	resultFile := [][]string{{"Artist name", "Listens amount", "Likes amount", "Yadex Music ID"}}
	similarArtistLinks := []string{}
	for _, link := range data.Links {
		artistInfo, similarLinks := utils.GetArtistInfo(data.BaseURL+link, false)
		artistInfo = append(artistInfo, link)
		resultFile = append(resultFile, artistInfo)
		similarArtistLinks = append(similarArtistLinks, similarLinks...)
	}
	for _, simLink := range similarArtistLinks {
		artistInfo, _ := utils.GetArtistInfo(data.BaseURL+simLink, true)
		artistInfo = append(artistInfo, simLink)
		resultFile = append(resultFile, artistInfo)

	}
	utils.WriteFile(resultFile)
}
