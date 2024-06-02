package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	if res, err := httpClient.Get(url); err != nil {
		return nil, err
	} else {
		defer res.Body.Close()
		if data, err := io.ReadAll(res.Body); err != nil {
			return nil, err
		} else {
			rssFeed := RSSFeed{}
			if err := xml.Unmarshal(data, &rssFeed); err != nil {
				return nil, err
			}
			return &rssFeed, nil
		}
	}
}
