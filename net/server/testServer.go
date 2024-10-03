// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// )

// func handleconnection(conn net.Conn) {
// 	//read msg from client
// 	for {
// 		reader, _ := bufio.NewReader(conn).ReadString('\n')
// 		fmt.Print("message from client: ")
// 		fmt.Println(reader)
// 		//send msg ok back to client
// 		fmt.Fprintln(conn, "ok")
// 	}
// }

// func main() {
// 	ln, _ := net.Listen("tcp", ":8080")
// 	for {
// 		conn, _ := ln.Accept()
// 		fmt.Println("Successfully connected with client!")
// 		fmt.Println("Waiting for message...")
// 		go handleconnection(conn)

// 	}
// }
