package common

import (
	"github.com/sirupsen/logrus"
)

type ChangeLog []*ChangeLogTagBody

type ChangeLogTagBody struct {
	Version *TagForShow
	Commits []*CommitForShow
}

const changeLogDebugFormat = "[%s:%s]%s ==> %s -> %s"

//调用的时候请确认commitList和tagList都是按时间倒序排列
func ToChangeLog(commitList []*Commit, tagList []*Tag) []*ChangeLogTagBody {
	list := make([]*ChangeLogTagBody, 0)
	if len(commitList) == 0 {
		logrus.Errorln("no commits in this repository")
		return list
	}
	commitIndex := 0
	//处理提交记录时间晚于最新Tag的情况
	logrus.Debugln("tag长度:", len(tagList))
	if len(tagList) == 0 || commitList[0].Time.After(tagList[0].Time) {
		tmpChangeLogBody := &ChangeLogTagBody{
			Version: &TagForShow{
				Name: messageNewestTag,
			},
			Commits: make([]*CommitForShow, 0),
		}
		for ; commitIndex < len(commitList); commitIndex++ {
			logrus.Debugf(changeLogDebugFormat, tmpChangeLogBody.Version.Name, "", commitList[commitIndex].Message, commitList[commitIndex].Time.Format(TimeFormat), commitList[commitIndex].Author)
			if len(tagList) > 0 && !commitList[commitIndex].Time.After(tagList[0].Time) {
				break
			}

			tmpChangeLogBody.Commits = append(tmpChangeLogBody.Commits, commitList[commitIndex].Show())
		}
		list = append(list, tmpChangeLogBody)
	}
	//处理提交记录在Tag内的情况
	for tagIndex := 0; tagIndex < len(tagList); tagIndex++ {
		tmpChangeLogBody := &ChangeLogTagBody{
			Version: tagList[tagIndex].Show(),
			Commits: make([]*CommitForShow, 0),
		}
		for ; commitIndex < len(commitList); commitIndex++ {
			logrus.Debugf(changeLogDebugFormat, tmpChangeLogBody.Version.Name, "", commitList[commitIndex].Message, commitList[commitIndex].Time.Format(TimeFormat), commitList[commitIndex].Author)
			//如果还有下一个标签，并且当前commit比下一个标签更早的话，需要开启下一个标签
			if tagIndex+1 < len(tagList) && !commitList[commitIndex].Time.After(tagList[tagIndex+1].Time) {
				break
			}

			tmpChangeLogBody.Commits = append(tmpChangeLogBody.Commits, commitList[commitIndex].Show())
		}
		list = append(list, tmpChangeLogBody)
		if commitIndex == len(commitList)-1 {
			break
		}
	}

	return list
}
