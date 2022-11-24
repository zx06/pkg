package broadcaster

import (
	"github.com/pkg/errors"
	"sync"
)

var LocalBroadcaster Broadcaster = &localBroadcaster{
	rooms: make(map[string]*ChanRoom),
}

type ChanRoom struct {
	users         map[RoomDataChan]struct{}
	dataChan      chan []byte
	joinChan      chan RoomDataChan
	leaveChan     chan RoomDataChan
	closeRoomChan chan []byte
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

func NewRoom(bufLen int) *ChanRoom {
	r := &ChanRoom{
		users:         make(map[RoomDataChan]struct{}),
		dataChan:      make(chan []byte, bufLen),
		closeRoomChan: make(chan []byte),
		joinChan:      make(chan RoomDataChan),
		leaveChan:     make(chan RoomDataChan),
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

var LocalCloser Closer = &localCloser{
	data: make(map[string]chan string),
}

type localCloser struct {
	sync.Mutex
	data map[string]chan string
}

func (l *localCloser) SendClose(key, reason string) {
	l.Lock()
	defer l.Unlock()
	c, ok := l.data[key]
	if ok {
		c <- reason
		// 由于使用的是无缓冲channel,上面的操作是阻塞的,可以在消息被消费后删除该channel
		delete(l.data, key)
	}

}

func (l *localCloser) SubClose(key string) chan string {
	l.Lock()
	defer l.Unlock()
	c := make(chan string)
	l.data[key] = c
	return c
}
