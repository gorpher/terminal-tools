package main

import (
	"os/exec"
	"syscall"
)

func XCopy(s, d string) error {
	cmd := exec.Command("xcopy", s, d, "/s/y/q")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := cmd.Output()
	return err
}
