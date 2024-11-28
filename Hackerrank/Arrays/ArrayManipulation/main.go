package main

import (
	"fmt"
)

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func arrayManipulation(n int32, queries [][]int32) int64 {
	res := make([]int32, n)
	for _, query := range queries {
		a, b, k := query[0], query[1], query[2]
		res[a-1] += k
		if b < n {
			res[b] -= k
		}
	}

	maxVal := int64(0)
	currValue := int64(0)

	for _, v := range res {
		currValue += int64(v)
		if currValue > maxVal {
			maxVal = currValue
		}
	}

	return maxVal
}

func main() {
	// Sample Input
	n := int32(10)
	queries := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}

	// Call the function
	result := arrayManipulation(n, queries)

	// Print the result
	fmt.Println("Maximum Value After Manipulation:", result)
}
