package preparation

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//we read the text file, return it as []string
func ReadFile() []string {

	if len(os.Args) < 1 {
		fmt.Println("Missing parameter, provide file name!")
		os.Exit(1)
	}
	// we get the whole text as []byte
	text, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		os.Exit(1)
	}
	
	// we get the whole text as []string each string is made of a separate line from txt file
	textslice := strings.Split(string(text), "\n")
	return textslice
}

// here we read the first line (first string) of text and return nr of ants
func NumAnts(textslice []string) int {
	var antNr int
	var err error
	for i := range textslice {
		if i == 0 {
			// convert the string to an int
			antNr, err = strconv.Atoi(string(textslice[i]))
			if antNr < 1 || err != nil {
				fmt.Println("ERROR: We need ants!")
				os.Exit(1)
				panic(err)

			}

		}
	}
	return antNr
}
