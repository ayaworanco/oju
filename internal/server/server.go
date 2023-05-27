package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"oju/internal/config"
	"oju/internal/parser"
	"oju/internal/requester"
)

func Start() {
	config_file, load_error := config.LoadConfigFile()

	// FIXME: there is needed a parse tree to each one of applications
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

		go handle_incoming_message(socket, config, tree)
	}
}

func handle_incoming_message(socket net.Conn, config config.Config, tree *parser.Tree) {
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
		request, request_error := requester.Parse(string(message), config.AllowedApplications)
		if request_error != nil {
			log.Println("Error on parsing request: ", request_error.Error())
			break
		}

		switch request.Header.Verb {
		case "LOG":
			parser.ParseLog(tree, request.Message)
			break
		case "WATCH":
			// install watcher
			fmt.Println("Installing watcher")
		default:
		}

	}
}
