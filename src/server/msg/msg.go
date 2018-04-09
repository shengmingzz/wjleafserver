package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

//var Processor network.Processor
var Processor = json.NewProcessor()
var ProcessorBuf = protobuf.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&HelloOld{})

	// 登录
	Processor.Register(&WjLogin{})


	ProcessorBuf.Register(&Hello{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type HelloOld struct {
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

//普通消息通知
type NormalMsg struct {
	Code   int
	Status int //0-success 1-failed
	Msg    interface{}
}

type ClientMsg struct {
	Code   int
	Status int //0-success 1-failed
	Msg    interface{}
}

//recv from client game msg
type GameMsg struct {
	Token string
	Code  int
	Msg   interface{}
}

type RoomMsg struct {
	Code   int
	RoomID int
	Msg    interface{}
}

const (
	MSG_CLIENT = 0
	MSG_CLIENT_LOGIN_REQ = MSG_CLIENT + 1 // 用户登录
	MSG_CLIENT_LOGIN_RSP = MSG_CLIENT + 2

	MSG_CLIENT_USERINFO_REQ = MSG_CLIENT + 3 // 用户信息
	MSG_CLIENT_USERINFO_RSP = MSG_CLIENT + 4

	MSG_CLIENT_RECONNECT_REQ = MSG_CLIENT + 5 // 重连
	MSG_CLIENT_RECONNECT_RSP = MSG_CLIENT + 6

	MSG_CLIENT_LOGOUT_REQ = MSG_CLIENT + 7 // 用户退出
	MSG_CLIENT_LOGOUT_RSP = MSG_CLIENT + 8

	MSG_CLIENT_REPEAT_NOTICE = MSG_CLIENT + 9 // repeat landing

	MSG_CLIENT_TOKEN_EXPIRED_NOTICE = MSG_CLIENT + 10 // 登录超时

	/*---------------------dbcenter------------------------*/
	MSG_DBCENTER = 2000

	MSG_DBCENTER_LOGIN_NOTICE       = MSG_DBCENTER + 1 //user login sys notice
	MSG_DBCENTER_LOGOUT_NOTICE      = MSG_DBCENTER + 2 //user logout sys notice
	MSG_DB_CENTER_DISCONNECT_NOTICE = MSG_DBCENTER + 3 //user disconnect sys notice

	/*---------------------room msg------------------------*/
	MSG_ROOM            = 3000
	MSG_ROOM_CREATE_REQ = MSG_ROOM + 1 // 创建房间
	MSG_ROOM_CREATE_RSP = MSG_ROOM + 2 //

	MSG_ROOM_JOIN_REQ = MSG_ROOM + 3 // 加入房间
	MSG_ROOM_JOIN_RSP = MSG_ROOM + 4

	MSG_ROOM_APPLY_DISSOLVE_REQ = MSG_ROOM + 5 // 解散房间申请
	MSG_ROOM_APPLY_DISSOLVE_RSP = MSG_ROOM + 6

	MSG_ROOM_AGREE_DISSOLVE_REQ = MSG_ROOM + 7 // 同意解散请求
	MSG_ROOM_AGREE_DISSOLVE_RSP = MSG_ROOM + 8

	MSG_ROOM_DISSOLVED_NOTICE = MSG_ROOM + 9  // notice room is dissolve
	MSG_ROOM_SYS_INFO_NOTICE  = MSG_ROOM + 10 // room info

	MSG_ROOM_LEAVE_REQ = MSG_ROOM + 11 // user leave
	MSG_ROOM_LEAVE_RSP = MSG_ROOM + 12 //

	MSG_ROOM_STATUS_CHANGE_NOTICE = MSG_ROOM + 13
	MSG_ROOM_SYS_DISSOLVE_NOTICE = MSG_ROOM + 14

	MSG_ROOM_REJOIN_SYS_NOTICE = MSG_ROOM + 15

	/*---------------------game msg------------------------*/
	MSG_GAME_BASE  = 4000
	MSG_GAME_BEGIN = MSG_GAME_BASE + 1 // 游戏开始

	SUM_MSG_STATUS         = 0
	SUM_MSG_STATUS_SUCCESS = SUM_MSG_STATUS
	SUM_MSG_STATUS_ERROR   = SUM_MSG_STATUS + 1
	SUM_MSG_STATUS_REPEAT  = SUM_MSG_STATUS + 2
	SUM_MSG_STATUS_TO_MUCH = SUM_MSG_STATUS + 3

	//0-wait 1-playing 2-apply dissolve 3-disssolved 4-over 5-room not exist 6-user login 7 user logout
	ROOM_STATUS_BASE      = 0
	ROOM_STATUS_WAIT      = ROOM_STATUS_BASE
	ROOM_STATUS_PLAYING   = ROOM_STATUS_BASE + 1
	ROOM_STATUS_APL_DIS   = ROOM_STATUS_BASE + 2
	ROOM_STATUS_DIS_DONE  = ROOM_STATUS_BASE + 3
	ROOM_STATUS_OVER      = ROOM_STATUS_BASE + 4
	ROOM_STATUS_NOT_EXIST = ROOM_STATUS_BASE + 5
	ROOM_STATUS_USER_LOGIN  = ROOM_STATUS_BASE + 6
	ROOM_STATUS_USER_LOGOUT = ROOM_STATUS_BASE + 7
)

type UserInfo struct {
	ID        uint
	UserName  string
	IsOnline  int
	RoomID    int
	RoomCards int
	Status    int
	Nickname  string
	Type      int
}

type RoomBaseInfo struct {

}

//game init msg
type GameConf struct {
	ID            int
	Name          string
	PeopleNumbers int
	GameNumbers   int
}