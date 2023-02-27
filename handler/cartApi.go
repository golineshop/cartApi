package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cart "github.com/golineshop/cart/proto"
	cartApi "github.com/golineshop/cartApi/proto"
	"github.com/micro/go-micro/v2/util/log"
	"strconv"
)

type CartApi struct {
	CartService cart.CartService
}

func (c *CartApi) FindAll(ctx context.Context, req *cartApi.Request, rsp *cartApi.Response) error {
	log.Info("接受到 /cartApi/findAll 访问请求")
	if _, ok := req.Get["user_id"]; !ok {
		//rsp.StatusCode = 500
		return errors.New("参数异常")
	}
	userIdString := req.Get["user_id"].Values[0]
	fmt.Println(userIdString)
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return err
	}
	//获取购物车所有商品
	cartAll, err := c.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	//数据类型转化s
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}
