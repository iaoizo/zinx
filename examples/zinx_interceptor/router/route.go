package router

import (
	"github.com/iaoizo/zinx/ziface"
	"github.com/iaoizo/zinx/zlog"
	"github.com/iaoizo/zinx/znet"
)

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF(string(request.GetData()))
}
