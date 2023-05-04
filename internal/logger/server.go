package logger

import (
	"bufio"
	"io"
	"log"
	"net"
	"oluwoye/internal/config"
	"oluwoye/internal/parser"
	"os"
)

func StartLogger() {
	config_file, load_error := config.LoadConfigFile()
	tree := parser.NewTree(8)

	if load_error != nil {
		log.Fatalln(load_error.Error())
	}

	config, load_config_error := config.BuildConfig(config_file)

	if load_config_error != nil {
		log.Fatalln("error loding config")
	}

	listener, listener_error := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if listener_error != nil {
		log.Fatalln("Logger error: ", listener_error.Error())
	}

	defer listener.Close()
	log.Println("[LOGGER STARTED] port:" + os.Getenv("PORT"))

	for {
		socket, socket_error := listener.Accept()
		if socket_error != nil {
			log.Println("Socket Error: ", socket_error.Error())
		}

		go handle_socket(socket, config, tree)
	}
}

func handle_socket(socket net.Conn, config config.Config, tree *parser.Tree) {
	log.Println("New connection accepted: ", socket.RemoteAddr().String())
	reader := bufio.NewReader(socket)

	id := 0
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

		request, request_error := parser.ParseRequest(string(message), config.AllowedApplications)
		if request_error != nil {
			log.Println("Error on parsing request: ", request_error.Error())
			break
		}

		parser.DrainParse(tree, request.Message, id)
		id++
	}
}
