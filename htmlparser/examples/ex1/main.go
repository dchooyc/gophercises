package main

import (
	"fmt"
	"htmlparser/link"
	"strings"
)

var exampleHtml = `
<html>
<body>
	<h1>Merry Christmas</h1>
	<a href="/other-page">Link over here</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
