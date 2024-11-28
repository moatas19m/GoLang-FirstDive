package main

import "fmt"

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY stringList
 *  2. STRING_ARRAY queries

 */

func matchingStrings(stringList []string, queries []string) []int32 {
	strFrequency := make(map[string]int32)
	res := make([]int32, len(queries))

	for _, str := range stringList {
		strFrequency[str]++
	}

	for i, query := range queries {
		res[i] = strFrequency[query]
	}

	return res
}

func main() {
	stringList := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}

	// Call the function
	results := matchingStrings(stringList, queries)

	// Print the results
	fmt.Println(results) // Output: [2, 1, 0]
}
