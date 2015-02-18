package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Initialize the map of regex
var r map[string]*regexp.Regexp

func main() {

	r = make(map[string]*regexp.Regexp)

	r["news"], _ = regexp.Compile(`<div class="news-summary">(.*?)</span></span>\s+</div>\s+</div>\s+</div>`)
	r["title"], _ = regexp.Compile(`<h2>\s+<a .*>(.*?)\s+</a>.*</h2>`)
	r["votes"], _ = regexp.Compile(`<div class="votes">\s+<a id=".*>(\d+)</a>\s+meneos\s+</div>`)

	type News struct {
		Votes     int
		Clicks    int
		Title     string
		User      string
		To        string
		Sent      string
		Published string
		Thumbnail string
		Details   struct {
			Users     int
			Anonymous int
			Comments  int
			Category  string
			Karma     int
		}
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

	news := r["news"].FindAllStringSubmatch(string(body), -1)
	fmt.Println(news)
	for k, _ := range news {
		//		fmt.Println(news[k][1])
		title := r["title"].FindStringSubmatch(string(news[k][1]))
		fmt.Println(title[1])
		votes := r["votes"].FindStringSubmatch(string(news[k][1]))
		fmt.Println(votes[1])
	}
	/*
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
	*/
}
