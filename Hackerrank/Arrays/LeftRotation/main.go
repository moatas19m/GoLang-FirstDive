package main

import (
	"fmt"
)

/*
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

func rotateLeft(d int32, arr []int32) []int32 {
	n := int32(len(arr))
	d = d % n // When d is greater than or equal to n, ew only need the remainder, i.e, the number of steps to actually rotate. For example if d==n, the remainder is 0 and the rotated array is the array from the input0
	res := make([]int32, len(arr))
	// Idea: slice the array into 2 parts, then swap the positions of the two parts

	res = append(arr[d:], arr[:d]...)
	return res
}

func main() {
	// Input values
	d := 10
	arr := []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51}

	// Print original array
	fmt.Println("Original Array:", arr)

	// Perform left rotation
	result := rotateLeft(int32(d), arr)

	// Print the rotated array
	fmt.Println("Rotated Array:", result)
}
