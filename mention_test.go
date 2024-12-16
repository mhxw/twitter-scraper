package twitterscraper_test

import (
	"encoding/json"
	twitterscraper "github.com/imperatrona/twitter-scraper"
	"testing"
	"time"
)

func TestScraper_GetMentionNotifications(t *testing.T) {
	testScraper.SetAuthToken(twitterscraper.AuthToken{Token: authToken, CSRFToken: csrfToken})
	if !testScraper.IsLoggedIn() {
		panic("Invalid AuthToken")
	}

	var cursor string
	for {
		notifications, err := testScraper.GetMentionNotifications(cursor, 40)
		if err != nil {
			t.Error(err)
			continue
		}

		marshal, err := json.Marshal(notifications)
		if err != nil {
			return
		}
		t.Log("notify: ", string(marshal))

		for k, v := range notifications.GlobalObjects.Tweets {
			t.Log(k, v.FullText)

		}

		for _, v2 := range notifications.Timeline.Instructions {
			if len(v2.AddEntries.Entries) > 1 {
				cursor = v2.AddEntries.Entries[0].Content.Operation.Cursor.Value
			}
		}

		t.Logf("cursor: %v\n", cursor)
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
		notifications, err := testScraper.GetMentionNotifications(cursor, 40)
		if err != nil {
			t.Error(err)
			continue
		}

		marshal, err := json.Marshal(notifications)
		if err != nil {
			return
		}
		t.Log("notify: ", string(marshal))

		for k, v := range notifications.GlobalObjects.Tweets {
			t.Log(k, v.FullText)

		}

		for _, v2 := range notifications.Timeline.Instructions {
			if len(v2.AddEntries.Entries) > 1 {
				cursor = v2.AddEntries.Entries[0].Content.Operation.Cursor.Value
			}
		}

		t.Logf("cursor: %v\n", cursor)
		time.Sleep(10 * time.Second)
	}
}
