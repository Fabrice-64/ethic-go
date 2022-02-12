package main

import (
	"log"
	"os"
	"os/exec"
)

func ExecCmd(cmd string, args_arr []string) (err error) {
	args := args_arr
	cmd_obj := exec.Command(cmd, args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	err = cmd_obj.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	return nil

}

func main() {
	cmd := "ls"
	ExecCmd(cmd, []string{"-l"})

}
