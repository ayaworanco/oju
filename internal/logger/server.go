package logger

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"oluwoye/internal/parser"
	"oluwoye/internal/ruler"
	"os"
)

func StartLogger() {
	rules_file, load_error := load_rules()
	if load_error != nil {
		log.Fatalln(load_error.Error())
	}

	listener, listener_error := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if listener_error != nil {
		panic("Logger error: " + listener_error.Error())
	}

	rules, load_rules_error := ruler.LoadRules(rules_file)
	if load_rules_error != nil {
		panic("Error loading rules")
	}

	defer listener.Close()
	log.Println("[LOGGER STARTED] port:" + os.Getenv("PORT"))

	for {
		socket, socket_error := listener.Accept()
		if socket_error != nil {
			log.Println("Socket Error: ", socket_error.Error())
		}

		go handle_socket(socket, rules)
	}
}

func load_rules() ([]byte, error) {
	var rule_file string
	rule_file = os.Getenv("RULES_YAML_PATH")

	if rule_file == "" {
		rule_file = "rules.yaml"
	}

	file, read_error := os.ReadFile(rule_file)

	if read_error != nil {
		return nil, errors.New(rule_file + " file not found")
	}

	return file, nil
}

func handle_socket(socket net.Conn, rules []ruler.Rule) {
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

		for _, rule := range rules {
			rule.Run(log_message.Message)
		}
	}
}
