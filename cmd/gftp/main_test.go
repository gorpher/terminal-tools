package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestDownload(t *testing.T) {
	getwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(getwd)
	err = exec.Command("go build -o gftp main.go").Run()
	if err != nil {
		t.Fatal(err)
	}

}
