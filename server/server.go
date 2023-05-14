package main

import (
	"fmt"
	"net"
	//"golang.org/x/text/message"
)

func main() {

	// inicio el servidor
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// Clientes: for para recibir todos los clientes
	for {
		fmt.Println("Esperando un cliente")

		//Aceptamos la conexión de un cliente
		client, err := server.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Un cliente se conectó")

		//funcion para manejar nuestro cliente
		go managerConnection(client)
	}

}

func managerConnection(Client net.Conn) {

	userConnection = append(userConnection, Client)
	for {
		//Leemos el mensaje
		var buff = make([]byte, 2048)

		_, err := Client.Read(buff)
		if err != nil {
			panic(err)
		}

		//Se lo mandamos a todos los users conectados
		writeMessageAllUsers(buff)

	}

}

func writeMessageAllUsers(message []byte) {
	for _, c := range userConnection {
		c.Write(message)
	}
}
