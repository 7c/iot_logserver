package udpserver

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var logger *log.Logger = log.New(os.Stdout, color.CyanString("UDPSERVER "), log.LstdFlags|log.Lmicroseconds)

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
	logger.Printf("UDP server started on port %d", u.listenPort)

	// read from udpConn
	buffer := make([]byte, 1024)
	for {
		n, addr, err := udpConn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		// logger.Printf("Received %d bytes from %s: '%s'", n, addr, buffer[:n])
		// send response
		udpConn.WriteToUDP([]byte("Received\n"), addr)
		// we expect line protocol
		// we need to split the buffer into lines
		lines := strings.Split(string(buffer[:n]), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			// u.influx2Writer.WritePoint(line)
			logger.Printf("received line from %s: %s", color.CyanString(addr.String()), color.BlueString(line))
			// lets add microsecond from current time to the line
			line = fmt.Sprintf("%s %d", line, time.Now().UnixMicro())
			logger.Printf("influx > %s", color.GreenString(line))
		}
	}
}

func (u *UDPServer) Stop() {
	logger.Println("Stopping UDP server")
	if u.udpConn != nil {
		logger.Println("Closing UDP connection")
		u.udpConn.Close()
		logger.Println("UDP connection closed")
	}
}
