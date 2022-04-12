package main

import (
	"flag"
	"fmt"
	"os"
	"wuxia"
)

func printUsage() {
	fmt.Println("copy [source directory or file] [dest directory or file]")
}

func main() {
	// check the number of args provided
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	flag.Parse() // get the source and destination directory

	source_dir := flag.Arg(0) // get the source directory from 1st argument

	dest_dir := flag.Arg(1) // get the destination directory from the 2nd argument

	fmt.Println("Source :" + source_dir)

	// check if the source dir exist
	src, err := os.Stat(source_dir)
	if err != nil {
		panic(err)
	}

	if !src.IsDir() {
		fmt.Println("Source is not a directory")
		os.Exit(1)
	}

	// create the destination directory
	fmt.Println("Destination :" + dest_dir)

	_, err = os.Open(dest_dir)
	if !os.IsNotExist(err) {
		fmt.Println("Destination directory already exists. Abort!")
		os.Exit(1)
	}

	err = wuxia.CopyDir(source_dir, dest_dir)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" Directory copied")
	}

}
