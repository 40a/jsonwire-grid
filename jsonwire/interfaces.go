package jsonwire

type ClientInterface interface {
	Sessions() (*Sessions, error)
	CloseSession(sessionId string) (*Message, error)
	Address() string
}

type NodeInterface interface {
	RemoveAllSessions() (int, error)
}
