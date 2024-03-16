package ordermgtclient

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/ordermgtmodel"
)

type OrderMgtClient interface {
	OrderCreate(req *ordermgtmodel.OrderCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError)
	OrderFindByID(orderID string) (*ordermgtmodel.Order, errorlib.AppError)
}
