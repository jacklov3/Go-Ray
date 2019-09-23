package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){
	conn,err :=net.Dial("tcp","www.google.com:80")
	if err!=nil{
		fmt.Println("can't connect to the server")
	}
	fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
	status,err :=bufio.NewReader(conn).ReadString('\n')
	if err!=nil{
		fmt.Println("something wrong!")
	}
	fmt.Println(status);
}