package impl

import (
	"fmt"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/ordermgtmodel"
	"github.com/prakash-p-3121/restclientlib"
	"log"
	"net/url"
)

type OrderMgtClientImpl struct {
	Host string
	Port uint
}

func (client *OrderMgtClientImpl) HostPort() string {
	return fmt.Sprintf("%s:%d", client.Host, client.Port)
}

func (client *OrderMgtClientImpl) OrderCreate(req *ordermgtmodel.OrderCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	restClient := restclientlib.NewRestClient()
	hostPort := client.HostPort()
	url := hostPort + orderCreate
	log.Println("url =", url)
	var resp idgenmodel.IDGenResp
	appErr := restClient.Post(url, req, &resp)
	return &resp, appErr
}

func (client *OrderMgtClientImpl) OrderFindByID(orderID string) (*ordermgtmodel.Order, errorlib.AppError) {
	restClient := restclientlib.NewRestClient()
	hostPort := client.HostPort()
	baseUrl := hostPort + orderFindByID
	log.Println("url =", baseUrl)

	params := url.Values{}
	params.Add("order-id", orderID)
	encodedParams := params.Encode()
	log.Println("encodedParams=" + encodedParams)
	finalUrl := baseUrl + "?" + encodedParams
	log.Println("finalUrl=" + finalUrl)

	var resp ordermgtmodel.Order
	appErr := restClient.Get(finalUrl, &resp)
	return &resp, appErr
}
