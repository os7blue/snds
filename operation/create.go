package operation

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
	"snds/model"
	"snds/option"
	"snds/util"
)

type create struct {
}

// Create new config.go
func (c *create) Create(taskName string) error {

	newTask := model.Task{
		Token:     "",
		Path:      "you are remote path",
		LocalPath: []string{"you are path"},
		Key:       "",
		Type:      "ALI,BAIDU,GOOGLE",
		Cron:      "",
	}
	yes := newTask.NameExist(option.Option.Tasks, taskName)
	if yes {
		return errors.New("已存在名为 %s 的任务配置")
	}

	key, err := util.GlobalUtil.GenerateRandomHex(16)
	if err != nil {
		return err
	}
	newTask.Key = key

	tomlData, err := toml.Marshal(newTask)
	if err != nil {
		return err
	}

	// 保存Toml数据到文件
	file, err := os.Create(fmt.Sprintf("%s.toml", filepath.Join(option.Option.ConfigPath, taskName)))

	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.Write(tomlData)
	if err != nil {
		return err
	}

	return nil
}
