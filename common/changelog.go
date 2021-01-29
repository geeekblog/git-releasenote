package common

import (
	"fmt"
	"os"
)

type ChangeLog []*ChangeLogTagBody

type ChangeLogTagBody struct {
	Tag     *Tag
	Commits []*Commit
}

//调用的时候请确认commitList和tagList都是按时间倒序排列
func ToChangeLog(commitList []*Commit, tagList []*Tag) []*ChangeLogTagBody {
	list := make([]*ChangeLogTagBody, 0)
	if len(commitList) == 0 {
		fmt.Fprintln(os.Stderr, "no commit list")
		return list
	}
	commitIndex := 0
	//处理提交记录时间晚于最新Tag的情况
	if len(tagList) == 0 || commitList[0].CommitTime.After(tagList[0].TagTime) {
		tmpChangeLogBody := &ChangeLogTagBody{
			Tag: &Tag{
				Name: "NEWEST",
			},
			Commits: make([]*Commit, 0),
		}
		for ; commitIndex < len(commitList); commitIndex++ {
			if len(tagList) > 0 && !commitList[commitIndex].CommitTime.After(tagList[0].TagTime) {
				break
			}

			tmpChangeLogBody.Commits = append(tmpChangeLogBody.Commits, commitList[commitIndex])
		}
		list = append(list, tmpChangeLogBody)
	}
	//处理提交记录在Tag内的情况
	for tagIndex := 0; tagIndex < len(tagList); tagIndex++ {
		tmpChangeLogBody := &ChangeLogTagBody{
			Tag:     tagList[tagIndex],
			Commits: make([]*Commit, 0),
		}
		for ; commitIndex < len(commitList); commitIndex++ {
			//如果还有下一个标签，并且当前commit比下一个标签更早的话，需要开启下一个标签
			if tagIndex+1 < len(tagList) && !commitList[commitIndex].CommitTime.After(tagList[tagIndex+1].TagTime) {
				break
			}

			tmpChangeLogBody.Commits = append(tmpChangeLogBody.Commits, commitList[commitIndex])
		}
		list = append(list, tmpChangeLogBody)
		if commitIndex == len(commitList)-1 {
			break
		}
	}

	return list
}
