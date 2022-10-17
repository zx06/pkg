package broadcaster

type RoomDataChan = chan<- []byte
type Room interface {
	Join(ch RoomDataChan)
	Leave(ch RoomDataChan)
	Message(d []byte)
	Close() error
}
type Broadcaster interface {
	NewRoom(id string) *ChanRoom
	CloseRoom(id string) error
	GetRoomByID(id string) (*ChanRoom, bool)
}
