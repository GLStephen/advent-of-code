package main

import (
	"fmt"
	"sync"
)

type inputValues struct {
	start      string
	directions []string
	nodes      map[string]inputNode
}

type inputNode struct {
	left  string
	right string
}

type inputValuesParallel struct {
	start      []string
	directions []string
	nodes      map[string]inputNode
}

var target string = "ZZZ"
var parallelTarget string = "Z"

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

type foundNode struct {
	steps       int
	nextNodeKey string
}

func solveParalell(input inputValuesParallel) int {
	fmt.Printf("Input: %+v\n", input)

	var wg sync.WaitGroup

	steps := 0
	found := false
	for !found {
		for i := 0; i < len(input.start); i++ {
			wg.Add(1)
			fmt.Printf("=============== Running: %s ================== \n", input.start[i])
			go func(ele string) {
				defer wg.Done()
				// TODO: Figure out return values from waitgroups - WORTH Working on Beyond Competition
				//foundNode := runStep(input, ele)
				runStep(input, ele)
			}(input.start[i])
		}
		wg.Wait()

		found = true
	}

	fmt.Printf("Target: %s Reached In: %d Steps\n", target, steps)
	return steps
}

func runStep(input inputValuesParallel, nextNodeKey string) foundNode {
	// for key, nodes := range input.nodes {
	// 	fmt.Printf("Key -%+v- Node %+v\n", key, nodes)
	// }

	steps := 0
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
				return foundNode{}
			}
			steps++
			fmt.Printf("Step: %d Current Node Key: -%s- Current Node: %+v Direction: -%s- Next Node Key: -%s-\n", steps, currentNodeKey, currentNode, direction, nextNodeKey)
			// trimmedItemKey := strings.Trim(linePieces[0], " ")
			// if string(trimmedItemKey[len(trimmedItemKey)-1:][0]) == "A" {
			// 	start = append(start, trimmedItemKey)
			// }
			if string(nextNodeKey[len(nextNodeKey)-1:][0]) == parallelTarget {
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
		// if iterations > 1 {
		// 	break
		// }
	}

	return foundNode{steps: steps, nextNodeKey: nextNodeKey}
}
