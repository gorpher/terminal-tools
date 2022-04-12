package main

import (
	"flag"
	"fmt"
	"github.com/gorpher/gone"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"wuxia"
)

func printUsage() {
	fmt.Println("start_chrome [chrome安装位置] [工作路径]\n比如：start_chrome.exe \"C:/Program Files/Google/Chrome/Application/chrome.exe\"  \"D:/Google\" http://www.baidu.com  ")
}

// windows下chrome程序多开
func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}
	flag.Parse()
	args := flag.Args()
	chrome := flag.Arg(0)
	workDir := flag.Arg(1)
	uuid := GetSID()
	if uuid == "" {
		uuid = gone.ID.SString()
	}
	homePath := os.Getenv("homepath")
	fmt.Println(filepath.Clean(homePath))

	defaultDir := filepath.Join(workDir, "default_dir")
	if _, err := os.Stat(defaultDir); err != nil {
		os.MkdirAll(defaultDir, os.ModePerm)
	}
	userDataDir := filepath.Join(workDir, "user_data_dir", filepath.Clean(homePath), uuid)
	err := wuxia.CopyDir(defaultDir, userDataDir)
	if err != nil {
		log.Fatalln(err)
	}
	targs := []string{fmt.Sprintf("--user-data-dir=%s", userDataDir)}
	targs = append(targs, args[2:]...)
	err = exec.Command(chrome, targs...).Start()
	if err != nil {
		log.Fatalln(err)
	}
}
