package logger

import (
	"bufio"
	"io"
	"log"
	"net"
	"oluwoye/internal/parser"
	"os"
)

func StartLogger() {
	listener, listener_error := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if listener_error != nil {
		panic("Logger error: " + listener_error.Error())
	}

	defer listener.Close()
	log.Println("[LOGGER STARTED] port:" + os.Getenv("PORT"))

	for {
		socket, socket_error := listener.Accept()
		if socket_error != nil {
			log.Println("Socket Error: ", socket_error.Error())
		}

		go handle_socket(socket)
	}
}

func handle_socket(socket net.Conn) {
	log.Println("New connection accepted: ", socket.RemoteAddr().String())
	reader := bufio.NewReader(socket)
	for {
		message, message_error := io.ReadAll(reader)
		if message_error != nil {
			if message_error == io.EOF {
				break
			}
			log.Println("Error on getting message: ", message_error.Error())
			break
		}
		if string(message) == "" {
			break
		}

		log_message, log_message_error := parser.Parse(string(message))
		if log_message_error != nil {
			log.Println("Error on parsing log: ", log_message_error.Error())
			break
		}
		log.Println(log_message)
	}
}
