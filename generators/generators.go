package generators

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getTitles(urls ...string) <-chan string {
	channel := make(chan string)

	for _, url := range urls {
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := ioutil.ReadAll(response.Body)

			regex, _ := regexp.Compile("<title>(.*?)<\\/title>")

			channel <- regex.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return channel
}

// Execute runs generators examples.
func Execute() {
	titles := getTitles("https://www.amazon.com", "https://www.google.com", "https://www.youtube.com")

	fmt.Println("Title 1:", <-titles)
	fmt.Println("Title 2:", <-titles)
	fmt.Println("Title 3:", <-titles)
}
