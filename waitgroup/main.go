package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"
)

func main() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer ctxCancel()

	urls := []string{
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.youtube.com",
	}

	titles := getTitles(ctx, urls)

	fmt.Println("titles: ", titles)
}

func getTitles(ctx context.Context, urls []string) []string {
	var waitgroup sync.WaitGroup
	var titles = make([]string, len(urls))

	for index, url := range urls {
		waitgroup.Add(1)

		go func(ctx context.Context, index int, url string) {
			defer waitgroup.Done()
			title, _ := getTitle(ctx, url)
			titles[index] = title
		}(ctx, index, url)
	}

	waitgroup.Wait()

	return titles
}

func getTitle(ctx context.Context, url string) (string, error) {
	client := http.DefaultClient

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	request = request.WithContext(ctx)

	httpResponse, err := client.Do(request)
	if err != nil {
		return "", err
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}

	findBodyTitleRegex, _ := regexp.Compile(`<title[^>]*>(.*?)</title>`)

	findResult := findBodyTitleRegex.FindStringSubmatch(string(responseBody))
	if len(findResult) == 0 {
		return "", errors.New("cannot retrieve website url")
	}

	return findResult[1], nil
}
