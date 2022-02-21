package db

import (
	"net/http"

	"golang.org/x/net/html"
)

// Parse returns a list of words from the specified URL
func ParseWordsFromUrl(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("get:\n", keepLines(string(body), 30))
	// fmt.Println(string(body))
	// fmt.Println(resp.StatusCode)

	nodes, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	words := []string{}
	dfs(nodes, &words)
	return words, nil
}

// searching for <a> tags, with the following pattern:
// <a class="word dynamictext" href="/dictionary/antenna">antenna</a>
// and then locate the text inside the tag
func dfs(n *html.Node, w *[]string) {
	// fmt.Println(n.Attr, n.Data, n.Type)
	if n.Type == html.ElementNode && n.Data == "a" {
		if n.Attr[0].Val == "word dynamictext" {
			*w = append(*w, n.LastChild.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, w)
	}
}
