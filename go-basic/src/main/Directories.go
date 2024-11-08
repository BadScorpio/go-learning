package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func checkd(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	err := os.Mkdir("/Users/spirit/Downloads/subdir", 0755)
	checkd(err)

	defer os.RemoveAll("/Users/spirit/Downloads/subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		checkd(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("/Users/spirit/Downloads/subdir/file1")

	err = os.MkdirAll("/Users/spirit/Downloads/subdir/parent/child", 0755)
	checkd(err)

	createEmptyFile("/Users/spirit/Downloads/subdir/parent/file2")
	createEmptyFile("/Users/spirit/Downloads/subdir/parent/file3")
	createEmptyFile("/Users/spirit/Downloads/subdir/parent/child/file4")

	c, err := os.ReadDir("/Users/spirit/Downloads/subdir/parent")
	checkd(err)

	fmt.Println("Listing /Users/spirit/Downloads/subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("/Users/spirit/Downloads/subdir/parent/child")
	checkd(err)

	c, err = os.ReadDir(".")
	checkd(err)

	fmt.Println("Listing /Users/spirit/Downloads/subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	checkd(err)

	fmt.Println("Visiting /Users/spirit/Downloads/subdir")
	err = filepath.WalkDir("/Users/spirit/Downloads/subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
