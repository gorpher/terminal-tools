package main

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

func GetSID() string {
	reader := &bytes.Buffer{}
	cmd := exec.Command("query", "session")
	cmd.Stdout = reader
	cmd.Run()
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if strings.HasPrefix(line, ">") {
			i := 0
			chars := strings.Split(line, " ")
			for _, c := range chars {
				if c != "" {
					i += 1
				}
				if i == 3 {
					return c
				}
			}
		}
	}
	return ""
}
