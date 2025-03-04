package main

import (
	"fmt"
	"github.com/iaoizo/zinx/zdecoder"
	"github.com/iaoizo/zinx/zpack"

	"github.com/iaoizo/zinx/ziface"
	"github.com/iaoizo/zinx/zinx_app_demo/mmo_game/api"
	"github.com/iaoizo/zinx/zinx_app_demo/mmo_game/core"
	"github.com/iaoizo/zinx/znet"
)

// 当客户端建立连接的时候的hook函数
func OnConnecionAdd(conn ziface.IConnection) {
	fmt.Println("=====> OnConnecionAdd is Called ...")
	//创建一个玩家
	player := core.NewPlayer(conn)

	//同步当前的PlayerID给客户端， 走MsgID:1 消息
	player.SyncPID()

	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	player.BroadCastStartPosition()

	//将当前新上线玩家添加到worldManager中
	core.WorldMgrObj.AddPlayer(player)

	//将该连接绑定属性PID
	conn.SetProperty("pID", player.PID)

	//同步周边玩家上线信息，与现实周边玩家信息
	player.SyncSurrounding()

	fmt.Println("=====> Player pIDID = ", player.PID, " arrived ====")
}

// 当客户端断开连接的时候的hook函数
func OnConnectionLost(conn ziface.IConnection) {
	//获取当前连接的PID属性
	pID, _ := conn.GetProperty("pID")
	var playerID int32
	if pID != nil {
		playerID = pID.(int32)
	}

	//根据pID获取对应的玩家对象
	player := core.WorldMgrObj.GetPlayerByPID(playerID)

	//触发玩家下线业务
	if player != nil {
		player.LostConnection()
	}

	fmt.Println("====> Player ", playerID, " left =====")

}

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//注册客户端连接建立和丢失函数
	s.SetOnConnStart(OnConnecionAdd)
	s.SetOnConnStop(OnConnectionLost)

	//注册路由
	s.AddRouter(2, &api.WorldChatApi{})
	s.AddRouter(3, &api.MoveApi{})

	//添加LTV数据格式Decoder
	s.SetDecoder(zdecoder.NewLTV_Little_Decoder())
	//添加LTV数据格式的Pack封包Encoder
	s.SetPacket(zpack.NewDataPackLtv())

	//启动服务
	s.Serve()
}
