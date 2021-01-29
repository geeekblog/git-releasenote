package common

import (
	"time"

	"github.com/go-git/go-git/v5/plumbing"
)

type Commit struct {
	Keyword    Keyword
	Message    string
	Body       string
	Author     string
	Email      string
	ShortHash  string
	Time       string
	Hash       plumbing.Hash
	CommitTime time.Time
}

type Tag struct {
	Name    string
	Time    string
	Hash    plumbing.Hash
	TagTime time.Time
}
