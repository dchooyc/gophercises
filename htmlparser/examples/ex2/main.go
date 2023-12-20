package main

import (
	"fmt"
	"htmlparser/link"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i <= 3; i++ {
		testNum := strconv.Itoa(i)
		htmlContent := readFile(testNum)
		r := strings.NewReader(htmlContent)
		fmt.Println("\nFinding links for test" + testNum + ".html")
		links, err := link.Parse(r)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", links)
	}
}

func readFile(testNum string) string {
	prefix, suffix := "test", ".html"
	file, err := os.Open(prefix + testNum + suffix)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return ""
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	return string(bytes)
}
