package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

//var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&Hello{})

	// 登录
	Processor.Register(&WjLogin{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type Hello struct {
	Name string
}

// 登录消息
type WjLogin struct {
	Name string
	Password string
}
// 微信登录
type WjWechatLogin struct {
	WechatId string
}

// 注册消息
type WjRegister struct {
	
}

// 心跳消息
type WjHeatbeat struct {

}

// 重连
type WjReConnect struct {

}

