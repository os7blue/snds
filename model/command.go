package model

import (
	"os/exec"
)

type Command struct {
	Run  *exec.Cmd
	Stop *exec.Cmd
}
