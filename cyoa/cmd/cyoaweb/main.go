package main

import (
	"cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "story line json file")
	flag.Parse()
	fmt.Printf("Using story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
