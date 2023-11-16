package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmt := exec.Command("date")

	dataOutput, err := dateCmt.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dataOutput))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing", err)
		case *exec.ExitError:
			fmt.Println("Command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write([]byte("hello grep \n goodbye grep"))

	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep Hello")
	fmt.Println(string(grepBytes))

	IsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	IsOut, err := IsCmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(IsOut))
}
