package wuxia

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"os"
	"time"
)

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	// get source size
	var sourceSize int64
	sourceStat, err := sourcefile.Stat()
	if err != nil {
		fmt.Printf("Can't stat %s: %v\n", sourcefile, err)
		return
	}
	sourceSize = sourceStat.Size()

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	// create bar
	bar := pb.New(int(sourceSize)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = false
	bar.ShowPercent = true
	bar.ShowTimeLeft = true
	bar.Start()

	// create multi writer
	writer := io.MultiWriter(destfile, bar)

	_, err = io.Copy(writer, sourcefile)

	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()
		srcFileName := obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			fmt.Println("")
			fmt.Println("FileName :" + srcFileName) //--------------------------------------------------<<<
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}
