package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// This Method is used to Generate the run time error.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// This Method is used the generate the random number for every new run
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var numberOfRandomWords int

	numberOfArgs := os.Args
	if len(numberOfArgs) == 1 {
		fmt.Println("Passing Default Value")
		numberOfRandomWords = 1
	} else {
		numberOfRandomWords, _ = strconv.Atoi(numberOfArgs[1])
	}

	var wordsArray []string
	//Opens the file
	file, err := os.Open("words_alpha.txt")
	check(err)
	defer file.Close()

	//Reads the entire file and returns the bytes and creating a array of all the words
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordsArray = append(wordsArray, scanner.Text())
	}

	//Generate Random Number
	randonNumber := RangeInt(0, len(wordsArray), numberOfRandomWords)
	if err := scanner.Err(); err != nil {
		check(err)
	}

	for _, index := range randonNumber {
		fmt.Println(wordsArray[index])
	}
}

// Generating random method
func RangeInt(min int, max int, n int) []int {
	arr := make([]int, n)
	var r int
	for r = 0; r <= n-1; r++ {
		arr[r] = rand.Intn(max) + min
	}
	return arr
}
