package entity

import "time"

type NewMatch struct {
	ID        int    `json:"id"`
	HomeTeam  string `json:"home_team"`
	AwayTeam  string `json:"away_team"`
	Date      string `json:"date"`
	Status    string `json:"status"`
	HomeScore int    `json:"home_score,omitempty"`
	AwayScore int    `json:"away_score,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
