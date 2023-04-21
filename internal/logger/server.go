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
	if load_error != nil {
		log.Fatalln(load_error.Error())
	}

	config, load_config_error := config.BuildConfig(config_file)

	if load_config_error != nil {
		panic("error loading config")
	}

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

		go handle_socket(socket, config)
	}
}

func handle_socket(socket net.Conn, config config.Config) {
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

		log_message, log_message_error := parser.Parse(string(message), config.AllowedApplications)
		if log_message_error != nil {
			log.Println("Error on parsing log: ", log_message_error.Error())
			break
		}

		for _, rule := range config.Rules {
			rule.Run(log_message.Message)
		}
	}
}
