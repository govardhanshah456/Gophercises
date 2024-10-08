package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	linkNodes := linkNodes(doc)
	var links []Link
	for _, node := range linkNodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var str string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str += buildText(c)
	}
	return str
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = buildText(n)
	ret.Text = strings.Join(strings.Fields(ret.Text), " ")
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
