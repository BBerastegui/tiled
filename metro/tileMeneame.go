package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// Initialize the map of regex
var r map[string]*regexp.Regexp

func main() {

	r = make(map[string]*regexp.Regexp)

	r["news"], _ = regexp.Compile(`<div class="news-summary">(.*?)</span></span>\s+</div>\s+</div>\s+</div>`)
	r["title"], _ = regexp.Compile(`<h2>\s+<a .*>(.*?)\s+</a>\s+</h2>`)
	type News struct {
		Level   int
		Title   string
		Content string
		Html    string
	}

	// Setup the request to the target
	req, err := http.NewRequest("GET", "https://www.meneame.net", nil)

	client := &http.Client{}
	// Perform request and store response on "resp"
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Store Body
	body, _ := ioutil.ReadAll(resp.Body)

	news := r["news"].FindStringSubmatch(string(body))
	fmt.Println(news[1])

	title := r["title"].FindStringSubmatch(string(news[1]))
	fmt.Println(title[1])

	var sn [1]News
	sn[0].Level = 0
	sn[0].Title = "TITLE 0"
	sn[0].Content = "CONTENT 0"
	sn[0].Html = "<b>HTML 0</b>"

	fmt.Println(sn[0])
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	// Create writer to file
	f, _ := os.Create("made.html")
	defer f.Close()

	w := bufio.NewWriter(f)
	t.Execute(w, sn)
	w.Flush()
}
