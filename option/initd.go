package option

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"snds/model"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pelletier/go-toml"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Option = new(model.Option)
var Logger = logrus.New()

func Init() error {

	Option.Os = runtime.GOOS

	//初始化目录
	err := initAppDir()
	if err != nil {
		return err
	}

	//初始化日志框架
	err = initLog()
	if err != nil {
		return err
	}

	//初始化配置列表
	err = initTask()
	if err != nil {
		return err
	}

	//init command
	initCommand()

	return nil
}

/*
*
初始化 log
flag分发使用自带 log 包
其他使用 logrus
*/
func initLog() error {

	Logger.SetLevel(logrus.InfoLevel)

	// 创建一个Rotating Log的Hook
	logPath := Option.LogPath
	logFileName := "snds.log"

	fullPath := filepath.Join(logPath, logFileName)

	// 按日期分割日志文件，保留7天的日志
	logWriter, _ := rotatelogs.New(
		fullPath+".%Y%m%d",
		rotatelogs.WithLinkName(fullPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	// 创建一个Hook，将日志输出到Rotating Log
	Logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
		},
		&logrus.JSONFormatter{},
	))

	return nil
}

/*
*
根据系统初始化命令行参数
*/
func initCommand() {

	if Option.Os == "darwin" || Option.Os == "linux" {
		unixCommand()
	}

	if Option.Os == "windows" {
		winCommand()
	}

}

func unixCommand() {

	Option.Command.Run = exec.Command(
		"nohup",
		fmt.Sprintf("./%s", Option.RunName),
		"-run",
		"&",
	)
	//c.Run = fmt.Sprintf("nohup ./%s -run >%s.log  &", appName, appName)

}

func winCommand() {

}

func initTask() error {

	//便利所有文件，帅选出 toml 文件并转化

	files, err := os.ReadDir(Option.ConfigPath)
	if err != nil {
		return err
	}

	var tasks []model.Task
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".toml") {

			cfg, err := toml.LoadFile(filepath.Join(Option.ConfigPath, file.Name()))
			if err != nil {
				return err

			}

			var task model.Task
			err = cfg.Unmarshal(&task)
			if err != nil {
				return err

			}

			tasks = append(tasks, task)

		}
	}

	Option.Tasks = tasks

	return nil
}

func initAppDir() error {
	// 获取当前程序的绝对路径
	absPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return err
	}
	Option.RunName = filepath.Base(absPath)

	// 获取当前程序所在的文件夹路径
	currentDir := filepath.Dir(absPath)
	Option.AppPath = currentDir

	// 要创建的目录名称
	appDirName := "snds_data"
	tempDirName := "temp"
	configDirName := "config"
	logDirName := "log"

	// 构建新目录的完整路径
	Option.DataPath = filepath.Join(currentDir, appDirName)
	Option.TempPath = filepath.Join(Option.DataPath, tempDirName)
	Option.ConfigPath = filepath.Join(Option.DataPath, configDirName)
	Option.LogPath = filepath.Join(Option.DataPath, logDirName)

	// 如果目录已存在跳过
	_, err = os.Stat(Option.DataPath)
	if err != nil {
		err = os.Mkdir(Option.DataPath, 0755) // 0755 是目录权限
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(Option.TempPath)
	if err != nil {
		err = os.Mkdir(Option.TempPath, 0755) // 0755 是目录权限
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(Option.ConfigPath)
	if err != nil {
		err = os.Mkdir(Option.ConfigPath, 0755) // 0755 是目录权限
		if err != nil {
			return err
		}
	}

	_, err = os.Stat(Option.ConfigPath)
	if err != nil {
		err = os.Mkdir(Option.LogPath, 0755) // 0755 是目录权限
		if err != nil {
			return err
		}
		return nil
	}

	return nil

}
