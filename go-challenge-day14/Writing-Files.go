package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	checkerr(err)

	f, err := os.Create("/tmp/dat2")
	checkerr(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkerr(err)
	fmt.Printf("Wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	checkerr(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	checkerr(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
