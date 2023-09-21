package option

import (
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
	"runtime"
	"snds/model"
	"strings"
)

var Option = new(model.Option)

func Init() error {

	Option.Os = runtime.GOOS

	//初始化目录
	err := initAppDir()
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
根据系统初始化命令行参数
*/
func initCommand() {

	if Option.Os == "darwin" || Option.Os == "linux" {
		Option.Command.UnixCommand()
	}

	if Option.Os == "windows" {
		Option.Command.WinCommand()
	}

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
	if err == nil {
		return err
	}

	err = os.Mkdir(Option.DataPath, 0755) // 0755 是目录权限
	if err != nil {
		return err
	}

	err = os.Mkdir(Option.TempPath, 0755) // 0755 是目录权限
	if err != nil {
		return err
	}
	err = os.Mkdir(Option.ConfigPath, 0755) // 0755 是目录权限
	if err != nil {
		return err
	}

	err = os.Mkdir(Option.LogPath, 0755) // 0755 是目录权限
	if err != nil {
		return err
	}
	return nil
}
