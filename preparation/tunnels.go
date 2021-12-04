package preparation

import (
	"regexp"
	"sort"
	"strings"
)

// with TunnelIndex (tunnel pairs) we get indexes of corrisponding rooms [(richard dinish) becomes (0 3)]
func TunnelIndex(tunnel []string, rooms []string, roomindexes []int) []int {

	var connections []int

	// we loop over the indexes of name pairs of a tunnel
	for s := range tunnel {
		// we loop over the indexes of rooms
		for t := range rooms {
			if tunnel[s] == rooms[t] {
				// we append indexes of name pairs
				connections = append(connections, roomindexes[t])
			}
		}
	}
	return connections
}

// we range through the whole text, find names of rooms, roomindexes, names of tunnelpairs
func FindRoomsTunnels(textslice []string) ([][]int, int, int, []string) {
	var roomindexslice []int
	var roomconnections = make([][]int, 0)
	var rooms []string
	var startRoom int
	var endRoom int

	// we range over the the whole text
	for i := range textslice {
		// since textslice[0] is ants, we skip it
		if i > 0 {
			// we set the pattern that matches the lines where are room names
			roomPattern := regexp.MustCompile(`^[a-zA-Z0-9]+\s[0-9]+\s[0-9]+$`)
			// if string matches with lines with room names
			if roomPattern.MatchString(textslice[i]) {
				// we append an empty slice to a 2-dim slice to put in each room's connections
				roomconnections = append(roomconnections, []int{})
				// we save the string with room name (richard 0 6)
				roomline := textslice[i]
				// we save the whole string with room name as [] string ([richard 0 6])
				roomlineSlice := strings.Fields(roomline)
				// room is index 0 of that slice
				room := roomlineSlice[0]
				// we save each room in []string
				rooms = append(rooms, room)
				// all rooms as indexes in []int
				roomindexslice = append(roomindexslice, len(rooms)-1)
			}
			// we find the index of start rooom
			if textslice[i] == "##start" {
				startRoom = len(rooms)
			}
			// we find the index of end room
			if textslice[i] == "##end" {
				endRoom = len(rooms)
			}
			// we set the pattern that matches the lines where are room names
			tunnelPattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9]+$`)
			// if string matches with lines with room names
			if tunnelPattern.MatchString(textslice[i]) {
				tunnelline := textslice[i]

				// tunnelslice []string has all tunnelpairs without "-" come through in a loop, ([richard dinish])
				tunnelslice := strings.Split(tunnelline, "-")

				// with tunnelIndex func we change tunnelpair names to indexes (that corrispond to room indexes)
				// save them in []int
				tunnelIndexes := TunnelIndex(tunnelslice, rooms, roomindexslice)
				// with changing index pairs we append to a corresponding room's connection the other index from the index pair
				roomconnections[tunnelIndexes[0]] = append(roomconnections[tunnelIndexes[0]], tunnelIndexes[1])
				roomconnections[tunnelIndexes[1]] = append(roomconnections[tunnelIndexes[1]], tunnelIndexes[0])
			}
		}
	}
	return roomconnections, startRoom, endRoom, rooms
}

var startcounter int
var pathslice []int
var allPaths [][]int

// we send in each room's connections and get back all possible paths. It's a recursive function.
func AllPossiblePaths(start int, changingroom int, end int, roomconnections [][]int) [][]int {

	// we enter the first condition only the first time
	if startcounter == 0 {
		pathslice = append(pathslice, start)
		// first we range through the connections of the start room
		changingroom = start
		startcounter++
	}
	for _, i := range roomconnections[changingroom] {
		// with the boolean we make sure the path does not go back to where it has been
		var containsI bool
		for t := 0; t < len(pathslice); t++ {
			if pathslice[t] == i {
				containsI = true
				break
			}
		}
		if containsI {
			continue
		}
		// we append the room from roomconnections to the path
		pathslice = append(pathslice, i)
		// if we find the end, we save the path in allPaths
		if i == end {
			temp := make([]int, len(pathslice))
			copy(temp, pathslice)
			allPaths = append(allPaths, temp)
			pathslice = pathslice[:len(pathslice)-1]
			continue
		}
		// here we change the room for the next recursion
		changingroom = pathslice[len(pathslice)-1]
		AllPossiblePaths(start, changingroom, end, roomconnections)
		pathslice = pathslice[:len(pathslice)-1]
	}
	return allPaths
}

// we sort the paths, shortest comes first
func SortByLength(allpaths [][]int) [][]int {

	sort.Slice(allpaths, func(i, j int) bool {
		return len(allpaths[i]) < len(allpaths[j])
	})
	return allpaths
}

// we remove the first and last index of each path
func RemoveStartEnd(orderedPaths [][]int) [][]int {
	var editedPaths [][]int
	for _, val := range orderedPaths {
		val = val[1 : len(val)-1]
		editedPaths = append(editedPaths, val)
	}
	return editedPaths
}
