package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func main() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer ctxCancel()

	titles := getTitles(
		ctx,
		getTitleAsync(ctx, "https://www.amazon.com"),
		getTitleAsync(ctx, "https://www.google.com"),
		getTitleAsync(ctx, "https://www.youtube.com"),
	)

	for title := range titles {
		fmt.Println(title)
	}
}

func getTitles(
	ctx context.Context,
	googleChan,
	amazonChan,
	youtubeChan <-chan string,
) <-chan string {
	titlesChan := make(chan string)

	go func() {
		for {
			select {
			case title := <-googleChan:
				titlesChan <- title
			case title := <-amazonChan:
				titlesChan <- title
			case title := <-youtubeChan:
				titlesChan <- title
			case <-ctx.Done():
				close(titlesChan)
				return
			}
		}
	}()

	return titlesChan
}

func getTitleAsync(ctx context.Context, url string) <-chan string {
	titleChan := make(chan string)

	go func(ctx context.Context, url string) {
		title, _ := getTitle(ctx, url)
		titleChan <- title
	}(ctx, url)

	return titleChan
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
