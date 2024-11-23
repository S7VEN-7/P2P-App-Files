package p2p

// Peer : interface --> node in the network
type Peer interface{}

// Transport : communication between peers --> TCP, UDP, websockets...
type Transport interface {
	ListenAndAccept() error
}
