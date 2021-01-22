package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Server represents FTP server based on TCP by default
type Server struct {
	listener net.Listener
	auth     bool
}

// NewServer contructs Servers
func NewServer(connection, address, port string, auth bool) *Server {
	listener, err := net.Listen(connection, address+": "+port)
	if err != nil {
		log.Fatalf("Failed to create new Server: %s", err)
	}

	return &Server{listener, auth}
}

// Serve starts the server
func (s *Server) Serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Printf("Error occured when conecting: %s", err)
			continue
		}

		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()

	buff := bufio.NewScanner(conn)

	for buff.Scan() {
		requestStr := strings.TrimSpace(buff.Text())

		if requestStr != "" {
			whitespaceIndex := strings.IndexByte(requestStr, ' ')
			var cmd string
			if whitespaceIndex != -1 {
				cmd = requestStr[:strings.IndexByte(requestStr, ' ')]
			} else {
				cmd = requestStr
			}
			rest := strings.TrimSpace(requestStr[len(cmd):])
			req := s.apply(cmd, rest)
			if req != nil {
				req.apply(conn)
			}
		}
	}
}

func (s *Server) apply(cmd, rest string) cmdRequest {
	switch cmd {
	case "get":
		return &getFileRequest{rest}
	case "ls":
		return &listRequest{rest}
	}
	return nil
}
