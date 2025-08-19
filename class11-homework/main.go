package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var r = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
	  </body>
</html>
`

func countWordsAndImages(doc *html.Node) (words int, pics int) {

	for n := range doc.Descendants() {
		if n.Type == html.TextNode {
			words = words + len(strings.Fields(n.Data))
		}
		if n.Type == html.ElementNode {
			if n.Data == "img" {
				pics++
			}
		}

	}
	return words, pics
}

func main() {
	//io.Reader Provides a standard way to read from any data source Files, Network connections, HTTP responses, in-memory buffers (strings)
	//Reads data in chunks (via the Read(p []byte) method) instead of loading everything at once.
	doc, err := html.Parse(strings.NewReader(r))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)
	fmt.Printf("%d words and %d images \n", words, pics)

}
