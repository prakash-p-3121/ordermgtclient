package ordermgtclient

import "github.com/prakash-p-3121/ordermgtclient/impl"

func NewOrderMgtClient(host string, port uint) OrderMgtClient {
	return &impl.OrderMgtClientImpl{Host: host, Port: port}
}
