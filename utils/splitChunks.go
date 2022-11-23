package utils

func SplitChunks(initSlice []string) [][]string {
	chunksNum := len(initSlice) / 11
	chunkLen := len(initSlice) / chunksNum
	result := [][]string{}
	endOfchunk := 0
	for i := 0; i < chunksNum; i++ {
		endOfchunk += chunkLen
		if endOfchunk > len(initSlice) {
			endOfchunk = len(initSlice)
		}
		result = append(result, initSlice[i*chunkLen:endOfchunk])
	}
	return result
}
