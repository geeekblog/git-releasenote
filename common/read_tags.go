package common

import (
	"io"
	"sort"

	"github.com/sirupsen/logrus"

	"github.com/go-git/go-git/v5"
)

func ReadTags(repoPath string) ([]*Tag, error) {
	repo, err := git.PlainOpenWithOptions(repoPath, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	tagIter, err := repo.Tags()
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	rs := make([]*Tag, 0, 16)

	for {
		if tag, err := tagIter.Next(); err == nil {
			hash := tag.Hash()
			logrus.Debugln("source tag.Hash:" + tag.Hash().String() + "[Tag]:" + tag.Name().Short())
			if t, err := repo.TagObject(hash); err == nil {
				logrus.Debugln("changed tag.Hash" + t.Target.String() + "[Tag]:" + tag.Name().Short())
				hash = t.Target
			}
			c, err := repo.CommitObject(hash)
			if err != nil {
				logrus.Debugln("getCommit Error:" + err.Error())
				continue
			}

			t := &Tag{
				Name: tag.Name().Short(),
				Time: c.Committer.When,
				Hash: hash,
			}
			rs = append(rs, t)
		} else {
			if err == io.EOF {
				break
			} else {
				logrus.Errorln("tar.Next Error:" + err.Error())
				break
			}
		}
	}

	return rs, nil
}

func ReadSortedTags(repoPath string) ([]*Tag, error) {
	list, err := ReadTags(repoPath)
	if err != nil {
		logrus.Debugln("ReadTags Error:" + err.Error())
		return nil, err
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Time.UnixNano() > list[j].Time.UnixNano()
	})

	return list, nil
}
