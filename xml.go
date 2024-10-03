package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// fetch feed from URL, return filled out RSSFeed struct
func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// use http.client, requests, and do instead of get.

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		msg("Error making request.")
		return &RSSFeed{}, err
	}
	// common practice to identify program to server
	req.Header.Add("User-Agent", "gator")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		msg("Error doing request.")
		return &RSSFeed{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		msg("Error reading request.")
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		msg("Error unmarshaling xml")
		return &RSSFeed{}, err
	}

	// run html escape string to decide escaped html entities through title description fields of both entire channel and items
	// lmao

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[item].Description = html.UnescapeString(rssFeed.Channel.Item[item].Description)
		rssFeed.Channel.Item[item].Title = html.UnescapeString(rssFeed.Channel.Item[item].Title)
	}

	return &rssFeed, nil
}
