package link

import "io"

// Link represents a link in a html document
type Link struct {
	Href string
	Text string
}

// Parse takes in a html document and returns a
// slice of links
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
