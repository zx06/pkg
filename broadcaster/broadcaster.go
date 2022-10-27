package broadcaster

type RoomDataChan = chan<- []byte

// Room 基于channel实现的发布订阅模型,以房间作为一组订阅者的抽象集合
type Room interface {
	// Join 加入房间,传入一个channel,开始向该channel广播房间里面的消息
	Join(ch RoomDataChan)
	// Leave 离开房间,不再向该队列广播消息
	Leave(ch RoomDataChan)
	// Message 向房间发送消息,消息会广播给房间中所有的channel
	Message(d []byte)
	// Close 关闭房间,停止广播
	Close() error
}

// Broadcaster 一个广播组
type Broadcaster interface {
	// NewRoom 根据id新建一个房间
	NewRoom(id string) *ChanRoom
	// CloseRoom 关闭房间
	CloseRoom(id string) error
	// GetRoomByID 根据id获取房间
	GetRoomByID(id string) (*ChanRoom, bool)
}

// Closer 一个基于channel的关闭通知
type Closer interface {
	// SendClose 向指定key发送关闭消息,可以的带上关闭原因
	SendClose(key, reason string)
	// SubClose 订阅关闭消息,返回一个channel; ps: 因为只有订阅了关闭才有意义,所以channel创建是在订阅时做的
	SubClose(key string) chan string
}
