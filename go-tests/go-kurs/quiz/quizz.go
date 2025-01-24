package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rightAnswer := 0
	wrongAnswer := 0
	start := time.Now()

	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	for _, eachrecord := range records {
		fmt.Println("Frage: ", eachrecord[0])
		answer := eachrecord[1]
		var userAnswer string

		fmt.Print("Antwort: ")
		fmt.Scan(&userAnswer)
		if answer == userAnswer {
			fmt.Println("Richtig")
			rightAnswer += 1
		} else {
			fmt.Println("Falsch")
			wrongAnswer += 1
		}
		dt := time.Now()
		fmt.Println("Abgelaufene Zeit ", dt.Sub(start))
	}
	fmt.Printf("Richtig = %d Falsch = %d", rightAnswer, wrongAnswer)
}
