package twitterscraper

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

const NotificationsMentionsAPI = "https://twitter.com/i/api/2/notifications/mentions.json"
const NotificationsAllAPI = "https://twitter.com/i/api/2/notifications/all.json"

type User struct {
	Id          int64       `json:"id"`
	IdStr       string      `json:"id_str"`
	Name        string      `json:"name"`
	ScreenName  string      `json:"screen_name"`
	Location    interface{} `json:"location"`
	Description string      `json:"description"`
	Url         interface{} `json:"url"`
	Entities    struct {
		Description struct {
			Urls []struct {
				Url         string `json:"url"`
				ExpandedUrl string `json:"expanded_url"`
				DisplayUrl  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	Protected                      bool          `json:"protected"`
	FollowersCount                 int           `json:"followers_count"`
	FriendsCount                   int           `json:"friends_count"`
	ListedCount                    int           `json:"listed_count"`
	CreatedAt                      string        `json:"created_at"`
	FavouritesCount                int           `json:"favourites_count"`
	UtcOffset                      interface{}   `json:"utc_offset"`
	TimeZone                       interface{}   `json:"time_zone"`
	GeoEnabled                     bool          `json:"geo_enabled"`
	Verified                       bool          `json:"verified"`
	StatusesCount                  int           `json:"statuses_count"`
	Lang                           interface{}   `json:"lang"`
	ContributorsEnabled            bool          `json:"contributors_enabled"`
	IsTranslator                   bool          `json:"is_translator"`
	IsTranslationEnabled           bool          `json:"is_translation_enabled"`
	ProfileBackgroundColor         string        `json:"profile_background_color"`
	ProfileBackgroundImageUrl      interface{}   `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps interface{}   `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool          `json:"profile_background_tile"`
	ProfileImageUrl                string        `json:"profile_image_url"`
	ProfileImageUrlHttps           string        `json:"profile_image_url_https"`
	ProfileBannerUrl               string        `json:"profile_banner_url"`
	ProfileLinkColor               string        `json:"profile_link_color"`
	ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string        `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
	DefaultProfile                 bool          `json:"default_profile"`
	DefaultProfileImage            bool          `json:"default_profile_image"`
	Following                      bool          `json:"following"`
	FollowRequestSent              interface{}   `json:"follow_request_sent"`
	Notifications                  interface{}   `json:"notifications"`
	Blocking                       bool          `json:"blocking"`
	BlockedBy                      bool          `json:"blocked_by"`
	WantRetweets                   bool          `json:"want_retweets"`
	ProfileInterstitialType        string        `json:"profile_interstitial_type"`
	TranslatorType                 string        `json:"translator_type"`
	WithheldInCountries            []interface{} `json:"withheld_in_countries"`
	FollowedBy                     bool          `json:"followed_by"`
	ExtIsBlueVerified              bool          `json:"ext_is_blue_verified"`
	ExtHighlightedLabel            struct {
	} `json:"ext_highlighted_label"`
}

type MentionResult struct {
	GlobalObjects struct {
		Users  map[string]User         `json:"users"`
		Tweets map[string]MentionTweet `json:"tweets"`
	} `json:"globalObjects"`
	Timeline struct {
		Id           string `json:"id"`
		Instructions []struct {
			ClearCache struct {
			} `json:"clearCache,omitempty"`
			AddEntries struct {
				Entries []struct {
					EntryId   string `json:"entryId"`
					SortIndex string `json:"sortIndex"`
					Content   struct {
						Operation struct {
							Cursor struct {
								Value      string `json:"value"`
								CursorType string `json:"cursorType"`
							} `json:"cursor"`
						} `json:"operation,omitempty"`
						Item struct {
							Content struct {
								Tweet struct {
									Id          string `json:"id"`
									DisplayType string `json:"displayType"`
								} `json:"tweet"`
							} `json:"content"`
							ClientEventInfo struct {
								Component string `json:"component"`
								Element   string `json:"element"`
								Details   struct {
									NotificationDetails struct {
										ImpressionId string `json:"impressionId"`
										Metadata     string `json:"metadata"`
									} `json:"notificationDetails"`
								} `json:"details"`
							} `json:"clientEventInfo"`
						} `json:"item,omitempty"`
					} `json:"content"`
				} `json:"entries"`
			} `json:"addEntries,omitempty"`
			ClearEntriesUnreadState struct {
			} `json:"clearEntriesUnreadState,omitempty"`
			MarkEntriesUnreadGreaterThanSortIndex struct {
				SortIndex string `json:"sortIndex"`
			} `json:"markEntriesUnreadGreaterThanSortIndex,omitempty"`
		} `json:"instructions"`
	} `json:"timeline"`
}

type MentionTweet struct {
	CreatedAt        string `json:"created_at"`
	Id               int64  `json:"id"`
	IdStr            string `json:"id_str"`
	FullText         string `json:"full_text"`
	Truncated        bool   `json:"truncated"`
	DisplayTextRange []int  `json:"display_text_range"`
	Entities         struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []struct {
			ScreenName string `json:"screen_name"`
			Name       string `json:"name"`
			Id         int64  `json:"id"`
			IdStr      string `json:"id_str"`
			Indices    []int  `json:"indices"`
		} `json:"user_mentions"`
	} `json:"entities"`
	Source               string      `json:"source"`
	InReplyToStatusId    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserId      int64       `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string      `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  string      `json:"in_reply_to_screen_name"`
	UserId               int64       `json:"user_id"`
	UserIdStr            string      `json:"user_id_str"`
	Geo                  interface{} `json:"geo"`
	Coordinates          interface{} `json:"coordinates"`
	Place                interface{} `json:"place"`
	Contributors         interface{} `json:"contributors"`
	IsQuoteStatus        bool        `json:"is_quote_status"`
	RetweetCount         int         `json:"retweet_count"`
	FavoriteCount        int         `json:"favorite_count"`
	ReplyCount           int         `json:"reply_count"`
	QuoteCount           int         `json:"quote_count"`
	ConversationId       int64       `json:"conversation_id"`
	ConversationIdStr    string      `json:"conversation_id_str"`
	ConversationMuted    bool        `json:"conversation_muted"`
	Favorited            bool        `json:"favorited"`
	Retweeted            bool        `json:"retweeted"`
	PossiblySensitive    bool        `json:"possibly_sensitive"`
	Lang                 string      `json:"lang"`
}

func (s *Scraper) GetMentionNotifications(cursor string, count uint64) (*MentionResult, error) {
	if !s.isLogged {
		return nil, errors.New("scraper is not logged in for notifications")
	}

	queryParams := "include_profile_interstitial_type=1" +
		"&include_blocking=1" +
		"&include_blocked_by=1" +
		"&include_followed_by=1" +
		"&include_want_retweets=1" +
		"&include_mute_edge=1" +
		"&include_can_dm=1" +
		"&include_can_media_tag=1" +
		"&include_ext_is_blue_verified=1" +
		"&include_ext_verified_type=1" +
		"&include_ext_profile_image_shape=1" +
		"&skip_status=1" +
		"&cards_platform=Web-12" +
		"&include_cards=1" +
		"&include_ext_alt_text=true" +
		"&include_ext_limited_action_results=true" +
		"&include_quote_count=true" +
		"&include_reply_count=1" +
		"&tweet_mode=extended" +
		"&include_ext_views=true" +
		"&include_entities=true" +
		"&include_user_entities=true" +
		"&include_ext_media_color=true" +
		"&include_ext_media_availability=true" +
		"&include_ext_sensitive_media_warning=true" +
		"&include_ext_trusted_friends_metadata=true" +
		"&send_error_codes=true" +
		"&simple_quoted_tweet=true" +
		"&requestContext=launch" +
		"&ext=mediaStats,highlightedLabel,voiceInfo,birdwatchPivot,superFollowMetadata,unmentionInfo,editControl,article"

	if cursor != "" {
		queryParams += "&cursor=" + url.QueryEscape(cursor)
	}

	if count > 0 {
		queryParams += "&count=20"
	} else {
		queryParams += "&count=" + strconv.FormatUint(count, 10)
	}

	reqURL := fmt.Sprintf("%s?%s", NotificationsMentionsAPI, queryParams)
	req, err := s.newRequest("GET", reqURL)
	if err != nil {
		return nil, err
	}

	var mentionResp *MentionResult
	err = s.RequestAPI(req, &mentionResp)
	if err != nil {
		return nil, err
	}
	return mentionResp, nil
}

func (s *Scraper) GetAllNotifications(cursor string, count uint64) (*MentionResult, error) {
	if !s.isLogged {
		return nil, errors.New("scraper is not logged in for notifications")
	}

	queryParams := "include_profile_interstitial_type=1" +
		"&include_blocking=1" +
		"&include_blocked_by=1" +
		"&include_followed_by=1" +
		"&include_want_retweets=1" +
		"&include_mute_edge=1" +
		"&include_can_dm=1" +
		"&include_can_media_tag=1" +
		"&include_ext_is_blue_verified=1" +
		"&include_ext_verified_type=1" +
		"&include_ext_profile_image_shape=1" +
		"&skip_status=1" +
		"&cards_platform=Web-12" +
		"&include_cards=1" +
		"&include_ext_alt_text=true" +
		"&include_ext_limited_action_results=true" +
		"&include_quote_count=true" +
		"&include_reply_count=1" +
		"&tweet_mode=extended" +
		"&include_ext_views=true" +
		"&include_entities=true" +
		"&include_user_entities=true" +
		"&include_ext_media_color=true" +
		"&include_ext_media_availability=true" +
		"&include_ext_sensitive_media_warning=true" +
		"&include_ext_trusted_friends_metadata=true" +
		"&send_error_codes=true" +
		"&simple_quoted_tweet=true" +
		"&requestContext=launch" +
		"&ext=mediaStats,highlightedLabel,voiceInfo,birdwatchPivot,superFollowMetadata,unmentionInfo,editControl,article"

	if cursor != "" {
		queryParams += "&cursor=" + url.QueryEscape(cursor)
	}

	if count > 0 {
		queryParams += "&count=20"
	} else {
		queryParams += "&count=" + strconv.FormatUint(count, 10)
	}

	reqURL := fmt.Sprintf("%s?%s", NotificationsAllAPI, queryParams)
	req, err := s.newRequest("GET", reqURL)
	if err != nil {
		return nil, err
	}

	var mentionResp *MentionResult
	err = s.RequestAPI(req, &mentionResp)
	if err != nil {
		return nil, err
	}
	return mentionResp, nil
}
