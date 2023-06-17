package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"oju/internal/config"
	"oju/internal/proxy"
	"oju/internal/requester"
	"os"
)

func main() {
	config_file, load_error := config.LoadConfigFile()

	if load_error != nil {
		log.Fatalln(load_error.Error())
	}

	config, load_config_error := config.BuildConfig(config_file)

	if load_config_error != nil {
		log.Fatalln("error loding config")
	}

	manager := proxy.NewManager(config.AllowedApplications)

	listener, listener_error := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if listener_error != nil {
		log.Fatalln("Oju error: ", listener_error.Error())
	}

	defer listener.Close()
	log.Println("[OJU STARTED] port:" + os.Getenv("PORT"))

	for {
		socket, socket_error := listener.Accept()
		if socket_error != nil {
			log.Println("Socket Error: ", socket_error.Error())
		}

		go handle_incoming_message(socket, config, manager)
	}
}

func handle_incoming_message(socket net.Conn, config config.Config, manager *proxy.Manager) {
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
			manager.Redirect(request.Header.AppKey, proxy.ApplicationMessage{
				Type:    "LOG",
				Payload: request.Message,
			})
			break
		case "TRACE":
			manager.Redirect(request.Header.AppKey, proxy.ApplicationMessage{
				Type:    "TRACE",
				Payload: request.Message,
			})
			break
		case "WATCH":
			// install watcher
			fmt.Println("Installing watcher")
		default:
		}

	}
}
