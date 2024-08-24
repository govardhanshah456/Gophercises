package main

import (
	"encoding/json"
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
	var story story.Story
	r := json.NewDecoder(file)
	err = r.Decode(&story)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", story)
}
