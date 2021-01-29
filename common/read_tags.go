package common

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/go-git/go-git/v5"
)

func ReadTags(repoPath string) ([]*Tag, error) {
	repo, err := git.PlainOpenWithOptions(repoPath, &git.PlainOpenOptions{})
	if err != nil {
		return nil, err
	}

	tagIter, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	rs := make([]*Tag, 0, 16)

	for {
		if tag, err := tagIter.Next(); err == nil {
			hash := tag.Hash()
			if t, err := repo.TagObject(hash); err == nil {
				hash = t.Target
			}
			c, err := repo.CommitObject(hash)
			if err != nil {
				continue
			}

			t := &Tag{
				Name:    tag.Name().Short(),
				TagTime: c.Committer.When,
				Time:    c.Committer.When.Format("2006-01-02 15:04:05"),
			}
			rs = append(rs, t)
		} else {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintln(os.Stderr, err)
				break
			}
		}
	}
	return rs, nil
}

func ReadSortedTags(repoPath string) ([]*Tag, error) {
	list, err := ReadTags(repoPath)
	if err != nil {
		return nil, err
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].TagTime.UnixNano() > list[j].TagTime.UnixNano() {
			return true
		}
		return false
	})

	return list, nil
}
