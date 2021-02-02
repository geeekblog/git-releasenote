package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Type uint8

const (
	_ Type = iota
	TypeChangeLog
	TypeReleaseNote

	ChangeLogFileName   = "CHANGELOG.template"
	ReleaseNoteFileName = "RELEASENOTE.template"
)

var (
	appDir           string
	appHomeConfigDir string
	systemConfigDir  string

	TemplateMaps = map[Type]string{
		TypeChangeLog:   ChangeLogFileName,
		TypeReleaseNote: ReleaseNoteFileName,
	}

	defaultChangeLogTemplate   = ``
	defaultReleaseNoteTemplate = ``
)

func init() {
	var err error
	//加载app所在文件夹
	app, err := os.Executable()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	appDir = filepath.Dir(app)
	appDir = strings.TrimRight(appDir, string(os.PathSeparator)) + string(os.PathSeparator) + "config"

	//加载home目录
	appHomeConfigDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	appHomeConfigDir = strings.TrimRight(appHomeConfigDir, string(os.PathSeparator)) + string(os.PathSeparator) + ".config/git-releasenote"

	//加载etc所在文件夹
	systemConfigDir = "/etc/git-releasenote"
}
