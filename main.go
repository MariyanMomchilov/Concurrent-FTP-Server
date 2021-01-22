package main

import "./server"

func main() {
	server := server.NewServer("tcp", "localhost", "8000", false)
	server.Serve()
}
