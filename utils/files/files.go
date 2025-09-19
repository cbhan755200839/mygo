package files

import (
	"fmt"
	"mygo/utils/logs"
	"os"
	"path/filepath"
)

func CreateFile(path string) (err error) {
	// 检查文件是否存在
	_, err = os.Stat(path)

	// 如果文件不存在
	if os.IsNotExist(err) {
		// 判断文件夹是否存在
		dir := filepath.Dir(path)
		// 文件夹若不存在，新建
		if err := os.MkdirAll(dir, 0744); err != nil {
			return err
		}

		var file *os.File
		// 新建文件
		file, err = os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	} else {
		logs.Log.Info(fmt.Sprintf("文件%s已存在", path))
	}

	return nil
}
