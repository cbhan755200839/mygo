package tmpls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExecuteTmpl(c *gin.Context, dir string, name string, data interface{}) {
	tmpl, err := loadTmpl(dir + "/" + name + ".tmpl")
	if err != nil {
		c.String(http.StatusInternalServerError, "模板解析失败：%v", err)
		return
	}
	if err = tmpl.ExecuteTemplate(c.Writer, name+".tmpl", nil); err != nil {
		c.String(http.StatusInternalServerError, "模板渲染失败：%v", err)
	}
}
