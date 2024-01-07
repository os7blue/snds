package operation

import (
	"snds/model"
)

type create struct {
}

var steps = []string{"Name", "Key"}

// Create new config.go
func (c *create) Create() {

	newTask := model.Task{
		Token:     "",
		Path:      "you are remote path",
		LocalPath: []string{"you are path"},
		Name:      "",
		Key:       "",
		Type:      "",
		Cron:      "",
	}

}
