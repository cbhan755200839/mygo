package main

import (
	"mygo/utils/files"
	"mygo/utils/logs"
	"mygo/utils/yamls/configs"

	"go.uber.org/zap"
)

func main() {
	if err := mainInit(); err != nil {
		return
	}
}

func mainInit() (err error) {
	if err = configs.InitConfig(); err != nil {
		logs.Log.Error("配置文件配置初始化失败", zap.Error(err))
		return err
	}
	appPath, errPath := configs.Config.LogFile.AppPath, configs.Config.LogFile.ErrPath
	if err = files.CreateFile(appPath); err != nil {
		logs.Log.Error("应用日志文件创建失败", zap.Error(err))
		return err
	}
	if err = files.CreateFile(errPath); err != nil {
		logs.Log.Error("错误日志文件创建失败", zap.Error(err))
		return err
	}
	if err = logs.InitLog(appPath, errPath); err != nil {
		logs.Log.Error("日志初始化失败", zap.Error(err))
		return err
	}
	return nil
}
