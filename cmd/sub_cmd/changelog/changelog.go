package changelog

import (
	"fmt"
	"git-releasenote/common"
	"git-releasenote/common/template"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "changelog",
		Short: "make change log",
		Long:  "",
		Run:   Run,
	}
	repoPath    *string
	newest      *bool
	sinceTime   time.Time
	tagsNum     *int
	outPutFile  *string
	since       *string
	templateDir *string
)

func init() {
	repoPath = Command.PersistentFlags().String("repo_path", "./", "target repo path")
	newest = Command.PersistentFlags().Bool("newest", false, "show new commit after latest tag")
	since = Command.PersistentFlags().String("since", "0000-00-00 00:00:00", "show commits from time")
	tagsNum = Command.PersistentFlags().Int("tags", -1, "show last few commits")
	outPutFile = Command.PersistentFlags().StringP("output", "o", "", "out put to file")
	templateDir = Command.PersistentFlags().StringP("template", "t", "", "set template dir")
}

func Run(cmd *cobra.Command, args []string) {
	//解析出since
	var err error
	if *since != "0000-00-00 00:00:00" {
		sinceTime, err = time.Parse(common.TimeFormat, *since)
		if err != nil {
			fmt.Println("since is invalid:", err)
		}
	}

	//获取排序后的tag，从小到大
	tags, err := common.ReadSortedTags(*repoPath)
	if err != nil {
		os.Exit(1)
	}

	var endTime time.Time
	var fromTime time.Time

	//根据args内容，进行时间的确定
	if *newest {
		fromTime = tags[0].Time
		endTime = time.Now()
	} else if *tagsNum > 0 && len(tags) > 0 {
		if *tagsNum > len(tags) {
			fromTime = tags[len(tags)-1].Time
		} else {
			fromTime = tags[*tagsNum].Time
		}
		endTime = tags[0].Time
	} else if !sinceTime.IsZero() {
		fromTime = sinceTime
		endTime = time.Now()
	} else {
		endTime = time.Now()
	}

	//拆分出时间范围内的tag
	fromTagIndex := len(tags)
	for index, t := range tags {
		if t.Time.Before(fromTime) {
			if index != 0 {
				fromTagIndex = index
			}
			break
		}
	}
	tags = tags[:fromTagIndex]

	//因为不能包含其实时间tag当时的commit
	fromTime = fromTime.Add(time.Nanosecond)

	list, err := common.ReadLogs(*repoPath, &fromTime, &endTime)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	t, err := template.GetTemplate(template.TypeChangeLog, *templateDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cl := common.ToChangeLog(list, tags)

	//整理模板需要的结构
	var wr io.Writer
	wr = os.Stdout
	if *outPutFile != "" {
		file, err := os.Create(*outPutFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer file.Close()

		wr = file
	}
	err = t.Execute(wr, cl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
