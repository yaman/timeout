package timeout

import "net"

type Timeout struct {
	Port     string
	listener net.Listener
}

func (timeout *Timeout) Listen() net.Listener {
	listener, err := net.Listen("tcp", timeout.Port)
	if err != nil {
		panic("An Error Occured while trying to open port: " + timeout.Port)
	}
	timeout.listener = listener
	return listener
}

func (timeout *Timeout) Close() error {
	return timeout.listener.Close()
}
