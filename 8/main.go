package main

import "fmt"

type inputValues struct {
	start      string
	directions []string
	nodes      map[string]inputNode
}

type inputNode struct {
	left  string
	right string
}

var target string = "ZZZ"

func main() {
	fmt.Printf("No Real Main Day 8")
	input := inputValues{}
	solve(input)
}

func solve(input inputValues) int {
	fmt.Printf("Input: %+v\n", input)

	for key, nodes := range input.nodes {
		fmt.Printf("Key -%+v- Node %+v\n", key, nodes)
	}

	steps := 1
	nextNodeKey := input.start
	nextNode := input.nodes[nextNodeKey]
	found := false
	iterations := 0
	for !found {
		fmt.Print("Still Not Found \n")
		for mapIndex := 0; mapIndex < len(input.directions); mapIndex++ {
			direction := input.directions[mapIndex]
			fmt.Printf("Moving: %s on %+v\n", direction, nextNode)
			if direction == "L" {
				nextNodeKey = nextNode.left
				fmt.Printf("left: %s\n", nextNodeKey)
			} else if direction == "R" {
				nextNodeKey = nextNode.right
				fmt.Printf("right: %s\n", nextNodeKey)
			} else {
				fmt.Printf("Unknown Direction: %s\n", direction)
				return 0
			}
			steps++
			if nextNodeKey == target {
				fmt.Print("Found")
				found = true
				break
			}
			nextNode := input.nodes[nextNodeKey]
			fmt.Printf("Input: %+v\n", input.nodes)
			fmt.Printf("Next Node Key: %s Node: %+v\n", nextNodeKey, nextNode)
		}
		iterations++
		if iterations > 5 {
			found = true
		}
	}

	fmt.Printf("Target: %s Reached In: %d Steps\n", target, steps)
	return steps
}
