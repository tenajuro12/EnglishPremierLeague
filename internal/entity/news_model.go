package entity

import "time"

type NewsArticle struct {
	ID        int
	Title     string
	Content   string
	Category  string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
