package selectcontrol

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func getFasterWebsite(urlA, urlB, urlC string) string {
	channelA := getTitles(urlA)
	channelB := getTitles(urlB)
	channelC := getTitles(urlC)

	select {
	case titleA := <-channelA:
		return titleA
	case titleB := <-channelB:
		return titleB
	case titleC := <-channelC:
		return titleC
	case <-time.After(1000 * time.Millisecond):
		return "too slow"
	}
}

// Execute runs selectcontrol examples.
func Execute() {
	winner := getFasterWebsite(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)

	fmt.Println(fmt.Sprintf("The winner website is: %s", winner))
}
