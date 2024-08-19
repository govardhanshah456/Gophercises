package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv","problems.csv","hello")
	flag.Parse();
	file,err := os.Open(*csvFileName)
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
	for i := 0;i < totalLen;i++ {
		if len(data[i]) != 2 {
			fmt.Print("Invalid Data")
			continue;
		}
		question,answer := data[i][0],data[i][1]
		var input string
		fmt.Printf("%s=",question)
		fmt.Scanln(&input);
		if answer == input {
			match++
		}
		
	}
	fmt.Printf("You gave %d correct Answers \n",match)
}