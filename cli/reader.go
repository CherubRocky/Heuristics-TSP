package cli

import(
	"strconv"
	"strings"
	"bufio"
	"os"
)

// Error checking pending
func readLine() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err 
	}
	line := scanner.Text()
	idsStrings := strings.SplitN(line, ",", -1)
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
