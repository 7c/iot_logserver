package main

import (
	"log"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/joho/godotenv"
)

var (
	app                 = kingpin.New("iot_logserver", "IoT Log Server")
	udpListenPortString = app.Flag("udpport", "UDP listen port.").Default("12345").String()
)

func ParseEnv() bool {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
		return false
	}

	log.Println("loading .env file")
	*udpListenPortString = os.Getenv("UDPPORT")
	log.Println(".env file loaded")
	return true
}

func ParseArgs() {
	log.Println("parsing args")
	kingpin.MustParse(app.Parse(os.Args[1:]))

}
