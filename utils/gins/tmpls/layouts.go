package tmpls

import (
	"path/filepath"

	"github.com/cbhan755200839/mygo/utils/logs"
	"go.uber.org/zap"
)

func loadBaseTmpl() (tmplFiles []string, err error) {
	tmplFiles, err = filepath.Glob("templates/layout/*.tmpl")
	if err != nil {
		logs.Log.Error("模板路径templates/layout/*.tmpl解析失败", zap.Error(err))
	}
	return tmplFiles, err
}
