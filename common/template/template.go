package template

import (
	"os"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
)

//按模板文件覆盖优先级获取模板
func GetTemplate(templateType Type, templateDir string) (*template.Template, error) {
	var temp *template.Template
	var err error
	pathSeparator := string(os.PathSeparator)

	//如果指定了文件夹，那么如果报错的话，则不会继续尝试其他的文件夹
	if templateDir != "" {
		templateDir = strings.TrimRight(templateDir, string(os.PathSeparator))
		return template.ParseFiles(templateDir + pathSeparator + TemplateMaps[templateType])
	}

	if temp, err := template.ParseFiles(appDir + pathSeparator + TemplateMaps[templateType]); err != nil {
		logrus.Debugln(err)
		if temp, err := template.ParseFiles(appHomeConfigDir + pathSeparator + TemplateMaps[templateType]); err != nil {
			logrus.Debugln(err)
			if temp, err := template.ParseFiles(systemConfigDir + pathSeparator + TemplateMaps[templateType]); err != nil {
				logrus.Debugln(err)
			} else {
				return temp, nil
			}
		} else {
			return temp, nil
		}
	} else {
		return temp, nil
	}

	switch templateType {
	case TypeChangeLog:
		return (&template.Template{}).Parse(defaultChangeLogTemplate)
	case TypeReleaseNote:
		return (&template.Template{}).Parse(defaultReleaseNoteTemplate)
	}
	return temp, err
}
