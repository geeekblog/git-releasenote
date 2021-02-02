package common

import (
	"time"

	"github.com/go-git/go-git/v5/plumbing"
)

type Commit struct {
	Keyword Keyword
	Message string
	Body    string
	Author  string
	Email   string
	Hash    plumbing.Hash
	Time    time.Time
}

type CommitForShow struct {
	Keyword   Keyword
	Message   string
	Body      string
	Author    string
	Email     string
	ShortHash string
	Time      string
	Hash      string
}

func (commit *Commit) Show() *CommitForShow {
	return &CommitForShow{
		Keyword:   commit.Keyword,
		Message:   commit.Message,
		Body:      commit.Body,
		Author:    commit.Author,
		Email:     commit.Email,
		ShortHash: commit.Hash.String()[:8],
		Time:      commit.Time.Format(TimeFormat),
		Hash:      commit.Hash.String(),
	}
}

type Tag struct {
	Name string
	Hash plumbing.Hash
	Time time.Time
}

type TagForShow struct {
	Name string
	Time string
}

func (tag *Tag) Show() *TagForShow {
	return &TagForShow{
		Name: tag.Name,
		Time: tag.Time.Format(TimeFormat),
	}
}
