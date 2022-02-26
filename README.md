# Findtheword

Assumptions:
  * All Input structures contains strings(collection of bytes) instead of a rune(unicode representation)
  * Only Ascii Small Alphabets characters (a-z) will be used in input
  * Slice are used as input parameters.But these are remain read only through the scope
  * Input keyword to be matched will not contain any trainling/leading spaces or any unwanted characters
  * Programs only prints the find occurrence of the input pattern. **NOT ALL**
  
 The Following Repistory contains two method to approach the same problem:
 
 * Brute Force Solution: The following solution leverages go routines to check each and every word pattern horizontally and vertically to acheive the required result.In worst scenario it will rotate the array 25 times to match the input pattern.
 * Difference Method: The following solution calculate a difference in rotations for each element of the 2-D array and then approaches each and every element by applying that delta.

## Execution

Clone the folowing git repository and change directory to bruteForceMethod or differenceMethod folder

```
go run main.go
```

By Default, sample input is present in main function(in main.go).This can be changed to try an other input.

## Testing


Clone the folowing git repository and change directory to bruteForceMethod or differenceMethod folder

```
go test -v
```

Four testcases are present in the test file.Two are Valid Scenarios(Netapp and Container Puzzles) and Two are Invalid Scenario(No Input to the program)

Sample Output Of test Files:
```
=== RUN   TestFindTheWord
=== RUN   TestFindTheWord/valid
x=5, y=15, count=8 found horizontally
=== RUN   TestFindTheWord/valid#01
x=10, y=2, count=17 found vertically
=== RUN   TestFindTheWord/invalid
    main_test.go:68: empty Input
=== RUN   TestFindTheWord/invalid#01
    main_test.go:68: empty Input
=== RUN   TestFindTheWord/invalid#02
    main_test.go:68: empty Input
--- PASS: TestFindTheWord (0.00s)
    --- PASS: TestFindTheWord/valid (0.00s)
    --- PASS: TestFindTheWord/valid#01 (0.00s)
    --- PASS: TestFindTheWord/invalid (0.00s)
    --- PASS: TestFindTheWord/invalid#01 (0.00s)
    --- PASS: TestFindTheWord/invalid#02 (0.00s)
PASS
ok      playground/findaword/differenceMethod   0.629s
```
