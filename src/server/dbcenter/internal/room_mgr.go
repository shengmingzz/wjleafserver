package internal

import (
	"qpgame/base/function"
	"wjleafserver/src/server/msg"
	"wjleafserver/src/server/base/model"
	"github.com/name5566/leaf/log"
	"time"
)

var (
	g_roomMgr *RoomMgr
)

type RoomMgr struct {
	rooms map[int]*Room
}

func init() {
	g_roomMgr = new(RoomMgr)
	g_roomMgr.rooms = map[int]*Room{}
}

func (rmg *RoomMgr)Get(room_id int) (*Room, bool)  {
	room, ok := rmg.rooms[room_id]
	if ok {
		return room, true
	} else {
		return nil, false
	}
}

// 创建房间号,随机获取
func (rmg *RoomMgr) CreateRoomID() int {
	room_id := function.RandInt64(100000,999999)
	_, ok := rmg.Get(room_id)
	if !ok {
		return room_id
	} else {
		return rmg.CreateRoomID()
	}
}

func (rmg *RoomMgr)IsInRoom(uid uint)  {

}

func (rmg *RoomMgr)Create(u *model.User, game_conf msg.GameConf) (*Room, int) {
	if u.RoomID != 0 { // 已经有房间
		log.Debug("Create User is in room %d", u.RoomID)
		return nil, msg.SUM_MSG_STATUS_REPEAT
	}

	room := new(Room)
	room.ID = rmg.CreateRoomID()

	if _, ok := rmg.rooms[room.ID]; ok { // 房间在使用
		log.Debug("room id %d is be used", room.ID)
		return nil, msg.SUM_MSG_STATUS_REPEAT
	}

	room.Name = game_conf.Name
	room.PeopleNumbers = game_conf.PeopleNumbers
	room.GameNumbers = game_conf.GameNumbers
	room.CreatedAt = time.Now()
	room.CreatedBy = u.ID
	room.Join(u)
	room.ZhuangIndex = 0

	room.SetReady(u)
	var gam
}