package main

import (
	"fmt"
	"strconv"
	"twt/download"
	"twt/readwrite"
)

func main() {
	//Scrape post
	// a := []string{}
	// b := []string{}
	// scraper := twitterscraper.New()
	// for tweet := range scraper.GetTweets(context.Background(), "FFOODFESS", 5) {
	// 	if tweet.Error != nil {
	// 		panic(tweet.Error)
	// 	}

	// 	b = append(b, tweet.Text)
	// 	a = append(a, tweet.Photos...)
	// }

	// if err := writeLines(a, "foo.txt"); err != nil {
	// 	log.Fatalf("writeLines: %s", err)
	// }

	// if err := writeLines(b, "text.txt"); err != nil {
	// 	log.Fatalf("writeLines: %s", err)
	// }

	lines, err := readwrite.ReadLines("images.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	for index, url := range lines {
		s := strconv.Itoa(index)
		filename := fmt.Sprintf("%s%s", s, ".jpg")
		err := download.Download(filename, url)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Download success")
		}
	}
}
