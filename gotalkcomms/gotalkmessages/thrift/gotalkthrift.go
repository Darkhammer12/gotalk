package thriftmsgs

import (
	"errors"
	"gotalk/gotalkcomms/gotalkmessages/thrift/gen-go/gotalkThrift"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var thriftProtocolFactory thrift.TProtocolFactory
var thriftTransportFactory thrift.TTransportFactory

//assume some defaults
func init() {
	thriftProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	thriftTransportFactory = thrift.NewTTransportFactory()
}

type GotalkHandler struct {
	resultChan chan interface{}
}

func NewHGotalkHandler(o chan interface{}) *GotalkHandler {
	return &GotalkHandler{
		resultChan: o,
	}
}

func (handler *GotalkHandler) AddShip(s *GotalkThrift.Ship) (err error) {
	handler.resultChan <- s
	return
}

func StartThriftServer(addr string, h *GotalkHandler) (err error) {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return
	}
	processor := gotalkThrift.NewHGotalkThriftServiceProcessor(h)
	server := thrift.NewTSimpleServer4(processor, transport, thriftTransportFactory, thriftProtocolFactory)
	return server.Serve()
}

func RunThriftClient(obj interface{}, addr string) (err error) {
	s, ok := obj.(*gotalkThrift.Ship)
	if !ok {
		return errors.New("Unknown type.. ")
	}
	sock, err := thrift.NewTSocket(addr)
	if err != nil {
		return
	}
	log.Println("Thrift socket created")
	transport := thriftTransportFactory.GetTransport(sock)
	err = transport.Open()
	if err != nil {
		return
	}
	defer transport.Close()
	log.Println("Thrift transport opened")
	client := gotalkThrift.NewGotalkThriftServiceClientFactory(transport, thriftProtocolFactory)
	log.Println("Calling add")
	err = client.AddShip(s)
	return
}
