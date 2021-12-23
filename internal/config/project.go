package config

import (
	"os"
	"os/exec"
	"strings"
)

func (p *Project) Reload() error {
	for _, command := range p.Commands {
		splittedCommand := strings.Split(command.Command, " ")
		cmd := exec.Command(splittedCommand[0], splittedCommand[1:]...)
		cmd.Dir = command.WorkDir
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
