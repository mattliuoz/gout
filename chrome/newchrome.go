package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {

	cmdName := "cmd"

	rand.Seed(time.Now().UTC().UnixNano())

	path := fmt.Sprintf("%s%d%s", "c:\\temp\\chrome\\", rand.Intn(100), "\\")

	fmt.Println(path)

	os.MkdirAll(path, 0777)

	cmdArgs := []string{"/c", "start", "chrome", "/new-window", "--user-data-dir=" + path}

	cmd := exec.Command(cmdName, append(cmdArgs, os.Args[1])...)
	cmd.Start()
	fmt.Println("Done:")
}
