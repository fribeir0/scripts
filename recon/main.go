package main

import (
	"fmt"
	"os/exec"

)

func ExecCommands (command string, args ...string) ([]byte,error){
	cmd:= exec.Command(command,args...)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return stdout, err 

}

func main () {
	fmt.Println("Iniciando reconhecimento...")
	hname, err := ExecCommands("hostname")
	fmt.Println("hostname:",hname,err.Error())
	dir, err := ExecCommands("ls","-la","/")
	fmt.Println("diretorios:",dir,err.Error())
	openP, err:= ExecCommands("ss","-tuln")
	fmt.Println("open ports",openP,err.Error())
	psAux,err:= ExecCommands("psaux")
	fmt.Println("process:",psAux,err.Error())

}