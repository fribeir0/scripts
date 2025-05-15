package main

import (
	"fmt"
	"os/exec"

)

func ExecCommands (command string, args ...string) ([]byte,error){
	cmd:= exec.Command(command,args...)
	stdout, err := cmd.Output()
	if err != nil {
		return nil,err
	}
	return stdout, nil 

}

func main () {
	fmt.Println("Iniciando reconhecimento...")
	hname, err := ExecCommands("hostname")
	if err != nil {
		fmt.Println("error",err)
	}
	fmt.Println("hostname",string(hname))
	dir, err := ExecCommands("ls","-la","/")
	if err != nil {
		fmt.Println("erro ao listar:",err)
	}
		fmt.Println("diretorios em /",string(dir))
	openP, err:= ExecCommands("ss","-tuln")
	if err != nil {
		fmt.Println("erro ao ver portas abertas",err)
	}
		fmt.Println("portas abertas:",openP)
	psAux,err:= ExecCommands("psaux")
	if err != nil {
		fmt.Println("erro ao listar processos:",err)
	}
	fmt.Println("process:",psAux)

}