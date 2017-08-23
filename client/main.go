package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	addr := os.Args[1]

	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		fmt.Print("> ")
		var msg string
		_, err := fmt.Scanf("%s\n", &msg)
		if err != nil {
			log.Print(err)
			break
		}
		if msg == "exit" {
			break
		}
		io.WriteString(conn, msg)
	}

}
