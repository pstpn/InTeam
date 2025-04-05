package model

import "time"

type Feedback struct {
	RacketID int       `json:"racket_id"`
	UserID   int       `json:"user_id"`
	Text     string    `json:"feedback"`
	Date     time.Time `json:"date"`
	Rating   int       `json:"rating"`
}

type FeedbackWithUsername struct {
	Username string `json:"username"`
	Feedback
}
