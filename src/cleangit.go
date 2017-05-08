package main

import "os"
import "fmt"
import "os/exec"

func main(){
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Should have a least one git branch specified. e.g. cleangit testbranch")
	}
	
	var (
		cmdOut []byte
		err    error
	)

	cmdName :="git"
	
	for i:=1; i <len(os.Args); i++ {
		cmdArgs := []string{"branch", "-D"}
		cmdArgs = append(cmdArgs,os.Args[i])
		if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error while deleting branch:{"+os.Args[i]+"} ", err)
			os.Exit(1)
		}
	sha := string(cmdOut)
	firstSix := sha[:6]
	fmt.Println("The first six chars of the SHA at HEAD in this repo are", firstSix)
	}
}