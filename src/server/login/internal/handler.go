package internal

import (
	"reflect"
	"wjleafserver/src/server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"wjleafserver/src/server/base/model"
	"wjleafserver/src/server/utils"
	"wjleafserver/src/server/base/funciton"
	"time"
	"fmt"
	"wjleafserver/src/server/dbcenter"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.WjLogin{}, handleLogin)
}

func handleLogin(args []interface{}) {
	var (
		normal_msg msg.NormalMsg
		client_msg msg.ClientMsg
	)
	//recive msg
	m := args[0].(*msg.WjLogin)
	//client agent
	a := args[1].(gate.Agent)

	log.Debug("repeat handleLogin %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())

	user := &model.User{}
	var err error
	user, err = user.FindByUserName(m.Name)
	if err != nil {
		// user not found
	}
	if user.ID > 0 && utils.CheckPassword(m.Password, user.Password) {
		if c, ok := isRepeatLanding(user.ID); ok { // 重复登录
			log.Debug("repeat landing %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())
			//notice old user repeat landing
			client_msg.Msg = a.RemoteAddr().String()
			client_msg.Status = msg.SUM_MSG_STATUS_SUCCESS
			client_msg.Code = msg.MSG_CLIENT_REPEAT_NOTICE
			funciton.SendMsg(c.Agent, &client_msg)

			//response
			client_msg.Msg = ""
			client_msg.Status = msg.SUM_MSG_STATUS_REPEAT
			client_msg.Code = msg.MSG_CLIENT_LOGIN_RSP
			funciton.SendMsg(a, &client_msg)
		} else { // 登录成功
			log.Debug("login success %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())
			token := utils.GetMd5String(a.RemoteAddr().String() + time.Local.String() + fmt.Sprintf("%v", user.ID))

			a.SetUserData(token)

			c := client{}
			c.Token = token
			c.UserID = user.ID
			c.Agent = a
			c.IsOnline = true
			Clients[token] = c

			// 通知游戏模块,处理登录后的用户信息下发, 房间重连,广播房间用户等
			normal_msg.Code = msg.MSG_DBCENTER_LOGIN_NOTICE
			normal_msg.Msg = user.ID
			dbcenter.ChanRPC.Go("DBCENTER", &normal_msg, a)

			// 通知客户端登录成功
			client_msg.Msg = token
			client_msg.Code = msg.MSG_CLIENT_LOGIN_RSP
			funciton.SendMsg(a, &client_msg)
		}
	} else {
		log.Debug("登陆失败 %v/%v@%v", m.Name, m.Password, a.RemoteAddr().String())
		client_msg.Code = msg.MSG_CLIENT_LOGIN_RSP
		client_msg.Status = msg.SUM_MSG_STATUS_ERROR
		client_msg.Msg = "登陆密码或者帐号错误"
		funciton.SendMsg(a, &client_msg)
	}
}

func isRepeatLanding(uid uint) (client, bool) {
	for _, v := range Clients {
		if v.UserID == uid {
			if v.IsOnline {
				return v, true
			}
		}
	}
	return client{}, false
}