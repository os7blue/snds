package model

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"snds/option"
)

type Command struct {
	Run  *exec.Cmd
	Stop *exec.Cmd
}

func (c *Command) UnixCommand() {

	c.Run = exec.Command(
		"nohup",
		fmt.Sprintf("./%s", option.Option.RunName),
		"-run",
		fmt.Sprintf(">%s.log", filepath.Join(option.Option.LogPath, option.Option.RunName)),
		"&",
	)
	//c.Run = fmt.Sprintf("nohup ./%s -run >%s.log  &", appName, appName)

}

func (c *Command) WinCommand() {

}
