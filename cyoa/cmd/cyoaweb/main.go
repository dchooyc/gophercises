package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "port to start cyoa")
	filename := flag.String("file", "gopher.json", "story line json file")
	flag.Parse()
	fmt.Printf("Using story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	story, err := cyoa.ParseStory(f)
	if err != nil {
		fmt.Println("Error parsing story:", err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("starting server on port:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
