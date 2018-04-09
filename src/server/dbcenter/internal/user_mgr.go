package internal

import (
	"wjleafserver/src/server/base/model"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"wjleafserver/src/server/msg"
)

var (
	g_user_mgr       *UserMgr // 用户记录
	g_tokenAndUidMap map[string]uint // 用户token记录
)

func init()  {
	g_user_mgr = new(UserMgr)
	g_user_mgr.users = map[uint]*model.User{}
	g_tokenAndUidMap = map[string]uint{}
}

type UserMgr struct {
	users map[uint]*model.User
}

func (mgr *UserMgr)Get(uid uint) *model.User {
	u, ok := mgr.users[uid]
	if ok {
		return u
	}
	return nil
}

func (mgr *UserMgr)Login(user_id uint, a gate.Agent) (*model.User, bool) {
	u, ok := mgr.users[user_id]
	if ok { // 有登录记录
		token := a.UserData().(string)

		u.IsOnline = 1
		u.Agent = a
		mgr.users[user_id] = u
		g_tokenAndUidMap[token] = user_id

		log.Debug("MEM=>userid=%d, name=%s, isonline=%d, room_id=%d", u.ID, u.UserName, u.IsOnline, u.RoomID)
	} else { // 查找库,存储登录信息
		var err error
		new_user := new(model.User)
		new_user, err = new_user.FindById(user_id)
		if err != nil {
			return new_user, false
		}
		new_user.IsOnline = 1
		new_user.Agent = a
		u = new_user

		token := a.UserData().(string)
		g_tokenAndUidMap[token] = user_id
		log.Debug("DB=>userid=%d, name=%s, isonline=%d, room_id=%d", u.ID, u.UserName, u.IsOnline, u.RoomID)
	}
	mgr.users[user_id] = u
	return u, true
}

func (mgr *UserMgr)GetUserBaseInfo(u *model.User) *msg.UserInfo {
	user_info := new(msg.UserInfo)
	user_info.ID = u.ID
	user_info.UserName = u.UserName
	user_info.IsOnline = u.IsOnline
	user_info.RoomID = u.RoomID
	user_info.RoomCards = u.RoomCards
	user_info.Status = u.Status
	user_info.Nickname = u.Nickname
	user_info.Type = u.Type
	return user_info
}