package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Init app")

	//Cliente se conecta al server
	server, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	//escucha al servidor
	// Solicita el nombre de usuario
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese su nombre de usuario: ")
	scanner.Scan()
	username := scanner.Text()

	//imprime mensajes
	go readMessages(server, username)

	//envia mensajes
	writeMessage(server, username)

}

func writeMessage(conn net.Conn, username string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		fullMessage := fmt.Sprintf("%s: %s\n", username, text)
		_, err := conn.Write([]byte(fullMessage))
		if err != nil {
			panic(err)
		}
	}
}

func readMessages(conn net.Conn, username string) {
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		message := string(buffer)
		parts := strings.SplitN(message, ": ", 2)
		if len(parts) < 2 {
			continue
		}

		sender := parts[0]
		text := parts[1]

		if sender != username {
			fmt.Printf("[%s]: %s", sender, text)
		}
	}
}
