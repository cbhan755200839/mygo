package tmpls

import (
	"html/template"

	"github.com/cbhan755200839/mygo/utils/logs"
	"go.uber.org/zap"
)

func loadTmpl(tmplName string) (tmpl *template.Template, err error) {
	tmplFiles, err := loadBaseTmpl()
	if err != nil {
		return tmpl, err
	}
	tmplFiles = append(tmplFiles, "templates/"+tmplName)
	tmpl, err = template.ParseFiles(tmplFiles...)
	if err != nil {
		logs.Log.Error("模板解析失败", zap.Error(err))
		return tmpl, err
	}
	return tmpl, err
}
