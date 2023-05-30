package routing

import (
	"errors"

	"oju/internal/application"
	"oju/internal/config"
)

type Proxy struct {
	Applications []*application.Application
	Mailbox      chan ProxyMessage
}

type ProxyMessage struct {
	Destination string
	Payload     interface{}
}

func NewProxy(allowed_applications []config.Application) *Proxy {
	proxy := &Proxy{
		Applications: make([]*application.Application, 0),
		Mailbox:      make(chan ProxyMessage),
	}

	for _, allowed := range allowed_applications {
		proxy.Applications = append(proxy.Applications, get_app(allowed))
	}
	go proxy.run()
	return proxy
}

func (proxy *Proxy) Redirect(destination string, payload interface{}) {
	message := ProxyMessage{
		Destination: destination,
		Payload:     payload,
	}
	proxy.Mailbox <- message
}

func (proxy *Proxy) GetApp(name string) (*application.Application, error) {
	for _, app := range proxy.Applications {
		metadata := app.GetMetadata()
		if metadata.Host == name || metadata.Key == name {
			return app, nil
		}
	}
	return nil, errors.New("application not found")
}

func (proxy *Proxy) run() {
	for {
		proxy.handle_message(<-proxy.Mailbox)
	}
}

func (proxy *Proxy) handle_message(message ProxyMessage) {
	for _, app := range proxy.Applications {
		metadata := app.GetMetadata()
		if metadata.Host == message.Destination || metadata.Key == message.Destination {
			app.SendMessage(message.Payload.(application.Message))
		}
	}
}

func get_app(config_app config.Application) *application.Application {
	return application.Start(10, application.Metadata{
		Key:  config_app.AppKey,
		Host: config_app.Host,
	})
}
