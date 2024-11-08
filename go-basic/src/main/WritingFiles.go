package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkw(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/Users/spirit/Downloads/text.txt", d1, 0644)
	checkw(err)

	f, err := os.Create("/Users/spirit/Downloads/text.txt")
	checkw(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkw(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	checkw(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	checkw(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
