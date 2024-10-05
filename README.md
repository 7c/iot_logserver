# iot_logserver

Have you been using Serial.print() or Serial.println() in your Arduino code to debug? It is good way to debug if your IOT (esp32/arduino/other) is attached to your computer but when it is not you need another way to debug or assure that your IOT is working as expected. Lightweight logging is what you are looking for. UDP over WiFi comes handy and this project launches a UDP server you can send UDP packets to be logged to console/files.


## run with arguments
```
go run *.go -udpport 12345

// 12345 is the default port if you do not specify one or .env file is not found.
```

## run with env variables
create a .env file inside the same folder as the executable with the following:
```env
UDPPORT=12345
```
then run:
```
go run *.go
```
