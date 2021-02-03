package common

import (
	"io"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-git/go-git/v5"
)

//从git log中读取需要的内容
func ReadLogs(repoPath string, from, to *time.Time) ([]*Commit, error) {
	repo, err := git.PlainOpenWithOptions(repoPath, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return nil, err
	}
	opt := &git.LogOptions{
		Order: git.LogOrderCommitterTime,
		Until: to,
	}
	if !from.IsZero() {
		opt.Since = from
	}
	commitIter, err := repo.Log(opt)
	if err != nil {
		return nil, err
	}

	rsList := make([]*Commit, 0, 16)
	for {
		if commit, err := commitIter.Next(); err == nil {
			k, m := parseMessage(commit.Message)
			tmpRs := &Commit{
				Keyword: k,
				Message: m,
				Body:    strings.TrimSpace(commit.Message),
				Author:  strings.TrimSpace(commit.Author.Name),
				Email:   strings.TrimSpace(commit.Author.Email),
				Hash:    commit.Hash,
				Time:    commit.Committer.When,
			}
			if k == KeywordUnknown || m == "" {
				continue
			}
			logrus.Debugln(tmpRs.Hash.String() + "[" + tmpRs.Author + "]" + string(tmpRs.Keyword) + ":" + tmpRs.Message)
			rsList = append(rsList, tmpRs)
		} else {
			if err == io.EOF {
				break
			}
		}
	}

	return rsList, nil
}

func parseMessage(message string) (Keyword, string) {
	body := strings.Split(message, "\n")
	header := body[0]
	for _, key := range KeywordList {
		if i := strings.Index(header, string(key)); i == 0 {
			return key, strings.TrimSpace(strings.Replace(header, string(key)+": ", "", 1))
		}
	}
	return KeywordUnknown, ""
}
