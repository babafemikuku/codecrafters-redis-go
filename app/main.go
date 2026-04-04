package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Client sends commands to the Redis server as RESP arrays
// "*3\r\n$3\r\nSET\r\n$4\r\nuser\r\n$5\r\nalice"
// ["*3", "$3SET", "$4user", "$5alice"]

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()

		withoutcrlf := strings.Split(line, "\r\n")

		arguments := []string{}
		for i, str := range withoutcrlf[1:] {
			arguments[i] = strings.ReplaceAll(str, "$", "")
		}

		if strings.ToUpper(arguments[0]) == "ECHO" {
			strLen := len(arguments[1])
			resp := fmt.Sprintf("$%d\r\n%s\r\n", strLen, arguments[1])
			conn.Write([]byte(resp))
		}

		conn.Write([]byte("+PONG\r\n"))
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
