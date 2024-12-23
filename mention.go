package twitterscraper

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

const NotificationsMentionsAPI = "https://twitter.com/i/api/2/notifications/mentions.json"
const NotificationsAllAPI = "https://twitter.com/i/api/2/notifications/all.json"

func (s *Scraper) GetMentionNotifications(cursor string, count uint64) ([]*Tweet, string, error) {
	if count > 40 {
		count = 40
	}
	if !s.isLogged {
		return nil, "", errors.New("scraper is not logged in for notifications")
	}

	var queryParams string
	if cursor != "" {
		queryParams += "&cursor=" + url.QueryEscape(cursor)
	}

	queryParams += "&count=" + strconv.FormatUint(count, 10)

	reqURL := fmt.Sprintf("%s?%s", NotificationsMentionsAPI, queryParams)
	req, err := s.newRequest("GET", reqURL)
	if err != nil {
		return nil, "", err
	}

	var timeline *timelineV1
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, "", err
	}

	tweets, nextCursor := timeline.parseTweets()
	return tweets, nextCursor, nil
}

func (s *Scraper) GetAllNotifications(cursor string, count uint64) ([]*Tweet, string, error) {
	if count > 40 {
		count = 40
	}
	if !s.isLogged {
		return nil, "", errors.New("scraper is not logged in for notifications")
	}

	var queryParams string
	if cursor != "" {
		queryParams += "&cursor=" + url.QueryEscape(cursor)
	}

	queryParams += "&count=" + strconv.FormatUint(count, 10)

	reqURL := fmt.Sprintf("%s?%s", NotificationsAllAPI, queryParams)
	req, err := s.newRequest("GET", reqURL)
	if err != nil {
		return nil, "", err
	}

	var timeline *timelineV1
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return nil, "", err
	}

	tweets, nextCursor := timeline.parseTweets()
	return tweets, nextCursor, nil
}
