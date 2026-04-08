package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Client sends commands to the Redis server as RESP arrays
// "*3\r\n$3\r\nSET\r\n$4\r\nuser\r\n$5\r\nalice"
// ["*3", "$3SET", "$4user", "$5alice"]

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if len(line) < 2 {
			break
		}

		line = strings.TrimSuffix(line, "\r\n")
		if line[0] == '*' {
			arrNum := strings.TrimPrefix(line, "*")
			count, err := strconv.Atoi(arrNum)
			if err != nil {
				break
			}
		}
	}
}

func main() {
	fmt.Println("Logs from your program will appear here!")
	
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}
