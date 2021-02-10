package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	connection, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting Client on :9000")

	defer connection.Close()

	fmt.Println("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	username = strings.Trim(username, "\r\n")
	fmt.Printf("Welcome %s ! write Below to send Messages :-", username)

	go Read(connection)
	Write(connection, username)

}

func Read(connection net.Conn) {
	for {
		reader := bufio.NewReader(connection)
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(message)

	}
}

func Write(connection net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		message = fmt.Sprintf("%s:-%s\n", username, strings.Trim(message, "\r\n"))
		connection.Write([]byte(message))

	}
}
