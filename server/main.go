package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ccding/go-stun/stun"
)

func main() {
	var server string
	flag.StringVar(&server, "s", "", "Server address")
	flag.Parse()

	client := stun.NewClient()
	if server != "" {
		client.SetServerAddr(server)
	}

	nat, host, err := client.Discover()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("NAT Type : %s\n", nat)
	fmt.Printf("Host     : %s\n", host)

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", host.Port()))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		buf := make([]byte, 1024)
	LOOP:
		for {
			size, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				log.Print(err)
				break LOOP
			}
			fmt.Printf("Addr : %v\n", addr)
			fmt.Println(hex.Dump(buf[:size]))
		}
	}()

	<-ch
}
