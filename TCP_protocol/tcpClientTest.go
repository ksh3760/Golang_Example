package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	G_PORT string = ":8081"
)

func main() {
	var (
		sStr string
		aCon net.Conn
		sErr error
		sNum int
	)

	aCon, sErr = net.Dial("tcp", G_PORT)
	if nil != sErr {
		log.Println(sErr)
	}

	log.Println("port : ", G_PORT)

	go func() {
		data := make([]byte, 4096)

		for {
			sNum, sErr = aCon.Read(data)
			if sErr != nil {
				log.Println(sErr)
				return
			}

			log.Println("Server send : " + string(data[:sNum]))
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	for {
		fmt.Print("msg >")
		fmt.Scanln(&sStr)
		aCon.Write([]byte(sStr))
		time.Sleep(time.Duration(1) * time.Second)
	}
}
