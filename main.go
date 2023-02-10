package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"piu/client"
	"piu/comm"
	"piu/server"
)

func RunAsDaemon(args []string) {
	_args := make([]string, 0)
	for _, v := range args {
		if v == "-d" || v == "--d" {
			continue
		}
		_args = append(_args, v)

	}
	cmd := exec.Command(_args[0], _args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Process started with PID:", cmd.Process.Pid)
	}
	os.Exit(0)
}

var service ServiceImp

func main() {
	daemon := flag.Bool("d", false, "is daemon")
	version := flag.Bool("version", false, "is daemon")
	if len(os.Args) <= 1 || (os.Args[1] != "server" && os.Args[1] != "client") {
		fmt.Println("need input: main server or main client")
		os.Exit(0)
	}
	switch os.Args[1] {
	case "server":
		service = &server.Server{}
	case "client":
		service = &client.Client{}
	}
	service.InitArgs()
	tempArgs := os.Args
	_args := []string{os.Args[0]}
	_args = append(_args, os.Args[2:]...)
	os.Args = _args
	flag.Parse()
	if version != nil && *version {
		fmt.Println(fmt.Sprintf("piu version is %s", comm.Version))
		os.Exit(0)
		return
	}

	if daemon != nil && *daemon {
		RunAsDaemon(tempArgs)
		os.Exit(0)
		return
	}
	service.Run()
}

type ServiceImp interface {
	InitArgs()
	Run()
}
