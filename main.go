package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var score int
var questionCount int

func main() {
	quizzPath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer (defaults to 'problems.csv'" )
	timeLimit := flag.Int("time", 30, "time limit to finish the quizz, in seconds")
	flag.Parse()

	fmt.Printf("You have %v seconds to finish the quizz. Press Enter when you're ready. ", *timeLimit)
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	go func()  {
		<-timer.C
		fmt.Println("\nTime out !")
		fmt.Printf("Your score is %v/%v\n", score, questionCount)
		os.Exit(0)
	}()

	quizzLines := readCSVFile(*quizzPath)
	questionCount = len(quizzLines)
	consoleReader := bufio.NewReader(os.Stdin)
	score = 0

	for _, line := range quizzLines {
		fmt.Print(line[0] + "=")
		input, err := consoleReader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if err != nil {
			log.Fatal(err)
		}

		if input == line[1] {
			score++
		}
	}

	fmt.Printf("Your score is %v/%v\n", score, questionCount)
}

func readCSVFile(path string) [][]string {
	quizzFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer quizzFile.Close()

	quizzFileReader := csv.NewReader(quizzFile)
	quizzLines, err := quizzFileReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return quizzLines
}