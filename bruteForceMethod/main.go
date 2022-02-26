package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
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
	//output variable
	count := 0
	result := make(chan Result)
	done := make(chan struct{})
	//brute force logic
	for {
		wg := &sync.WaitGroup{}
		for i := range puzzleInput {
			for j := range puzzleInput[i] {
				if j+len(feed) <= MaxXBound {
					wg.Add(1)
					go matchKeywordWithPuzzle(puzzleInput, feed, result, i, j, true, wg)
				}
				if i+len(feed) <= MaxYBound {
					wg.Add(1)
					go matchKeywordWithPuzzle(puzzleInput, feed, result, i, j, false, wg)
				}
			}
		}
		go func() {
			wg.Wait()
			done <- struct{}{}
		}()
		select {
		case op := <-result:
			op.count = count
			fmt.Println(op)
			return nil
		case <-done:
		}
		//result not found rotate the array
		count += 1
		//25 rotation completed the alphabet series a->z
		if count > 25 {
			//bruteforce completed got the same matrix again
			return errors.New("No Pattern matched")
		}
		rotatePuzzle(puzzleInput)
	}
}

//matchKeywordWithPuzzle matches keyword input with the puzzle array
func matchKeywordWithPuzzle(puzzleInput [][]string, feed []string, result chan Result, i, j int, horizontal bool, wg *sync.WaitGroup) {
	defer wg.Done()
	if horizontal {
		for x := range feed {
			if puzzleInput[i][j+x] != feed[x] {
				return
			}
		}
		result <- Result{x: j, y: i, horizontal: true}
	} else {
		for y := range feed {
			if puzzleInput[i+y][j] != feed[y] {
				return
			}
		}
		result <- Result{x: j, y: i, horizontal: false}
	}
}

//rotatePuzzle increments each character by one
func rotatePuzzle(puzzleInput [][]string) {
	for i := range puzzleInput {
		for j := range puzzleInput[i] {
			if AscciChar := []byte(puzzleInput[i][j])[0]; AscciChar < 97 && AscciChar > 122 {
				panic("Input Out Of Scope!!")
			} else if AscciChar == 122 {
				puzzleInput[i][j] = string([]byte{97})
			} else {
				puzzleInput[i][j] = string([]byte(puzzleInput[i][j])[0] + 1)
			}
		}
	}
}
