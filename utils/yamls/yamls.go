package yamls

import (
	"encoding/json"
	"mygo/utils/logs"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var dir = "yamls/"

func LoadYaml(fileName string) (jsonStr string, err error) {
	fileData, err := os.ReadFile(dir + fileName)
	if err != nil {
		logs.Log.Fatal("读取yaml文件失败", zap.Error(err))
		return jsonStr, err
	}
	var yamlData = map[string]interface{}{}
	if err = yaml.Unmarshal(fileData, yamlData); err != nil {
		logs.Log.Fatal("解析yaml内容失败", zap.Error(err))
		return jsonStr, err
	}
	jsonData, err := json.Marshal(yamlData)
	if err != nil {
		logs.Log.Fatal("yaml内容转换json失败", zap.Error(err))
		return jsonStr, err
	}
	jsonStr = string(jsonData)
	return jsonStr, err
}
