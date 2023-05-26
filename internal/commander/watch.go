package commander

import (
	"flag"
	"fmt"
	"net"
	"os"

	"oju/internal/parser"
)

var (
	host string
	app  string
)

var watch_handler = func(cmd *Command, args []string) {
	// TODO: this should listen to a config.json file for oju
	// if there is this config, we can check all inside the file

	if len(args) == 0 {
		fmt.Println("This command should have at least a query\n-------------------")
		cmd.flags.Usage()
		os.Exit(0)
	}

	query := args[0]
	if query == "" {
		fmt.Println("This query should not be empty\n-------------------")
		cmd.flags.Usage()
		os.Exit(0)
	}

	if len(host) == 0 {
		fmt.Println("This command should have a host\n-------------------")
		cmd.flags.Usage()
		os.Exit(0)
	}

	if app == "" {
		fmt.Println("This command should have an app key\n-------------------")
		cmd.flags.Usage()
		os.Exit(0)
	}

	send_watch_message(host, app, query)
	os.Exit(0)
}

func WatchCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("watch", flag.ExitOnError),
		Execute: watch_handler,
	}

	cmd.flags.StringVar(&host, "host", "", "--host 127.0.0.1:8080")
	cmd.flags.StringVar(&app, "app", "", "--app ABC@123")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, watch_usage)
	}
	return cmd
}

func send_watch_message(host, app, query string) {
	request, error_request := create_watch_request(app, query)

	if error_request != nil {
		error_message(error_request.Error())
	}

	connection, connection_error := net.Dial("tcp", host)
	if connection_error != nil {
		error_message(connection_error.Error())
	}

	defer connection.Close()

	_, error_send_message := connection.Write([]byte(request.String()))
	if error_send_message != nil {
		error_message(error_send_message.Error())
	}

	for {
		// wait for responses from server
	}
}

func create_watch_request(app, query string) (parser.Request, error) {
	head := fmt.Sprintf("WATCH %s AWO1.1", app)

	request, request_error := parser.NewRequest(head, query)
	if request_error != nil {
		return parser.Request{}, request_error
	}

	return request, nil
}

func error_message(msg string) {
	fmt.Printf("%s\n-------------------", msg)
}
