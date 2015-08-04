package parse

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

func BuildShootings(body io.Reader) []Shooting {
	contents := TextAreaContents(body)
	shootings, err := ShootingsFromTextArea(contents)
	if err != nil {
		log.Fatal(err)
	}
	return shootings
}

func TextAreaContents(body io.Reader) string {
	doc, err := html.Parse(body)
	if err != nil {
		log.Fatalf("HTML token parse failure: %s", err)
	}
	return SelectTextArea(doc)
}

func SelectTextArea(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "textarea" {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "wpTextbox1" {
				return n.FirstChild.Data
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content := SelectTextArea(c)
		if content != "skip" {
			return content
		}
	}
	return "skip"
}
