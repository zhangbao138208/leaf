package room

type CodeState struct {
	Code    int    // const
	Message string //警告信息
}

var MSG_NOT_IN_ROOM = &CodeState{Code: 1008, Message: "not in room"} // 你不在房间
var MSG_ROOM_CLOSED = &CodeState{Code: 1009, Message: "room closed"} // 房间已经关闭
