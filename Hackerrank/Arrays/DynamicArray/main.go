package main

import "fmt"

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

/*
Problem Recap
1. Initialize Structures

    Create a 2D array, arr, containing n empty arrays.
        Example: If n = 2, initialize as arr = [[], []].
    Declare an integer lastAnswer and initialize it to 0.

2. Queries

There are two types of queries:
Query 1: 1 x y

    Compute the index idx:
    idx=(x⊕lastAnswer)%n
    idx=(x⊕lastAnswer)%n
        ^ is the XOR operator, and % is the modulo operator.
    Append the integer y to arr[idx].

Query 2: 2 x y

    Compute the index idx:
    idx=(x⊕lastAnswer)%n
    idx=(x⊕lastAnswer)%n

    Access the subarray arr[idx] and compute the value at:
    element=arr[idx][y%size(arr[idx])]
    element=arr[idx][y%size(arr[idx])]

    Set lastAnswer = element.

    Store the value of lastAnswer in the results array (answers).

Input Format

    First line: Two integers, n (size of arr) and q (number of queries).
    Next q lines: Each line contains a query in the form:
        1 x y
        2 x y

Output

    Return the results of all Query 2 operations as an array.

Constraints

    1 ≤ n, q ≤ 10^5
    0 ≤ x, y ≤ 10^9
    For Query 2, it is guaranteed that arr[idx] is non-empty.

Plan to Solve the Problem

    Parse Input:
        Read integers n and q.
        Parse the list of queries into a suitable data structure.

    Initialize Structures:
        Create an array of n empty arrays: arr = [[], [], ...].
        Set lastAnswer = 0.
        Create an empty results array: answers = [].

    Process Queries:
        For each query:
            Compute idx = (x ^ lastAnswer) % n.
            Handle based on query type:
                Query 1: Append y to arr[idx].
                Query 2: Compute lastAnswer and append it to answers.

    Return Results:
        Return the answers array as the output.
*/

func dynamicArray(n int32, queries [][]int32) []int32 {
	arr := make([][]int32, n)
	var lastAns int32 = 0
	result := []int32{}

	for _, query := range queries {
		t, x, y := query[0], query[1], query[2]
		idx := (x ^ lastAns) % n

		if t == 1 {
			arr[idx] = append(arr[idx], y)
		} else if t == 2 {
			lastAns = arr[idx][y%int32(len(arr[idx]))]
			result = append(result, lastAns)
		}
	}

	return result
}

func main() {
	// Directly define the input
	n := int32(2)
	queries := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

	// Call the function and get the results
	results := dynamicArray(n, queries)

	// Print the results
	for _, result := range results {
		fmt.Println(result)
	}
}
