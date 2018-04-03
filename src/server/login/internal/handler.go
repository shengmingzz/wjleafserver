package internal

import (
	"reflect"
	"wjleafserver/src/server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.WjLogin{}, handleLogin)
}

func handleLogin(args []interface{})  {
	//recive msg
	m := args[0].(*msg.WjLogin)
	//client agent
	a := args[1].(gate.Agent)

	log.Debug("repeat handleLogin %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())
}