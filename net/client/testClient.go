 package main

 import (
 	"bufio"
 	"fmt"
 	"net"
 	"os"
 )

 func readMessage(conn net.Conn) {
 	msg, _ := bufio.NewReader(conn).ReadString('\n')
 	fmt.Println(msg)
 }
 func main() {
 	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
 	stdin := bufio.NewReader(os.Stdin)
 	for {
 		//read input from terminal and send msg to server
 		fmt.Print("Enter text: ")
 		msg, _ := stdin.ReadString('\n')
 		fmt.Fprintln(conn, msg)
 		//reveive msg from server
 		fmt.Print("response from server: ")
 		readMessage(conn)
 	}
 }
