package broadcaster

import (
	"github.com/pkg/errors"
	"sync"
)

var LocalBroadcaster *localBroadcaster

func init() {
	LocalBroadcaster = &localBroadcaster{
		rooms: make(map[string]*ChanRoom),
	}
}

type ChanRoom struct {
	users     map[RoomDataChan]struct{}
	dataChan  chan []byte
	joinChan  chan RoomDataChan
	leaveChan chan RoomDataChan
}
type localBroadcaster struct {
	sync.Mutex
	rooms map[string]*ChanRoom
}

func (b *localBroadcaster) NewRoom(id string) *ChanRoom {
	b.Lock()
	defer b.Unlock()
	room := NewRoom(5)
	b.rooms[id] = room
	return room
}

func (b *localBroadcaster) CloseRoom(id string) error {
	b.Lock()
	defer b.Unlock()
	room, ok := b.rooms[id]
	if !ok {
		return errors.New("room not found")
	}
	err := room.Close()
	if err != nil {
		return err
	}
	delete(b.rooms, id)
	return nil
}

func (b *localBroadcaster) GetRoomByID(id string) (*ChanRoom, bool) {
	room, ok := b.rooms[id]
	return room, ok
}

func NewRoom(buflen int) *ChanRoom {
	r := &ChanRoom{
		users:     make(map[RoomDataChan]struct{}),
		dataChan:  make(chan []byte, buflen),
		joinChan:  make(chan RoomDataChan),
		leaveChan: make(chan RoomDataChan),
	}
	go r.run()
	return r
}

func (r *ChanRoom) run() {
	for {
		select {
		// 处理提交数据
		case d := <-r.dataChan:
			r.broadcast(d)
		// 处理用户加入消息
		case ch, ok := <-r.joinChan:
			if ok {
				r.users[ch] = struct{}{}
			} else {
				return
			}
		// 处理用户离开消息
		case ch := <-r.leaveChan:
			delete(r.users, ch)
		}
	}
}

// 把消息广播给每一个用户
func (r *ChanRoom) broadcast(d []byte) {
	for ch := range r.users {
		ch <- d
	}
}

func (r *ChanRoom) Join(ch RoomDataChan) {
	r.joinChan <- ch
}

func (r *ChanRoom) Leave(ch RoomDataChan) {
	r.leaveChan <- ch
}

func (r *ChanRoom) Message(d []byte) {
	r.dataChan <- d
}

func (r *ChanRoom) Close() error {
	close(r.joinChan)
	close(r.leaveChan)
	return nil
}
