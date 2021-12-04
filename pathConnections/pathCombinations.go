package pathConnections

import "math"

// we send in the all possible paths without start and end and look for unique, non-overlapping combinations
func PathCombinations(removedStartEndPaths [][]int) ([][]int, [][]int) {
	var combinationIndex []int
	var combinationRooms []int
	var multicomboIndexes [][]int
	var multicomboRooms [][]int
	// we have two loops over all paths
	for i := range removedStartEndPaths {
		combinationRooms = nil
		combinationIndex = nil
		combinationRooms = append(combinationRooms, removedStartEndPaths[i]...)
		combinationIndex = append(combinationIndex, i)
		multicomboIndexes = append(multicomboIndexes, combinationIndex)
		multicomboRooms = append(multicomboRooms, combinationRooms)

		for j := range removedStartEndPaths {
			// we check that they are not the same paths
			if i != j {
				if ComparePaths(combinationRooms, removedStartEndPaths[j]) {
					combinationRooms = append(combinationRooms, removedStartEndPaths[j]...)
					combinationIndex = append(combinationIndex, j)
				}
			}
		}
		multicomboIndexes = append(multicomboIndexes, combinationIndex)
		multicomboRooms = append(multicomboRooms, combinationRooms)
	}
	return multicomboIndexes, multicomboRooms
}

func ComparePaths(comb []int, add []int) bool {
	for _, room1 := range comb {
		for _, room2 := range add {
			if room1 == room2 {
				return false
			}
		}
	}
	return true
}
// here we find the fastest combination of all possible combnations depending on the number of ants
func FastestCombo(combinations [][]int, combinationIndexs [][]int, paths [][]int, ants int) [][]int {
	var fastestComb [][]int
	var timeslice []int
	var times float64
	var timeTemp int
	var time int
	var index int
	// we convert ants int to float64 to get the answer how fast each combination is depending on the number of ants
	antsF := float64(ants)
	for i, val := range combinations {
		valF := float64(len(val))
		for i2, j := range combinationIndexs {
			jF := float64(len(j))
			if i == i2 {
				times = (valF + antsF) / jF
				timeTemp = int(math.Round(times))
				timeslice = append(timeslice, timeTemp)
			}
		}
	}
	// from all the path times we sort through and find the path with shortest time
	for k, num := range timeslice {
		if k == 0 {
			time = num
			index = k
		}
		if num < time {
			time = num
			index = k
		}
	}
	fastestCombIn := combinationIndexs[index]
	var fastestPath []int
	for _, pathIndex := range fastestCombIn {
		fastestPath = append(fastestPath, paths[pathIndex]...)
		fastestComb = append(fastestComb, fastestPath)
		fastestPath = nil
	}
	return fastestComb
}

// we put back the end index to every path
func AddEnd(paths [][]int, end int) [][]int {
	for y := range paths {
		paths[y] = append(paths[y], end)

	}
	return paths
}

// we get the length of the fastest combination which helps us to decide which ant takes which path from the combination
func FastestComboLen(fastest [][]int) ([]int, [][]int) {
	var lenFastest []int
	var tempS []int
	var l int
	for _, val := range fastest {
		if len(val) < 2 {
			tempS = val
			fastest = nil
			fastest = append(fastest, tempS)
		}
	}
	for _, val := range fastest {
		l = len(val)
		lenFastest = append(lenFastest, l)
	}
	return lenFastest, fastest
}
