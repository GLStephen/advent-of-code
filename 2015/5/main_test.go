package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGiven(t *testing.T) {
	nice, naughty, err := solve([]string{"ugknbfddgicrmopn"})

	assert.Nil(t, err)
	assert.Equal(t, 1, nice, "Expected 1 nice string")
	assert.Equal(t, 0, naughty, "Expected 0 naughty string")

	nice, naughty, err = solve([]string{"aaa"})

	assert.Nil(t, err)
	assert.Equal(t, 1, nice, "Expected 1 nice string")
	assert.Equal(t, 0, naughty, "Expected 0 naughty string")

	nice, naughty, err = solve([]string{"jchzalrnumimnmhp"})

	assert.Nil(t, err)
	assert.Equal(t, 0, nice, "Expected 0 nice string")
	assert.Equal(t, 1, naughty, "Expected 1 naughty string")

	nice, naughty, err = solve([]string{"haegwjzuvuyypxyu"})

	assert.Nil(t, err)
	assert.Equal(t, 0, nice, "Expected 0 nice string")
	assert.Equal(t, 1, naughty, "Expected 1 naughty string")

	nice, naughty, err = solve([]string{"dvszwmarrgswjxmb"})

	assert.Nil(t, err)
	assert.Equal(t, 0, nice, "Expected 0 nice string")
	assert.Equal(t, 1, naughty, "Expected 1 naughty string")

	nice, naughty, err = solve([]string{"ugknbfddgicrmopn", "aaa", "jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb"})

	assert.Nil(t, err)
	assert.Equal(t, 2, nice, "Expected 3 naughty string")
	assert.Equal(t, 3, naughty, "Expected 3 naughty string")
}

func TestSolve(t *testing.T) {
	testStrings, err := loadInput()
	if err != nil {
		assert.Fail(t, "Error loading input", err)
	}
	nice, naughty, err := solve(testStrings)

	assert.Nil(t, err)
	assert.Equal(t, 238, nice)
	assert.Equal(t, 762, naughty)
}

func TestBonusGiven(t *testing.T) {
	nice, naughty, err := solve([]string{"qjhvhtzxzqqjkmpb"})

	assert.Nil(t, err)
	assert.Equal(t, 1, nice, "Expected 1 nice string")
	assert.Equal(t, 0, naughty, "Expected 0 naughty string")

	// nice, naughty, err = solve([]string{"xxyxx"})

	// assert.Nil(t, err)
	// assert.Equal(t, 1, nice, "Expected 1 nice string")
	// assert.Equal(t, 0, naughty, "Expected 0 naughty string")

	// nice, naughty, err = solve([]string{"uurcxstgmygtbstg"})

	// assert.Nil(t, err)
	// assert.Equal(t, 0, nice, "Expected 0 nice string")
	// assert.Equal(t, 1, naughty, "Expected 1 naughty string")

	// nice, naughty, err = solve([]string{"ieodomkazucvgmuy"})

	// assert.Nil(t, err)
	// assert.Equal(t, 0, nice, "Expected 0 nice string")
	// assert.Equal(t, 1, naughty, "Expected 1 naughty string")

	// nice, naughty, err = solve([]string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"})

	// assert.Nil(t, err)
	// assert.Equal(t, 2, nice, "Expected 3 naughty string")
	// assert.Equal(t, 3, naughty, "Expected 3 naughty string")
}

func loadInput() ([]string, error) {
	results := make([]string, 0)
	file, err := os.Open("./input.txt")
	if err != nil {
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, line)
		//fmt.Printf("%d: %s\n", count, line)
		count++
	}

	return results, nil
}
