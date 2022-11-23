package main

import (
	"log"
	"time"

	"github.com/sklvv/yandex-music-parser/input"
	"github.com/sklvv/yandex-music-parser/utils"
)

func main() {
	logger := log.Default()
	resultFile := [][]string{{"Yadex Music ID", "Last realease", "Last realease date", "Artist name", "Listens amount", "Likes amount", "Total Likes"}}
	aristsLinks := input.Links

	logger.Printf("Program started! Have %v initial links.", len(input.Links))
	logger.Print("Getting similar artists links. Please wait!")

	for _, link := range input.Links {
		similarLinks := utils.GetSimilarArtistLinks(input.BaseURL+link, link, aristsLinks)
		aristsLinks = append(aristsLinks, similarLinks...)
	}

	logger.Printf("Got similar artists! Total %v links.", len(aristsLinks))

	artistsChunks := utils.SplitChunks(aristsLinks)
	res := make([][]string, len(aristsLinks))
	logger.Printf("Splited to %v chunks, with each around %v links!", len(artistsChunks), len(artistsChunks[0]))

	for index, arChunk := range artistsChunks {
		logger.Printf("Getting info! %v of %v chunks.", index+1, len(artistsChunks))

		arChann := make(chan []string, len(arChunk))

		for _, link := range arChunk {
			go utils.GetArtistInfo(input.BaseURL+link, link, arChann)
			time.After(time.Millisecond * 500)
		}

		res := make([][]string, len(arChunk))

		for ind := range res {
			res[ind] = <-arChann
		}

		resultFile = append(resultFile, res...)
	}

	resultFile = append(resultFile, res...)

	logger.Print("Result is ready! Check input.csv")

	utils.WriteFile(resultFile)
}
