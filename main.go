package main

import (
	"iot_logserver/udpserver"
	"log"
	"strconv"
)

var logDebug *log.Logger = log.Default()
var listenPort int

func main() {

	if !ParseEnv() {
		ParseArgs()
	}
	if _listenPort, err := strconv.Atoi(*udpListenPortString); err != nil {
		log.Fatal(err)
	} else {
		listenPort = _listenPort
		logDebug.Println("listen port :", listenPort)
	}
	ListInterfaceIPs()

	udpserver := udpserver.NewUDPServer(listenPort)
	go udpserver.Start()

	select {}
}
