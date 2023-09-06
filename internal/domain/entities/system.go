package entities

type System struct {
	Graph     Graph
	Resources []Resource
	Mailbox   chan Command
}
