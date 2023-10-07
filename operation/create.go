package operation

import (
	"fmt"
	"snds/model"
)

type create struct {
}

// create new config.go
func (c *create) Create(name string) {

	newTask := model.Task{
		Token:     "",
		Path:      "",
		LocalPath: nil,
		Name:      name,
		Key:       "",
		Type:      "",
		Cron:      "",
	}

	newTask.LocalPath = []string{""}

	fmt.Println(newTask)

}
