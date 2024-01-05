package operation

import (
	"fmt"
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

	for i := 0; i < len(steps); i++ {
		step := steps[i]

		switch step {
		case "Name":
			err := name(&newTask)
			if err != nil {
				return
			}
			break

		}

	}

}

func name(task *model.Task) error {

	var input string

	for {
		fmt.Print("请输入名称：")
		_, err := fmt.Scan(&input)
		if err != nil {
			return err
		}

		fmt.Print(input)

	}

}
