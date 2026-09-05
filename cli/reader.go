package cli

import(
	"strconv"
	"strings"
	"os"
)

// Error checking pending
func readArguments() ([]int, error) {
	argsWithProg := os.Args
	csv := argsWithProg[1]
	idsStrings := strings.SplitN(csv, ",", -1)
	return stringToIntArr(idsStrings)
}

func stringToIntArr(words []string) ([]int, error) {
	intSlice := make([]int, len(words))
	for i, val := range words {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
