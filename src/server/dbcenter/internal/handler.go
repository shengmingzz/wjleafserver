package internal

import (
	"reflect"
	"wjleafserver/src/server/msg"
	"github.com/name5566/leaf/gate"
	_ "github.com/name5566/leaf/log"
)

func init()  {
	handler(&msg.RoomMsg{}, handlerRoomMsg)

	skeleton.RegisterChanRPC("DBCENTER", onMessage)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerRoomMsg(args []interface{}) {

}

func onMessage(args []interface{}) {
	m := args[0].(*msg.NormalMsg)
	code := m.Code
	switch code {
		case msg.MSG_DBCENTER_LOGIN_NOTICE:
		onLogin(args)
	}
}

func onLogin(args []interface{}) {
	nmsg := msg.ClientMsg{}
	m := args[0].(*msg.NormalMsg)
	a := args[1].(gate.Agent)
	user_id := m.Msg.(uint)

	u, ok := g_user_mgr.Login(user_id, a)
	if ok {

	} else {
		// 没有获取用户信息
	}
}