package main

import (
	// "flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// var htmlFile = flag.String("htmlFile", "", "link to the html file to be parsed")
func parseText(n *html.Node, text string) string {
	if n.Type == html.TextNode {
		text = text + strings.TrimSpace(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = parseText(c, text)
	}
	return text
}

func parseAnchors(n *html.Node, anchor_list *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		// fmt.Println("\n", n.Attr)
		var href string
		for _, i := range n.Attr {
			if i.Key == "href" {
				href = i.Val
			}
		}
		text := parseText(n, "")
		*anchor_list = append(*anchor_list, Link{Href: href, Text: text})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseAnchors(c, anchor_list)
	}
}
func main() {
	r, err := os.Open("temp.html")
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	var anchor_list []Link
	parseAnchors(doc, &anchor_list)
	fmt.Println(anchor_list)
}
