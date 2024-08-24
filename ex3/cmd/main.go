package main

import (
	"flag"
	"fmt"
	"os"
	"story"
)

func main() {
	jsonFilePath := flag.String("jsonFileName", "gopher", "na")
	flag.Parse()
	file, err := os.Open(*jsonFilePath)
	if err != nil {
		panic(err)
	}
	storyy, err := story.DecodeToJson(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", storyy)
}
