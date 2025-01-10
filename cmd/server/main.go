package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	host, err := net.DialTCP("tcp", nil, nil)
	temp := host.LocalAddr().String()
	fmt.Printf(temp, err)
	fmt.Printf("Starting listening on port %s ...", port)
}
