package main

import "fmt"

type (
	ServerOptions func(options) options
	TransportType int
)

const (
	UDP TransportType = iota
	TCP
)

type Server struct {
	options
	isAlive bool
}

type options struct {
	MaxConnection int
	TransportType TransportType
	Name          string
}

func MaxConnection(n int) ServerOptions {
	return func(o options) options {
		o.MaxConnection = n
		return o
	}
}
func ServerName(n string) ServerOptions {
	return func(o options) options {
		o.Name = n
		return o
	}
}
func Transport(t TransportType) ServerOptions {
	return func(o options) options {
		o.TransportType = t
		return o
	}
}
func NewServer(opts ...ServerOptions) Server {
	options := options{
		TransportType: TCP,
	}
	for _, opt := range opts {
		options = opt(options)
	}
	return Server{
		options: options,
		isAlive: true,
	}
}

func main() {
	server := NewServer(MaxConnection(10), ServerName("test"))
	fmt.Printf("%+v\n", server)
}
