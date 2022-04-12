package main

import (
	"bytes"
	"flag"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	method := flag.String("m", "", "请求方式： download,upload,mkdir,delete,login")
	u := flag.String("u", "anonymous", "用户名")
	p := flag.String("p", "anonymous", "密码")
	h := flag.String("h", "ftp.example.org:21", "主机地址")
	flag.Parse()
	args := flag.Args()
	switch *method {
	case "mkdir":
		if len(args) != 1 {
			log.Fatal("gftp -m mkdir -u anonymous -p anonymous -h ftp.example.org:21 <dirpath>")
		}
		MKDIR(*u, *p, *h, args[0])
		return
	case "upload":
		if len(args) != 2 {
			log.Fatal("gftp -m upload -u anonymous -p anonymous -h ftp.example.org:21 <source> <target>")
		}
		Upload(*u, *p, *h, args[0], args[1])
		return
	case "download":
		if len(args) != 2 {
			log.Fatal("gftp -m download -u anonymous -p anonymous -h ftp.example.org:21 <source> <target>")
		}
		Download(*u, *p, *h, args[0], args[1])
		return
	case "delete":
		if len(args) != 1 {
			log.Fatal("gftp -m delete -u anonymous -p anonymous -h ftp.example.org:21 <dirpath>")
		}
		Delete(*u, *p, *h, args[0])
		return
	case "login":
		Login(*u, *p, *h)
		return
	}
	flag.PrintDefaults()
	os.Exit(1)
}

func Login(u, p, host string) {
	c, err := ftp.Dial(host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(u, p)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("登录成功")
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
func MKDIR(u, p, host, sourceFile string) {
	if sourceFile == "" {
		log.Fatal("源路径不存在")
	}
	c, err := ftp.Dial(host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(u, p)
	if err != nil {
		log.Fatal(err)
	}

	err = c.MakeDir(sourceFile)
	if err != nil {
		log.Println(err)
		return
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
func Delete(u, p, host, sourceFile string) {
	if sourceFile == "" {
		log.Fatal("源路径不存在")
	}
	c, err := ftp.Dial(host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(u, p)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Delete(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

func Upload(u, p, host, sourceFile, targetFile string) {
	if sourceFile == "" {
		log.Fatal("源路径不存在")
	}
	if targetFile == "" {
		log.Fatal("目标路径不存在")
	}
	body, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	c, err := ftp.Dial(host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(u, p)
	if err != nil {
		log.Fatal(err)
	}
	//c.MakeDir(filepath.Dir(targetFile))
	// don't care error

	err = c.Stor(targetFile, bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

func Download(u, p, host, sourceFile, targetFile string) {
	if sourceFile == "" {
		log.Fatal("源路径不存在")
	}
	if targetFile == "" {
		log.Fatal("目标路径不存在")
	}
	c, err := ftp.Dial(host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(u, p)
	if err != nil {
		log.Fatal(err)
	}
	//dir, filename := filepath.Split(sourceFile)
	//err = c.ChangeDir(dir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("cd %s \n", dir)
	r, err := c.Retr(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(targetFile, body, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}

}
