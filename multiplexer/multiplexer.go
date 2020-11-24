package multiplexer

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

// origin <-chan string: read channel only.
// destiny chan string: read and write channel.
func forward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

func merge(channels ...<-chan string) <-chan string {
	channel := make(chan string)
	for _, c := range channels {
		go forward(c, channel)
	}

	return channel
}

// Execute runs multiplexer examples.
func Execute() {
	titles := merge(
		getTitles("https://www.amazon.com", "https://www.google.com"),
		getTitles("https://www.youtube.com"),
	)

	fmt.Println("Title 1:", <-titles)
	fmt.Println("Title 2:", <-titles)
	fmt.Println("Title 3:", <-titles)
}
