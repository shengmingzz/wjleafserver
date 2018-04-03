package internal

import (
	"reflect"
	"wjleafserver/src/server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"wjleafserver/src/server/base/model"
	"wjleafserver/src/server/utils"
	"wjleafserver/src/server/base/funciton"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.WjLogin{}, handleLogin)
}

func handleLogin(args []interface{}) {
	//recive msg
	m := args[0].(*msg.WjLogin)
	//client agent
	a := args[1].(gate.Agent)

	log.Debug("repeat handleLogin %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())

	user := model.User{}
	var err error
	user, err = user.FindByUserName(m.Name)
	if err != nil {
		// user not found
	}
	if user.ID > 0 && utils.CheckPassword(m.Password, user.Password) {

	} else {
		log.Debug("登陆失败 %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())

		funciton.SendMsg(a, map[string]interface{})
	}
}