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
	fmt.Printf("Input: %+v\n", input.directions)

	// for key, nodes := range input.nodes {
	// 	fmt.Printf("Key -%+v- Node %+v\n", key, nodes)
	// }

	steps := 0
	nextNodeKey := input.start
	currentNode := input.nodes[nextNodeKey]
	found := false
	iterations := 0
	for !found {
		directionsIndex := 0
		//fmt.Print("Still Not Found \n")
		for directionsIndex < len(input.directions) {
			currentNodeKey := nextNodeKey
			//fmt.Printf("Node: %+v\n", nextNode)
			direction := input.directions[directionsIndex]
			if direction == "L" {
				nextNodeKey = currentNode.left
				//fmt.Printf("left: %s\n", nextNodeKey)
			} else if direction == "R" {
				nextNodeKey = currentNode.right
				//fmt.Printf("right: %s\n", nextNodeKey)
			} else {
				//fmt.Printf("Unknown Direction: %s\n", direction)
				return 0
			}
			steps++
			fmt.Printf("Step: %d Current Node Key: -%s- Current Node: %+v Direction: -%s- Next Node Key: -%s-\n", steps, currentNodeKey, currentNode, direction, nextNodeKey)
			if nextNodeKey == target {
				fmt.Print("Found \n")
				found = true
				break
			}
			currentNode = input.nodes[nextNodeKey]
			// fmt.Printf("Moving: %s to %+v\n", direction, nextNode)
			// fmt.Printf("Input: %+v\n", input.nodes)
			directionsIndex++
		}
		iterations++
	}

	fmt.Printf("Target: %s Reached In: %d Steps\n", target, steps)
	return steps
}
