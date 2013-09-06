package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const version = "0.1"

func main() {
	var bind, logfile string
	var showVer bool

	flag.StringVar(&bind, "b", "0.0.0.0:8080", "Address to bind")
	flag.StringVar(&logfile, "l", "stderr", "Log file. Defaults to stderr")
	flag.BoolVar(&showVer, "v", false, "Show version and exit")
	flag.Parse()

	if showVer {
		fmt.Fprintf(os.Stderr, "%s %s\n", os.Args[0], version)
		return
	}

	if logfile != "stderr" {
		f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, "open logfile ", err.Error())
			os.Exit(1)
		}
		log.SetOutput(f)
		defer f.Close()
	}

	log.Print("Listening on ", bind)

	ln, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatal("listen ", err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("accept ", err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	defer conn.Close()

	who := conn.RemoteAddr().String()
	msg := who + "\n"
	n, err := conn.Write([]byte(msg))

	if err != nil {
		log.Print("write ", err.Error())
		return
	}

	if n != len(msg) {
		log.Printf("incomplete write %d of %d (%#v)", n, len(msg), msg)
		return
	}

	log.Print("sent ", who)
}
