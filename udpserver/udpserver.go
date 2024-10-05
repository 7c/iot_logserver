package udpserver

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
)

var logger *log.Logger = log.New(os.Stdout, color.CyanString("UDPLOGGER "), log.LstdFlags|log.Lmicroseconds)

type UDPServer struct {
	listenPort int
	udpConn    *net.UDPConn
}

func NewUDPServer(listenPort int) *UDPServer {
	return &UDPServer{
		listenPort: listenPort,
	}
}

func (u *UDPServer) Start() {
	logger.Printf("Starting UDP server on port %d", u.listenPort)
	// start udp server
	udpAddr := &net.UDPAddr{
		Port: u.listenPort,
		IP:   net.ParseIP("0.0.0.0"),
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	u.udpConn = udpConn
	defer udpConn.Close()
	logger.Printf("UDP Server started on port %d", u.listenPort)

	// read from udpConn
	buffer := make([]byte, 1024)
	for {
		n, addr, err := udpConn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		// send response
		udpConn.WriteToUDP([]byte("Received\n"), addr)
		// we need to split the buffer into lines
		lines := strings.Split(string(buffer[:n]), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			logger.Printf("%s: %s", color.CyanString(addr.String()), color.BlueString(line))
		}
	}
}

func (u *UDPServer) Stop() {
	logger.Println("Stopping UDP Server")
	if u.udpConn != nil {
		logger.Println("Closing UDP connection")
		u.udpConn.Close()
		logger.Println("UDP connection closed")
	}
}
