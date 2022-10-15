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
	ctx, ctxCancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
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

func getTitles(ctx context.Context, titlesChans ...<-chan string) <-chan string {
	titlesChan := make(chan string, len(titlesChans))

	go func() {
		<-ctx.Done()
		close(titlesChan)
	}()

	for _, ch := range titlesChans {
		go func(ch <-chan string) {
			for {
				select {
				case title := <-ch:
					titlesChan <- title
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}

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
