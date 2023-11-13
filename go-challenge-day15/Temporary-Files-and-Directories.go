package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkerrorSecond(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	checkerrorSecond(err)
	fmt.Println("temp file name", f.Name())

	defer os.Remove(f.Name())
	_, err = f.Write([]byte{1, 2, 3, 4})
	checkerrorSecond(err)

	dname, err := os.MkdirTemp("", "samples")
	checkerrorSecond(err)
	fmt.Println("Temp dir name", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	checkerrorSecond(err)
}
