package sendAnts

import (
	"fmt"
	"strconv"
	"strings"
)

// here we decide which is the optimal path for each ant and in which order they go
func SendAntsonTheirWay(ants int, rooms []string, fastestCombLen []int, bestComb [][]int) {
	var printOut [][]string
	var line int

	for ant := 1; ant <= ants; ant++ {
		i := ShortestIndex(fastestCombLen)
		// fastestCombLen[i] changes with every ant, here we decide on which order and line certain ant's moves are printed
		line = fastestCombLen[i] - len(bestComb[i])
		fastestCombLen[i] = fastestCombLen[i] + 1

		for w, val := range bestComb[i] {
			// this is the case when the start and end room are connected and there are no rooms between them
			if len(bestComb) < 2 && len(bestComb[i]) == 1 {

				fmt.Print("L" + strconv.Itoa(ant) + "-" + rooms[bestComb[i][0]] + " ")
		
				// otherwise we do the following
			} else {
				// if length of the printOut slice is bigger than the (index of room + line)
				if len(printOut) > w + line {

					printOut[w+line] = append(printOut[w+line], "L"+ strconv.Itoa(ant)+"-"+rooms[val])

				} else {
					printOut = append(printOut, []string{"L" + strconv.Itoa(ant) + "-" + rooms[val]})
				}
			}
		}
	}
	for _, text := range printOut {

		fmt.Println(strings.Join(text, " "))
	}
	fmt.Println()
}

// here we get the index of currently most optimal path in the fastest path combination
func ShortestIndex(fastestLen []int) int {
	minValue := fastestLen[0]
	index := 0
	for i := 1; i < len(fastestLen); i++ {
		if fastestLen[i] < minValue {
			minValue = fastestLen[i]
			index = i
		}
	}
	return index
}
