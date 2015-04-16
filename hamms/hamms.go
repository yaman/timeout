package hamms

import "net"

type Hamms struct {
	Port     string
	listener net.Listener
}

func (hamms *Hamms) Listen() net.Listener {
	listener, err := net.Listen("tcp", hamms.Port)
	if err != nil {
		panic("An Error Occured while trying to open port: " + hamms.Port)
	}
	hamms.listener = listener
	return listener
}

func (hamms *Hamms) Close() error {
	return hamms.listener.Close()
}
