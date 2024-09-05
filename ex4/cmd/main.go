package main

import (
	"flag"
	"fmt"
	parser "linkParser"
	"os"
)

func main() {
	htmlFlag := flag.String("htmlFile", "ex1.html", "desc")
	flag.Parse()
	reader, err := os.Open(*htmlFlag)
	if err != nil {
		fmt.Print("Somethign went wrong while opening html.")
		fmt.Print(err)
		return
	}
	links, err := parser.Parse(reader)
	if err != nil {
		fmt.Print("Somethign went wrong while parsing html.")
		return
	}
	for _, node := range links {
		fmt.Print(node)
	}
}
