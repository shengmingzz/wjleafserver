package gate

import (
	"wjleafserver/src/server/msg"
	"wjleafserver/src/server/game"
	"wjleafserver/src/server/login"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.HelloOld{}, game.ChanRPC)

	// 登录
	msg.Processor.SetRouter(&msg.WjLogin{}, login.ChanRPC)

	// protobuf
	msg.ProcessorBuf.SetRouter(&msg.Hello{}, game.ChanRPC)
}
