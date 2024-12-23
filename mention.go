package twitterscraper

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

const NotificationsMentionsAPI = "https://twitter.com/i/api/2/notifications/mentions.json"
const NotificationsAllAPI = "https://twitter.com/i/api/2/notifications/all.json"

// GetMentionNotifications
// Fetches mention notifications for the authenticated user, showing tweets that mention the user.
//
// Rate limit: 180 actions per 15 minutes.
//
// Parameters:
// cursor (str, optional): Pagination cursor for fetching a specific page of mention notifications.
//
// Returns:
// tuple: A tuple containing a list of tweets mentioning the user, the next cursor, and the previous cursor for pagination.
func (s *Scraper) GetMentionNotifications(cursor string, count uint64) ([]*Tweet, string, string, error) {
	if count > 40 {
		count = 40
	}
	if !s.isLogged {
		return nil, "", "", errors.New("scraper is not logged in for notifications")
	}

	var queryParams string
	if cursor != "" {
		queryParams += "&cursor=" + url.QueryEscape(cursor)
	}

	queryParams += "&count=" + strconv.FormatUint(count, 10)

	reqURL := fmt.Sprintf("%s?%s", NotificationsMentionsAPI, queryParams)
	req, err := s.newRequest("GET", reqURL)
	if err != nil {
		return nil, "", "", err
	}

	var timeline *timelineV1
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, "", "", err
	}

	tweets, bottomCursor, topCursor := timeline.parseTweetsDouble()
	return tweets, bottomCursor, topCursor, nil
}

func (s *Scraper) GetAllNotifications(cursor string, count uint64) ([]*Tweet, string, string, error) {
	if count > 40 {
		count = 40
	}
	if !s.isLogged {
		return nil, "", "", errors.New("scraper is not logged in for notifications")
	}

	var queryParams string
	if cursor != "" {
		queryParams += "&cursor=" + url.QueryEscape(cursor)
	}

	queryParams += "&count=" + strconv.FormatUint(count, 10)

	reqURL := fmt.Sprintf("%s?%s", NotificationsAllAPI, queryParams)
	req, err := s.newRequest("GET", reqURL)
	if err != nil {
		return nil, "", "", err
	}

	var timeline *timelineV1
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, "", "", err
	}

	tweets, bottomCursor, topCursor := timeline.parseTweetsDouble()
	return tweets, bottomCursor, topCursor, nil
}
