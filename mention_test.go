package twitterscraper_test

import (
	"encoding/json"
	twitterscraper "github.com/imperatrona/twitter-scraper"
	"testing"
	"time"
)

func TestScraper_GetMentionNotifications(t *testing.T) {
	testScraper.SetAuthToken(twitterscraper.AuthToken{Token: authToken,
		CSRFToken: csrfToken})
	if !testScraper.IsLoggedIn() {
		panic("Invalid AuthToken")
	}

	var cursor string
	for {
		tweets, nextCursor, err := testScraper.GetMentionNotifications(cursor, 40)
		if err != nil {
			t.Error(err)
			continue
		}

		marshal, err := json.Marshal(tweets)
		if err != nil {
			return
		}
		t.Log("notify: ", string(marshal))

		t.Logf("cursor: %v\n", nextCursor)
		time.Sleep(10 * time.Second)
	}
}

func TestScraper_GetAllNotifications(t *testing.T) {
	testScraper.SetAuthToken(twitterscraper.AuthToken{Token: authToken, CSRFToken: csrfToken})
	if !testScraper.IsLoggedIn() {
		panic("Invalid AuthToken")
	}

	var cursor string
	for {
		tweets, nextCursor, err := testScraper.GetAllNotifications(cursor, 40)
		if err != nil {
			t.Error(err)
			continue
		}

		marshal, err := json.Marshal(tweets)
		if err != nil {
			return
		}
		t.Log("notify: ", string(marshal))

		t.Logf("cursor: %v\n", nextCursor)
		time.Sleep(10 * time.Second)
	}
}
