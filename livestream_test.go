package rumblelivestreamlib

import (
	"encoding/json"
	"reflect"
	"testing"
)

const exampleS string = `{
  "now": 1695059500,
  "type": "user",
  "user_id": "XXXXX",
  "channel_id": null,
  "since": null,
  "max_num_results": 50,
  "followers": {
    "num_followers": 165,
    "num_followers_total": 362,
    "latest_follower": {
      "username": "UserNameA",
      "followed_on": "2023-09-17T20:43:56-04:00"
    },
    "recent_followers": [
      {
        "username": "UserNameA",
        "followed_on": "2023-09-17T20:43:56-04:00"
      },
      {
        "username": "UserNameB",
        "followed_on": "2023-09-17T19:10:02-04:00"
      },
      {
        "username": "UserNameC",
        "followed_on": "2023-09-16T21:17:47-04:00"
      }
    ]
  },
  "subscribers": {
    "num_subscribers": 5,
    "latest_subscriber": {
      "username": "UserNameD",
      "user": "UserNameD",
      "amount_cents": 500,
      "amount_dollars": 5,
      "subscribed_on": "2023-09-15T21:03:05+00:00"
    },
    "recent_subscribers": [
      {
        "user": "UserNameE",
        "username": "UserNameE",
        "amount_cents": 500,
        "amount_dollars": 5,
        "subscribed_on": "2023-09-15T21:03:05+00:00"
      },
      {
        "user": "UserNameF",
        "username": "UserNameF",
        "amount_cents": 500,
        "amount_dollars": 5,
        "subscribed_on": "2023-09-15T20:50:50+00:00"
      },
      {
        "user": "UserNameG",
        "username": "UserNameG",
        "amount_cents": 500,
        "amount_dollars": 5,
        "subscribed_on": "2023-04-09T01:12:26+00:00"
      }
    ]
  },
  "livestreams": [
    {
      "id": "XXXXX",
      "title": "Title of Live Stream",
      "created_on": "2023-09-18T17:51:19+00:00",
      "is_live": true,
      "categories": {
        "primary": { 
          "slug": "gaming",
          "title": "Gaming"
        },
        "secondary": {
          "slug": "simulators",
          "title": "Simulators"
        }
      }, 
      "stream_key": "XXXXXX",
      "likes": 3,
      "dislikes": 1,
      "watching_now": 19,
      "chat": {
        "latest_message": {
          "username": "UserNameH",
          "badges": [
            "admin",
            "premium",
            "whale-blue"
          ],
          "text": "test chat",
          "created_on": "2023-09-18T17:51:08+00:00"
        },
        "recent_messages": [
          {
            "username": "UserNameA",
            "badges": [
              "admin",
              "premium",
              "whale-blue"
            ],
            "text": "This is my chat message",
            "created_on": "2023-09-18T17:51:08+00:00"
          }
        ],
        "latest_rant": {
          "username": "UserNameB",
          "badges": [
            "admin",
            "premium",
            "whale-blue"
          ],
          "text": "Rant Message Goes Here",
          "created_on": "2023-09-18T17:51:37+00:00",
          "expires_on": "2023-09-18T17:53:37+00:00",
          "amount_cents": 100,
          "amount_dollars": 1
        },
        "recent_rants": [
          {
            "username": "UserNameB",
            "badges": [
              "admin",
              "premium",
              "whale-blue"
            ],
            "text": "Rant Message",
            "created_on": "2023-09-18T17:51:37+00:00",
            "expires_on": "2023-09-18T17:53:37+00:00",
            "amount_cents": 100,
            "amount_dollars": 1
          },
          {
            "username": "UserNameC",
            "badges": [
              "admin",
              "premium",
              "whale-blue"
            ],
            "text": "Rant Message",
            "created_on": "2023-09-18T17:51:37+00:00",
            "expires_on": "2023-09-18T17:53:37+00:00",
            "amount_cents": 100,
            "amount_dollars": 1
          },{
            "username": "UserNameD",
            "badges": [
              "admin",
              "premium",
              "whale-blue"
            ],
            "text": "Rant Message",
            "created_on": "2023-09-18T17:51:37+00:00",
            "expires_on": "2023-09-18T17:53:37+00:00",
            "amount_cents": 100,
            "amount_dollars": 1
          }
        ]
      }
    }
  ]
}`

var example = LivestreamResponse{
	Now:           1695059500,
	Type:          "user",
	UserID:        "XXXXX",
	MaxNumResults: 50,
	Followers: Followers{
		NumFollowers:      165,
		NumFollowersTotal: 362,
		LatestFollower: Follower{
			Username:   "UserNameA",
			FollowedOn: "2023-09-17T20:43:56-04:00",
		},
		RecentFollowers: []Follower{
			{
				Username:   "UserNameA",
				FollowedOn: "2023-09-17T20:43:56-04:00",
			},
			{
				Username:   "UserNameB",
				FollowedOn: "2023-09-17T19:10:02-04:00",
			},
			{
				Username:   "UserNameC",
				FollowedOn: "2023-09-16T21:17:47-04:00",
			},
		},
	},
	Subscribers: Subscribers{
		NumSubscribers: 5,
		LatestSubscriber: Subscriber{
			Username:      "UserNameD",
			User:          "UserNameD",
			AmountCents:   500,
			AmountDollars: 5,
			SubscribedOn:  "2023-09-15T21:03:05+00:00",
		},
		RecentSubscribers: []Subscriber{
			{
				Username:      "UserNameE",
				User:          "UserNameE",
				AmountCents:   500,
				AmountDollars: 5,
				SubscribedOn:  "2023-09-15T21:03:05+00:00",
			},
			{
				Username:      "UserNameF",
				User:          "UserNameF",
				AmountCents:   500,
				AmountDollars: 5,
				SubscribedOn:  "2023-09-15T20:50:50+00:00",
			},
			{
				Username:      "UserNameG",
				User:          "UserNameG",
				AmountCents:   500,
				AmountDollars: 5,
				SubscribedOn:  "2023-04-09T01:12:26+00:00",
			},
		},
	},
	Livestreams: []Livestream{
		{
			ID:        "XXXXX",
			Title:     "Title of Live Stream",
			CreatedOn: "2023-09-18T17:51:19+00:00",
			IsLive:    true,
			Categories: Categories{
				Primary: Category{
					Slug:  "gaming",
					Title: "Gaming",
				},
				Secondary: Category{
					Slug:  "simulators",
					Title: "Simulators",
				},
			},
			StreamKey:   "XXXXXX",
			Likes:       3,
			Dislikes:    1,
			WatchingNow: 19,
			Chat: Chat{
				LatestMessage: Message{
					Username: "UserNameH",
					Badges: []Badge{
						"admin",
						"premium",
						"whale-blue",
					},
					Text:      "test chat",
					CreatedOn: "2023-09-18T17:51:08+00:00",
				},
				RecentMessages: []Message{
					{
						Username: "UserNameA",
						Badges: []Badge{
							"admin",
							"premium",
							"whale-blue",
						},
						Text:      "This is my chat message",
						CreatedOn: "2023-09-18T17:51:08+00:00",
					},
				},
				LatestRant: Rant{
					Message: Message{
						Username: "UserNameB",
						Badges: []Badge{
							"admin",
							"premium",
							"whale-blue",
						},
						Text:      "Rant Message Goes Here",
						CreatedOn: "2023-09-18T17:51:37+00:00",
					},
					ExpiresOn:     "2023-09-18T17:53:37+00:00",
					AmountCents:   100,
					AmountDollars: 1,
				},
				RecentRants: []Rant{
					{
						Message: Message{
							Username: "UserNameB",
							Badges: []Badge{
								"admin",
								"premium",
								"whale-blue",
							},
							Text:      "Rant Message",
							CreatedOn: "2023-09-18T17:51:37+00:00",
						},
						ExpiresOn:     "2023-09-18T17:53:37+00:00",
						AmountCents:   100,
						AmountDollars: 1,
					},
					{
						Message: Message{
							Username: "UserNameC",
							Badges: []Badge{
								"admin",
								"premium",
								"whale-blue",
							},
							Text:      "Rant Message",
							CreatedOn: "2023-09-18T17:51:37+00:00",
						},
						ExpiresOn:     "2023-09-18T17:53:37+00:00",
						AmountCents:   100,
						AmountDollars: 1,
					},
					{
						Message: Message{
							Username: "UserNameD",
							Badges: []Badge{
								"admin",
								"premium",
								"whale-blue",
							},
							Text:      "Rant Message",
							CreatedOn: "2023-09-18T17:51:37+00:00",
						},
						ExpiresOn:     "2023-09-18T17:53:37+00:00",
						AmountCents:   100,
						AmountDollars: 1,
					},
				},
			},
		},
	},
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want LivestreamResponse
	}{
		{"Website example", exampleS, example},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got LivestreamResponse
			err := json.Unmarshal([]byte(tt.arg), &got)
			if err != nil {
				t.Fatalf("want json.Unmarshal = nil, got %v", err)
			}

			equal := reflect.DeepEqual(got, tt.want)
			if !equal {
				t.Fatalf("want reflect.DeepEqual = true, got false")
			}
		})
	}
}
