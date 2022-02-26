package main

import (
	"errors"
	"fmt"
	"strings"
)

//This structure is a placeholder for computed Result
type Result struct {
	x          int
	y          int
	count      int
	horizontal bool
}

//Stringer interface implementation for the output object
func (r Result) String() string {
	if r.horizontal {
		return fmt.Sprintf("x=%d, y=%d, count=%d found horizontally", r.x, r.y, r.count)
	} else {
		return fmt.Sprintf("x=%d, y=%d, count=%d found vertically", r.x, r.y, r.count)
	}
}

//The main function.Contains a sample snippet to execute code
func main() {
	puzzleInput := [][]string{
		{"z", "e", "z"},
		{"w", "n", "a"},
		{"n", "q", "b"},
	}
	keyword := "abc"
	err := FindTheWord(puzzleInput, keyword)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func FindTheWord(puzzleInput [][]string, keyword string) error {
	//check for empty data
	if keyword == "" || len(puzzleInput) == 0 || len(puzzleInput[0]) == 0 {
		return errors.New("empty Input")
	}
	//preparing input for computation
	feed := strings.Split(keyword, "")
	MaxXBound := len(puzzleInput)
	MaxYBound := len(puzzleInput[0])
	//computation logic
	for i := range puzzleInput {
		for j := range puzzleInput[i] {
			delta := calculateDelta(puzzleInput[i][j], feed[0])
			if j+len(feed) <= MaxXBound {
				result, match := matchDelta(puzzleInput, feed, i, j, true, delta)
				if match {
					fmt.Println(result)
					break
				}
			}
			if i+len(feed) <= MaxYBound {
				result, match := matchDelta(puzzleInput, feed, i, j, false, delta)
				if match {
					fmt.Println(result)
					break
				}
			}
		}
	}
	return nil
}

//matchDelta matches keyword input with the puzzle array by incrementing delta for each character
func matchDelta(puzzleInput [][]string, feed []string, i, j int, horizontal bool, delta int) (Result, bool) {
	if horizontal {
		for x := range feed {
			if rotate(puzzleInput[i][j+x], delta) != feed[x] {
				return Result{}, false
			}
		}
		return Result{x: j, y: i, horizontal: true, count: delta}, true
	} else {
		for y := range feed {
			if rotate(puzzleInput[i+y][j], delta) != feed[y] {
				return Result{}, false
			}
		}
		return Result{x: j, y: i, horizontal: false, count: delta}, true
	}
}

//calculateDelta calculates the number of rotations pending for character(in input array) to match the first character of input keyword
func calculateDelta(param1, param2 string) int {
	AscciChar1 := byte(param1[0])
	AscciChar2 := byte(param2[0])
	if (AscciChar1 < 97 && AscciChar1 > 122) || (AscciChar2 < 97 && AscciChar2 > 122) {
		panic("Input Out Of Scope!!")
	} else if AscciChar1 <= AscciChar2 {
		return int(AscciChar2) - int(AscciChar1)
	} else {
		return 25 + (int(AscciChar2) - int(AscciChar1)) + 1
	}
}

//rotate increments input character by "X" times
func rotate(input string, times int) string {
	if AscciChar := []byte(input)[0]; (int(AscciChar)+times)/122 == 0 {
		input = string([]byte{AscciChar + byte(times)})
	} else {
		input = string([]byte{96 + byte((int(AscciChar)+times)%122)})
	}
	return input
}
