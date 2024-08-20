package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "hello")
	timer := flag.Int64("timer", 30, "hello")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err1 := reader.ReadAll()
	if err1 != nil {
		panic(err1)
	}
	totalLen := len(data)
	match := 0
	timeLeft := time.NewTimer(time.Duration(*timer) * time.Second)
	fmt.Println("Press 'Enter' to start the quiz...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	done := make(chan bool)
	go func() {
		for i := 0; i < totalLen; i++ {
			if len(data[i]) != 2 {
				fmt.Print("Invalid Data")
				continue
			}
			question, answer := data[i][0], data[i][1]
			var input string
			fmt.Printf("%s=", question)
			fmt.Scanln(&input)
			if answer == input {
				match++
			}

		}
		done <- true
	}()
	select {
	case <-timeLeft.C:
		fmt.Printf("\n\nTime Up\n")
	case <-done:
		fmt.Println("\n\n\n\nQuiz over before time expired! Bravo")

	}
	fmt.Printf("You gave %d correct Answers \n", match)
}
