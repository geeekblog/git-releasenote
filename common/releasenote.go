package common

import (
	"github.com/sirupsen/logrus"
)

type ReleaseNote []*ReleaseNoteBody

type ReleaseNoteBody struct {
	Version *TagForShow
	Group   []*ReleaseNoteGroup
}

type ReleaseNoteGroup struct {
	Keyword  string
	Messages []string
}

//调用的时候请确认commitList和tagList都是按时间倒序排列
func ToReleaseNote(commitList []*Commit, tagList []*Tag) []*ReleaseNoteBody {
	list := make([]*ReleaseNoteBody, 0)
	if len(commitList) == 0 {
		logrus.Errorln("no commits in this repository")
		return list
	}

	commitIndex := 0
	//处理提交记录时间晚于最新Tag的情况
	if len(tagList) == 0 || commitList[0].Time.After(tagList[0].Time) {
		tmpReleaseNoteBody := &ReleaseNoteBody{
			Version: &TagForShow{
				Name: messageNewestTag,
			},
			Group: make([]*ReleaseNoteGroup, 0),
		}
		tmpGroupMap := make(map[Keyword]*ReleaseNoteGroup, 2)
		for ; commitIndex < len(commitList); commitIndex++ {
			if len(tagList) > 0 && !commitList[commitIndex].Time.After(tagList[0].Time) {
				break
			}
			setToGroup(commitList[commitIndex], tmpGroupMap)
		}

		//map中的内容合并到slice中
		if v, ok := tmpGroupMap[GroupKeywordFeature]; ok {
			v.Messages = stringSliceUnique(v.Messages)
			tmpReleaseNoteBody.Group = append(tmpReleaseNoteBody.Group, v)
		}
		if v, ok := tmpGroupMap[GroupKeywordOther]; ok {
			v.Messages = stringSliceUnique(v.Messages)
			tmpReleaseNoteBody.Group = append(tmpReleaseNoteBody.Group, v)
		}
		list = append(list, tmpReleaseNoteBody)
	}
	//处理提交记录在Tag内的情况
	for tagIndex := 0; tagIndex < len(tagList); tagIndex++ {
		tmpReleaseNoteBody := &ReleaseNoteBody{
			Version: tagList[tagIndex].Show(),
			Group:   make([]*ReleaseNoteGroup, 0),
		}
		tmpGroupMap := make(map[Keyword]*ReleaseNoteGroup, 2)
		for ; commitIndex < len(commitList); commitIndex++ {
			//如果还有下一个标签，并且当前commit比下一个标签更早的话，需要开启下一个标签
			if tagIndex+1 < len(tagList) && !commitList[commitIndex].Time.After(tagList[tagIndex+1].Time) {
				break
			}
			setToGroup(commitList[commitIndex], tmpGroupMap)
		}
		//map中的内容合并到slice中
		if v, ok := tmpGroupMap[GroupKeywordFeature]; ok {
			v.Messages = stringSliceUnique(v.Messages)
			tmpReleaseNoteBody.Group = append(tmpReleaseNoteBody.Group, v)
		}
		if v, ok := tmpGroupMap[GroupKeywordOther]; ok {
			v.Messages = stringSliceUnique(v.Messages)
			tmpReleaseNoteBody.Group = append(tmpReleaseNoteBody.Group, v)
		}
		list = append(list, tmpReleaseNoteBody)
		if commitIndex == len(commitList)-1 {
			break
		}
	}

	return list
}

func setToGroup(commit *Commit, groupMap map[Keyword]*ReleaseNoteGroup) {
	switch commit.Keyword {
	case KeywordFeature:
		if g, ok := groupMap[GroupKeywordFeature]; ok {
			g.Messages = append(g.Messages, commit.Message)
		} else {
			groupMap[GroupKeywordFeature] = &ReleaseNoteGroup{
				Keyword:  GroupKeywordFeature,
				Messages: []string{commit.Message},
			}
		}
	case KeywordBugFix:
		if g, ok := groupMap[GroupKeywordOther]; ok {
			g.Messages = append(g.Messages, messageBugFix)
		} else {
			groupMap[GroupKeywordOther] = &ReleaseNoteGroup{
				Keyword:  GroupKeywordOther,
				Messages: []string{messageBugFix},
			}
		}
	case KeywordPref:
		if g, ok := groupMap[GroupKeywordOther]; ok {
			g.Messages = append(g.Messages, messagePerformanceOptimization)
		} else {
			groupMap[GroupKeywordOther] = &ReleaseNoteGroup{
				Keyword:  GroupKeywordOther,
				Messages: []string{messagePerformanceOptimization},
			}
		}
	}
}

func stringSliceUnique(list []string) []string {
	m := make(map[string]struct{}, len(list))
	for _, v := range list {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	rs := make([]string, 0, len(m))
	for v, _ := range m {
		rs = append(rs, v)
	}
	return rs
}
