package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

for{
	conn,err:=l.Accept()
	if err!=nil{
		fmt.Println("Error accepting connection: ",err.Error())
		// go handleConnection(conn)
		
	}
	go handleConnection(conn)

 }
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	reader:= bufio.NewScanner(conn)
	for reader.Scan(){
		text:=reader.Text()
		if strings.TrimSpace(text)=="PING"{
			conn.Write([]byte("+PONG\r\n"))
		}
	}
	 _,err := conn.Write([]byte("+PONG\r\n"))
	if err != nil{
fmt.Print("Error writing to connection",err)
return
	}
}
	
	


	


