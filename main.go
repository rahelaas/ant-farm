package main

import (
	"fmt"
	"lem-in/pathConnections"
	"lem-in/preparation"
	"lem-in/sendAnts"
	"os"
)

func main() {

	textSlice := preparation.ReadFile()

	ants := preparation.NumAnts(textSlice)

	roomConnections, start, end, rooms := preparation.FindRoomsTunnels(textSlice)

	allPaths := preparation.AllPossiblePaths(start, start, end, roomConnections)
	if len(allPaths) < 1 {
		fmt.Println("ERROR: No valid paths found!")
		os.Exit(1)
	}
	for _, string := range textSlice {
		fmt.Println(string)
	}
	fmt.Println()
	
	sortedPaths := preparation.SortByLength(allPaths)

	noStartEndpaths := preparation.RemoveStartEnd(sortedPaths)

	allCombinationIndexes, allCombinations := pathConnections.PathCombinations(noStartEndpaths)

	fastestComb := pathConnections.FastestCombo(allCombinations, allCombinationIndexes, noStartEndpaths, ants)

	fastestComb = pathConnections.AddEnd(fastestComb, end)

	fastestLen, newFastestComb := pathConnections.FastestComboLen(fastestComb)

	sendAnts.SendAntsonTheirWay(ants, rooms, fastestLen, newFastestComb)
}
