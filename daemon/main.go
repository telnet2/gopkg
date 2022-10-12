package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

// StartServer listens from the addr and waits for the connection.
func StartServer(name string, tcpListener *net.TCPListener, args ...string) error {
	proc := exec.Command("./serverApp/serverApp", args...)

	fd, err := tcpListener.File()
	if err != nil {
		return err
	}
	proc.ExtraFiles = []*os.File{fd}
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	log.Printf("starting server: %s", name)
	return proc.Run()
}

// ListenAddr opens up a listener on the addr
func ListenAddr(addr string) (*net.TCPListener, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	tcpListener, ok := listener.(*net.TCPListener)
	if ok {
		return tcpListener, nil
	}

	return nil, errors.New("fail to create TCPListener")
}

func main() {
	l, err := ListenAddr(":9999")
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			err = StartServer(":9999", l, "-name", fmt.Sprintf("service-%d", i))
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(i)
	}

	time.AfterFunc(time.Second*2, func() {
		for i := 0; i < 10; i++ {
			go func(i int) {
				for j := 0; j < 10; j++ {
					res, err := http.Get(fmt.Sprintf("http://localhost:9999/hi?id=%d", i))
					if err != nil {
						log.Println(err)
						continue
					}
					body, _ := io.ReadAll(res.Body)
					fmt.Println(string(body))
				}
			}(i)
		}
	})

	wg.Wait()
}
