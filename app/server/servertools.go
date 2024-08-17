package server

import (
	"net"
)

func GetFreeTcpPort() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return -1, err
	}

	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}
