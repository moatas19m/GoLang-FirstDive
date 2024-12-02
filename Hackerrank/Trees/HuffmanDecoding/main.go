package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Node struct {
	data  rune
	left  *Node
	right *Node
}

// Implementing a Priority Queue for huffman nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].data < pq[j].data
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// encode creates a Huffman tree and encodes the string
func encode(input string) (*Node, string, map[rune]string) {
	// Step 1: Calculate frequencies
	frequency := make(map[rune]int)
	for _, char := range input {
		frequency[char]++
	}

	// Step 2: Build a priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)
	for char, freq := range frequency {
		heap.Push(pq, &Node{data: char, freq: freq})
	}

	// Step 3: Build the Huffman Tree
	for pq.Len() > 1 {
		left := heap.Pop(pq).(*Node)
		right := heap.Pop(pq).(*Node)

		// Create a new internal node with combined frequency
		internal := &Node{
			data:  '\x00',
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}
		heap.Push(pq, internal)
	}

	// Root of the Huffman Tree
	root := heap.Pop(pq).(*Node)

	// Step 4: Generate Huffman codes
	huffmanCodes := make(map[rune]string)
	generateCodes(root, "", huffmanCodes)

	// Step 5: Encode the input string
	encodedString := ""
	for _, char := range input {
		encodedString += huffmanCodes[char]
	}

	return root, encodedString, huffmanCodes
}

// generateCodes generates Huffman codes recursively
func generateCodes(node *Node, code string, huffmanCodes map[rune]string) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		huffmanCodes[node.data] = code
		return
	}
	generateCodes(node.left, code+"0", huffmanCodes)
	generateCodes(node.right, code+"1", huffmanCodes)
}

// decode decodes the encoded string using the Huffman tree
func decode(root *Node, encodedString string) string {
	decoded := ""
	current := root

	for _, bit := range encodedString {
		if bit == '0' {
			current = current.left
		} else if bit == '1' {
			current = current.right
		}

		// If a leaf node is reached, append the character
		if current.left == nil && current.right == nil {
			decoded += string(current.data)
			current = root
		}
	}

	return decoded
}

func main() {
	// Input string
	fmt.Println("Enter the string to encode:")
	var input string
	fmt.Scanln(&input)

	// Encode the input and generate the Huffman tree
	root, encodedString, huffmanCodes := encode(input)

	// Print the encoded string and Huffman codes
	fmt.Println("Encoded String:", encodedString)
	fmt.Println("Huffman Codes:")
	for char, code := range huffmanCodes {
		fmt.Printf("%c: %s\n", char, code)
	}

	// Decode the encoded string
	decodedString := decode(root, encodedString)
	fmt.Println("Decoded String:", decodedString)
}
