package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"story"
)

func main() {
	jsonFilePath := flag.String("jsonFileName", "gopher.json", "na")
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
	h := story.NewStoryHandler(storyy)
	log.Fatal(http.ListenAndServe(":8080", h))
}
