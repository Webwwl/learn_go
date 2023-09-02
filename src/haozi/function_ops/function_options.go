package main

import (
	"fmt"
	"time"
)

// Ip、Port 必填, 其他非必填
type Server struct {
	Ip       string
	Port     int
	Timeout  time.Duration
	Protocol string
	MaxConns int
}

func (s *Server) LogInfo() {
	fmt.Printf("%v\n", s)
}

func main() {
	s1 := createServer1()
	s1.LogInfo()
	s2 := createServer2()
	s2.LogInfo()
	s3 := createServer3()
	s3.LogInfo()
}

type Server1 struct {
	Ip     string
	Port   int
	Config *ServerConfig
}

func (s *Server1) LogInfo() {
	fmt.Printf("ip: %s\nport: %d\nConfig: %v\n", s.Ip, s.Port, s.Config)
}

type ServerConfig struct {
	Timeout  time.Duration
	Protocol string
	MaxConns int
}

func createServer1() Server1 {
	// conf := &ServerConfig{Timeout: time.Second * 20, Protocol: "https", MaxConns: 10}
	// return _createServer1("127.0.0.1", 8080, conf)
	return _createServer1("127.0.0.1", 8080, nil)
}

func _createServer1(ip string, port int, config *ServerConfig) Server1 {
	defConfig := ServerConfig{time.Second * 10, "https", 5}
	if config == nil {
		return Server1{ip, port, &defConfig}
	}
	return Server1{ip, port, config}
}

// 一下是链式调用的实现
type ServerBuild struct {
	server Server
}

func (sb *ServerBuild) create(ip string, port int) *ServerBuild {
	sb.server.Ip = ip
	sb.server.Port = port
	return sb
}

func (sb *ServerBuild) withTimeout(timeout time.Duration) *ServerBuild {
	sb.server.Timeout = timeout
	return sb
}

func (sb *ServerBuild) withProtocol(protocol string) *ServerBuild {
	sb.server.Protocol = protocol
	return sb
}

func (sb *ServerBuild) withMaxConns(maxConns int) *ServerBuild {
	sb.server.MaxConns = maxConns
	return sb
}

func (sb *ServerBuild) build() Server {
	return sb.server
}

func createServer2() Server {
	sb := ServerBuild{}
	s := sb.create("localhost", 9999).withTimeout(time.Second * 100).withProtocol("http").withMaxConns(2).build()
	return s
}

// Functional Options 实现方式
type Option func(s *Server)

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func createServer3() Server {
	return _createServer3("127.0.0.1", 8888, Timeout(time.Second*1))
}

func _createServer3(ip string, port int, options ...func(s *Server)) Server {
	s := Server{ip, port, time.Second * 10, "https", 20}
	for _, fn := range options {
		fn(&s)
	}
	return s
}
