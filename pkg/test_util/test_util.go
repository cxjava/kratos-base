package test_util

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"
	"time"
)

// OpenPort returns an open port.
func OpenPort() int {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	u, err := url.Parse("//" + l.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	if u.Port() == "" {
		log.Fatalf("unable to parse a port from %s", l.Addr())
	}

	p, err := strconv.Atoi(u.Port())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(p)
	return p
}
// CheckServerIsReady check server is ready and close done chan.
func CheckServerIsReady(port int, done chan struct{}) {
	timeout := time.After(10 * time.Second)
	defer func ()  {
		if done != nil {
			close(done)
		}
	}()
	for {
		select {
		case <-timeout:
			log.Println("Time out on port ", port)
			return
		default:
			if isServerRunning(port) {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
// isServerRunning check server is running or not.
func isServerRunning(port int) bool {
	_, err := net.Dial("tcp", "localhost:"+fmt.Sprintf("%d", port))
	if err != nil {
		return false
	}
	log.Println("Server is running on port", port)
	return true
}
