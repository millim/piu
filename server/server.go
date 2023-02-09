package server

import (
	"flag"
	"fmt"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) InitArgs() {
	flag.StringVar(&s.Host, "host", "0.0.0.0", "server host")
	flag.IntVar(&s.Port, "port", 8080, "server port")
}

func (s *Server) Run() {
	fmt.Println("hello world")
}
